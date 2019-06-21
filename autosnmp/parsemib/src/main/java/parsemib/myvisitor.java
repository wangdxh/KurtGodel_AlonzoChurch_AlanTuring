package parsemib;

import java.util.*;

class iteminfo{
    public iteminfo(String strname, String stroid)
    {
        this.name = strname;
        this.oid = stroid;
    }
    public void addchild(String itemname){
        if (children==null){
            children=new ArrayList<>();
        }
        children.add(itemname);
    }
    public String name;
    public String oid;
    public List< String> children;
}

public class myvisitor extends SNMPMIBBaseVisitor<String>
{

    //public Map<String, String> m_mapnametooid = new LinkedHashMap<>();
    public Map<String, iteminfo> m_maptree= new LinkedHashMap<>();
    public iteminfo rootitem = null;

    public myvisitor(){
        //this.m_mapnametooid.put("mgmt", ".1.3.6.1.2");
        rootitem = new iteminfo("mgmt", ".1.3.6.1.2");
        this.m_maptree.put("mgmt", rootitem);
    }

    public void recursivetree(iteminfo item, int deep){
        String strtab = String.join("", Collections.nCopies(deep, "  "));
        System.out.println(strtab + item.name + " : "+item.oid);
        if (item.children != null){
            for (String sonname :item.children)
            {
                iteminfo son = this.m_maptree.get(sonname);
                recursivetree(son, deep+1);
            }
        }
    }

    public void addnewitem(String fathername, String itemname, String itemnumber){
        if (this.m_maptree.containsKey(fathername) == false){
            System.out.println("do not find father 's oid, there one error , fahtername is :"+fathername);
            System.exit(-1);
        }
        iteminfo father = this.m_maptree.get(fathername);
        iteminfo son = new iteminfo(itemname, father.oid+"."+itemnumber);
        this.m_maptree.put(itemname, son);
        father.addchild(itemname);
    }

    @Override
    public String visitTypeAssignment(SNMPMIBParser.TypeAssignmentContext ctx)
    {
        System.out.println("new type:" + ctx.IDENTIFIER().getText());
        return super.visitTypeAssignment(ctx);
    }

    @Override
    public String visitValueAssignment(SNMPMIBParser.ValueAssignmentContext ctx)
    {
        System.out.println("new value:" + ctx.IDENTIFIER().get(0).getText());

        String myname = ctx.IDENTIFIER().get(0).getText();
        String fathername = ctx.IDENTIFIER().get(1).getText();
        String mynumber = ctx.NUMBER().getText();

        addnewitem(fathername, myname, mynumber);

        return super.visitValueAssignment(ctx);
    }
}
