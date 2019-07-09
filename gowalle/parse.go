package main

import (
	"./parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	where := Where{Operate: "atom", Test: ctx.Test().GetText(), Sql: ctx.CSTRING().GetText()}
	return where
}

func (v *Visitor) VisitExpstring(ctx *parser.ExpstringContext) interface{} {
	fmt.Println("string ", ctx.GetText())
	where := Where{Operate: "atom", Test: "true", Sql: ctx.GetText()}
	return where
}

func (v *Visitor) VisitOperate(ctx *parser.OperateContext) interface{} {
	fmt.Println("operate ", ctx.GetText())
	return ctx.GetText()
}

func (v *Visitor) VisitTest(ctx *parser.TestContext) interface{} {
	fmt.Println("test ", ctx.GetText())
	return ctx.GetText()
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

func isallatom(where *Where) bool {
	if where.Test != "true" {
		return false
	}
	for _, value := range where.Itemlist {
		bret := isallatom(&value)
		if bret == false {
			return false
		}
	}
	return true
}
func test() {
	_ = `(and, true, "a = 123", (atom, true, "b in [123, 354 ]"),  (or, true, "testor1", "testor2"))`
	where := parsewhere(`(and, true, "a = 123", "test good")`)
	fmt.Println("result is ", isallatom(&where))
}

func genereatesql() {

}
func generategocode(where *Where, n int) {

}
