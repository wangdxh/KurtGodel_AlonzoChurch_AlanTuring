package parsemib;

import com.alibaba.fastjson.JSON;
import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.TokenSource;

import javax.json.Json;
import javax.json.JsonObject;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.util.*;

/**
 * Hello world!
 *
 */
public class App 
{
    public static Map<String, iteminfo> m_tree = new LinkedHashMap<>();
    public static void main( String[] args )
    {
        //Map<String, iteminfo> tree = new LinkedHashMap<>();
        iteminfo rootitem = null;
        rootitem = new iteminfo("mgmt", ".1.3.6.1.2", ItemType.NONE, "");
        m_tree.put(rootitem.name, rootitem);

        //String input = getstrfromarg();
        String[] mibifles = {"D:\\github\\kill-that-programmer\\autosnmp\\parsemib\\mibs\\RFC1213-MIB",
                "D:\\github\\kill-that-programmer\\autosnmp\\parsemib\\mibs\\Bridge.mib",
                "D:\\github\\kill-that-programmer\\autosnmp\\parsemib\\mibs\\HOST-RESOURCES-MIB"};

        for (String fileitem:mibifles)
        {
            String input = getstrfromarg(fileitem);
            CharStream stream = new ANTLRInputStream(input);
            TokenSource lexer = new SNMPMIBLexer(stream);

            CommonTokenStream tokenStream = new CommonTokenStream(lexer);
            SNMPMIBParser parser = new SNMPMIBParser(tokenStream);

            myvisitor visitor = new myvisitor(m_tree);
            visitor.visit(parser.moduleDefinition());
        }
        App.recursivetree(rootitem, 0);
        App.getallmibtypes(rootitem);
        m_mapMibTypes.forEach((key, val)->{
            System.out.printf("put(\"%s\", \"%s\");\n", key, val);
        });
        App.generatecode(rootitem, 0);
        String strdata = JSON.toJSONString(m_tree, true);
        System.out.println(strdata);
        FileWriter writer;
        try {
            writer = new FileWriter("D:\\github\\kill-that-programmer\\autosnmp\\snmpmibtoken.txt");
            writer.write("");//清空原文件内容
            writer.write(strdata);
            writer.flush();
            writer.close();
        } catch (IOException e)
        {
            e.printStackTrace();
        }

        return;
    }

    public  static String getstrfromarg(String arg){
        File file = new File(arg);
        if (file.exists()){
            String encoding = "UTF-8";
            Long filelength = file.length();
            byte[] filecontent = new byte[filelength.intValue()];
            FileInputStream in = null;
            try
            {
                in = new FileInputStream(file);
                try {
                    in.read(filecontent);
                    return new String(filecontent, encoding);
                } catch (Exception e) {
                    e.printStackTrace();
                    return "sble";
                } finally {
                    in.close();
                }
            }
            catch (IOException e)
            {
                e.printStackTrace();
                return arg;
            }
        }

        return arg;
    }

    public static SortedMap<String, String> m_mapMibTypes = new TreeMap<>();
    public static void getallmibtypes(iteminfo item){
        //String strtab = String.join("", Collections.nCopies(deep, "  "));
        //System.out.println(strtab + item.name + " : "+item.oid+":"+item.strtype);
        if (item.itemtype == ItemType.OBJECT && item.children==null &&
        !m_mapMibTypes.containsKey(item.strtype)){
            m_mapMibTypes.put(item.strtype, item.strtype);
        }

        if (item.children != null){
            for (String sonname :item.children)
            {
                iteminfo son = m_tree.get(sonname);
                getallmibtypes(son);
            }
        }
    }

    public static void recursivetree(iteminfo item, int deep){
        String strtab = String.join("", Collections.nCopies(deep, "  "));
        System.out.println(strtab + item.name + " : "+item.oid+":"+item.strtype);
        if (item.children != null){
            for (String sonname :item.children)
            {
                iteminfo son = m_tree.get(sonname);
                recursivetree(son, deep+1);
            }
        }
    }
    public static String getgotype(String mibtype){
        if (!m_mapTypetoGoType.containsKey(mibtype)){
            System.out.println("when generate code no type to go type:" + mibtype);
            int y = 100/10;
        }
        return m_mapTypetoGoType.get(mibtype);
    }
    private static String captureName(String str) {
        // 进行字母的ascii编码前移，效率要高于截取字符串进行转换的操作
        char[] cs=str.toCharArray();
        cs[0]-=32;
        return String.valueOf(cs);
    }

    public static void generatecode(iteminfo item, int deep){
        if (item.itemtype == ItemType.SEQUENCE && item.children !=null && item.children.size() ==1){
            System.out.printf("type %s struct  {\n", captureName(item.name));
            iteminfo itemson = m_tree.get(item.children.get(0));
            for (String sunzi:itemson.children)
            {
                iteminfo itemsunzi = m_tree.get(sunzi);
                String strtag = String.format(" `snmp:\"%s\"`", itemsunzi.strtype);
                System.out.println("\t" + captureName(itemsunzi.name) + " " + getgotype(itemsunzi.strtype) +
                        strtag);
            }
            System.out.println("}");
            return;
        }
        if (item.itemtype == ItemType.OBJECT && item.children == null){
            System.out.println(item.name + ":" +getgotype(item.strtype));
        }

        if (item.children != null){
            for (String sonname :item.children)
            {
                iteminfo son = m_tree.get(sonname);
                generatecode(son, deep+1);
            }
        }
    }
    public static Map<String , String>m_mapTypetoGotag = new HashMap<String, String>(){{

    }};

    public static Map<String , String>m_mapTypetoGoType = new HashMap<String, String>(){{
        put("AutonomousType", "string");
        put("BridgeId", "string");
        put("Counter", "int");
        put("Counter32", "int");
        put("DateAndTime", "string");
        put("DisplayString", "string");
        put("Gauge", "int");
        put("Gauge32", "int");
        put("INTEGER", "int");
        put("InterfaceIndex", "int");
        put("InterfaceIndexOrZero", "int");
        put("InternationalDisplayString", "string");
        put("IpAddress", "string");
        put("KBytes", "int");
        put("MacAddress", "string");
        put("NetworkAddress", "string");
        put("PhysAddress", "string");
        put("ProductID", "string");
        put("String", "string");
        put("TimeTicks", "int");
        put("Timeout", "int");
        put("TruthValue", "int");
    }};
}
