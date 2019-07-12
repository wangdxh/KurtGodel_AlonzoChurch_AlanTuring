package main

import (
	"./parser"
	"encoding/json"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/flosch/pongo2"
	"strings"
)

type Where struct {
	Operate  string // atom and or
	Test     string
	Sql      string
	Itemlist []Where
}

type Visitor struct {
	parser.BasePREFIXVisitor
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *Visitor) VisitExpress(ctx *parser.ExpressContext) interface{} {
	ope := v.visitRule(ctx.Operate())
	test := v.visitRule(ctx.Test())
	where := Where{Operate: ope.(string), Test: test.(string)}

	fmt.Println(ope, test, "in expression")
	subexplist := ctx.AllExpression()
	for _, value := range subexplist {
		where.Itemlist = append(where.Itemlist, v.visitRule(value).(Where))
	}

	return where
}

func (v *Visitor) VisitExpsignle(ctx *parser.ExpsignleContext) interface{} {
	fmt.Println("string ", ctx.GetText())
	where := Where{
		Operate: "atom", Test: strings.Trim(ctx.Test().GetText(), `"`),
		Sql: strings.Trim(ctx.CSTRING().GetText(), `"`),
	}
	return where
}

func (v *Visitor) VisitExpstring(ctx *parser.ExpstringContext) interface{} {
	fmt.Println("string ", ctx.GetText())
	where := Where{Operate: "atom", Test: "true", Sql: strings.Trim(ctx.GetText(), `"`)}
	return where
}

func (v *Visitor) VisitOperate(ctx *parser.OperateContext) interface{} {
	fmt.Println("operate ", ctx.GetText())
	return strings.Trim(ctx.GetText(), `"`)
}

func (v *Visitor) VisitTest(ctx *parser.TestContext) interface{} {
	fmt.Println("test ", ctx.GetText())
	return strings.Trim(ctx.GetText(), `"`)
}

func parsewhere(where string) Where {
	is := antlr.NewInputStream(where)

	// Create the Lexer
	lexer := parser.NewPREFIXLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewPREFIXParser(tokens)

	v := &Visitor{}
	//Start is rule name of Calc.g4
	a := p.Expression().Accept(v).(Where)
	printwhere(&a, 0)
	return a
}

func printwhere(where *Where, n int) {
	fmt.Println(strings.Repeat("\t", n), where.Operate, where.Test, where.Sql)
	for _, value := range where.Itemlist {
		printwhere(&value, n+1)
	}
}

func test() {
	//where := parsewhere(`(and, true, "a = 123", (atom, "item!=aaa", "b in [123, 354 ]"),  (or, "iftestor ", (atom, "iftestor1", "testor1"), "testor2"))`)
	//where := parsewhere(`(and, true, "a = 123", (atom, "item!=aaa", "b in [123, 354 ]"),  (or, true, "testor1", "testor2"))`)
	//where := parsewhere(`(and, true, "a = 123", "test '' 'good'")`)
	//where := parsewhere(`"a betewen in"`)

	where := parsewhere(`(and, true, "a = 123", (atom, "if item!=aaa", "b in ['123', 354 ]"),  (or, "iftestor ", (atom, "iftestor1", "testor1"), (and, "if orand", (atom, "testfff", "testfff != zz"), "testfff3"), "testor2"))`)
	fmt.Println("result is ", isallatom(where))
	fmt.Println("-------------------------------------")
	genereatesql(where)
	pongo2.SetAutoescape(false)
	tplstruct, err := pongo2.FromFile("./tpl/where.tpl")
	if err != nil {
		fmt.Println(err)
		return
	}

	strdata, err := tplstruct.Execute(pongo2.Context{
		"where": where,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(strdata)

}

func isallatom(where Where) bool {
	if where.Test != "true" {
		return false
	}
	for _, value := range where.Itemlist {
		bret := isallatom(value)
		if bret == false {
			return false
		}
	}
	return true
}

func generateatom(where Where) string {
	var sqls []string
	if len(where.Sql) > 0 {
		return "(" + strings.Trim(where.Sql, "\"") + ")"
	}
	for _, value := range where.Itemlist {
		sqls = append(sqls, generateatom(value))
	}
	return "(" + strings.Join(sqls, " "+where.Operate+" ") + ")"
}

func genereatesql(where Where) {
	if isallatom(where) {
		fmt.Println("where := ", generateatom(where))
	}
	if true {
		aa := "aa"
		if true {
			aa := "bb"
			fmt.Println(aa)
		}
		fmt.Println(aa)
	}
}
func generategocode(where *Where, n int) {

}

type TestA struct {
	Id int
	X  string
}
type TestB struct {
	Id int
	X  string
}
type TestC struct {
	Id int
	X  string
}

type Testresult struct {
	TestA
	TestB
	TestC
}

type TestBMerge struct {
	TestB `merge:"id"`
	C     []*TestC `json:"privis"`
}

type TestMerge struct {
	TestA `merge:"id"`
	B     []*TestBMerge `merge:"" json:"roles"`
}

func init() {
	x := []Testresult{
		{TestA{Id: 1, X: "x"}, TestB{2, "xx"}, TestC{3, "xxx"}},
		{TestA{Id: 1, X: "x"}, TestB{2, "xx"}, TestC{4, "xxx"}},
		{TestA{Id: 2, X: "x"}, TestB{2, "xx"}, TestC{3, "xxx"}},
		{TestA{Id: 3, X: "x"}, TestB{2, "xx"}, TestC{2, "xxx"}},
		{TestA{Id: 3, X: "x"}, TestB{2, "xx"}, TestC{3, "xxx"}},
		{TestA{Id: 3, X: "x"}, TestB{3, "xx"}, TestC{3, "xxx"}},
	}
	y := Groupby(x)
	/*for _, value := range y {
		fmt.Println(value.TestA)
		for _, val := range value.B {
			fmt.Println("\t", val.TestB)
			for _, val2 := range val.C {
				fmt.Println("\t\t", val2.Id, val2.X)
			}
		}
	}*/
	data, _ := json.MarshalIndent(y, "", "   ")
	fmt.Printf(string(data))
}

func Groupby(testresult []Testresult) (merge []*TestMerge) {
	for _, valres := range testresult {
		var m0 *TestMerge
		for _, valmerg := range merge {
			if valmerg.TestA.Id == valres.TestA.Id {
				m0 = valmerg
				break
			}
		}
		if m0 == nil {
			m0 = &TestMerge{}
			merge = append(merge, m0)
		}

		m0.TestA = valres.TestA

		var m1 *TestBMerge
		for _, value := range m0.B {
			if value.TestB.Id == valres.TestB.Id {
				m1 = value
				break
			}
		}
		if m1 == nil {
			m1 = &TestBMerge{}
			m0.B = append(m0.B, m1)
		}
		m1.TestB = valres.TestB

		var m2 *TestC
		for _, value := range m1.C {
			if valres.TestC.Id == value.Id {
				m2 = value
			}
		}
		if m2 == nil {
			m2 = &TestC{}
			m1.C = append(m1.C, m2)
		}
		*m2 = valres.TestC
	}
	return
}
