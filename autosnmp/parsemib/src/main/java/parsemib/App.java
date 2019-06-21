package parsemib;

import org.antlr.v4.runtime.ANTLRInputStream;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.CommonTokenStream;
import org.antlr.v4.runtime.TokenSource;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
        String input = getstrfromarg("D:\\github\\kill-that-programmer\\autosnmp\\parsemib\\mibs\\RFC1213-MIB");
        CharStream stream = new ANTLRInputStream(input);
        TokenSource lexer = new SNMPMIBLexer(stream);

        CommonTokenStream tokenStream = new CommonTokenStream(lexer);
        SNMPMIBParser parser = new SNMPMIBParser(tokenStream);

        myvisitor visitor = new myvisitor();
        visitor.visit(parser.moduleDefinition());
        /*visitor.m_mapnametooid.forEach((key, val)->{
            System.out.println(key + " : "+val);
        });*/
        visitor.recursivetree(visitor.rootitem, 0);
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


}
