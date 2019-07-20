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

type TestMerge struct {
	TestA `groupby:"id"`
	B     []TestBMerge `groupby:"" json:"roles"` //when we get the list will go into
}

type TestBMerge struct {
	TestB `groupby:"id"`
	C     []TestC `json:"privis"`
}

func init() {
	x := []Testresult{
		{TestA{Id: 1, X: "x"}, TestB{2, "xxb2"}, TestC{1, "xxxc1"}},
		{TestA{Id: 1, X: "x"}, TestB{2, "xxb2"}, TestC{2, "xxxc2"}},
		{TestA{Id: 2, X: "x"}, TestB{2, "xxb2"}, TestC{3, "xxxc3"}},
		{TestA{Id: 3, X: "x"}, TestB{2, "xxb2"}, TestC{4, "xxxc4"}},
		{TestA{Id: 3, X: "x"}, TestB{2, "xxb2"}, TestC{5, "xxxc5"}},
		{TestA{Id: 3, X: "x"}, TestB{3, "xxb3"}, TestC{6, "xxxc6"}},
	}
	y := Groupby(x)

	data, _ := json.MarshalIndent(y, "", "   ")
	fmt.Printf(string(data))
}

func Groupby(testresult []Testresult) (merge []TestMerge) {
	for _, valres := range testresult {
		var m0 *TestMerge
		for inx := 0; inx < len(merge); inx++ {
			if merge[inx].TestA.Id == valres.TestA.Id {
				m0 = &merge[inx]
				break
			}
		}
		if m0 == nil {
			merge = append(merge, TestMerge{})
			m0 = &merge[len(merge)-1]

			m0.TestA = valres.TestA
		}

		var m1 *TestBMerge
		for inx := 0; inx < len(m0.B); inx++ {
			if m0.B[inx].TestB.Id == valres.TestB.Id {
				m1 = &m0.B[inx]
				break
			}
		}

		if m1 == nil {
			m0.B = append(m0.B, TestBMerge{})
			m1 = &m0.B[len(m0.B)-1]

			m1.TestB = valres.TestB
		}

		var m2 *TestC
		for inx := 0; inx < len(m1.C); inx++ {
			if valres.TestC.Id == m1.C[inx].Id {
				m2 = &m1.C[inx]
				break
			}
		}
		if m2 == nil {
			m1.C = append(m1.C, TestC{})
			m2 = &m1.C[len(m1.C)-1]

			*m2 = valres.TestC
		}

	}
	return
}
