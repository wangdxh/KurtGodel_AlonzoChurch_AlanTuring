package parsemib;

import java.util.*;
enum ItemType{ NONE, OBJECT, SEQUENCE};

class iteminfo{
    public iteminfo(String strname, String stroid, ItemType itemtype, String strtype)
    {
        this.name = strname;
        this.oid = stroid;
        this.itemtype = itemtype;
        this.strtype = strtype;
    }
    public void addchild(String itemname){
        if (children==null){
            children=new ArrayList<>();
        }
        children.add(itemname);
    }
    public String name;
    public String oid;
    public String strtype;
    public ItemType itemtype;
    public List< String> children;
}

public class myvisitor extends SNMPMIBBaseVisitor<String>
{
    //public Map<String, String> m_mapnametooid = new LinkedHashMap<>();
    private Map<String, iteminfo> m_maptree;


    public myvisitor(Map<String, iteminfo> tree){
        m_maptree = tree;
    }

    public void addnewitem(String fathername, String itemname, String itemnumber){
        if (!this.m_maptree.containsKey(fathername)){
            System.out.println("do not find father 's oid, there one error , fahtername is :"+fathername);
            System.exit(-1);
        }
        iteminfo father = this.m_maptree.get(fathername);
        iteminfo son = new iteminfo(itemname, father.oid+"."+itemnumber, m_itemtype, m_strSyntax);
        this.m_maptree.put(itemname, son);
        father.addchild(itemname);
    }

    @Override
    public String visitTypeAssignment(SNMPMIBParser.TypeAssignmentContext ctx)
    {
        System.out.println("new type:" + ctx.IDENTIFIER().getText());
        return super.visitTypeAssignment(ctx);
    }

    private String m_strSyntax = "";
    private ItemType m_itemtype = ItemType.NONE;

    @Override
    public String visitValueAssignment(SNMPMIBParser.ValueAssignmentContext ctx)
    {
        m_itemtype= ItemType.NONE;
        m_strSyntax = "";
        if (ctx.valuetype()!=null && ctx.valuetype().getText().equals("OBJECT-TYPE")){
            // object type define should generate the code for access
            m_itemtype= ItemType.OBJECT;

            if (ctx.valuedesctype()!=null){
                ctx.valuedesctype().forEach((item)->{
                    super.visitValuedesctype(item);
                });
            }
        }
        String myname = ctx.IDENTIFIER().get(0).getText();
        String fathername = ctx.IDENTIFIER().get(1).getText();
        String mynumber = ctx.NUMBER().getText();
        System.out.println(myname + " ---" + m_itemtype + "---" + m_strSyntax);
        this.addnewitem(fathername, myname, mynumber);

        //return super.visitValueAssignment(ctx);
        return "";
    }

    @Override
    public String visitSyntaxdesctype(SNMPMIBParser.SyntaxdesctypeContext ctx)
    {
        if (ctx.IDENTIFIER()!=null){
            // this is sequence
            m_itemtype = ItemType.SEQUENCE;
            m_strSyntax = ctx.IDENTIFIER().getText();
        }
        if (ctx.asnType() != null){
            if (ctx.asnType().IDENTIFIER()!=null){
                m_strSyntax =ctx.asnType().IDENTIFIER().getText();
            }else if (ctx.asnType().builtinType().integerType()!=null){
                m_strSyntax ="INTEGER";
            }else {
                m_strSyntax ="String";
            }
        }
        return "";
    }



}
