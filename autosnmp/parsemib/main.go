package main

import (
	"./parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"fmt"
)

 var g_mapnametooid map[string]string = map[string]string{
"mgmt" : ".1.3.6.1.2",
}

func main()  {
	data , err := ioutil.ReadFile(`D:\github\kill-that-programmer\autosnmp\parsemib/mibs/RFC1213-MIB`)
	if err!=nil {
		fmt.Println(err)
		return
	}

	fmt.Println("---------------------1")
	input := antlr.NewInputStream(string(data))
	lexer := parser.NewASNLexer(input)
	token:=antlr.NewCommonTokenStream(lexer, 0)
	parser :=parser.NewASNParser(token)
	fmt.Println("---------------------2")
	myvisitor := Myvisitor{}
	myvisitor.Visit(parser.ModuleDefinition())
}

type Myvisitor struct {
	parser.BaseASNVisitor
}

func (v *Myvisitor) VisitModuleDefinition(ctx *parser.ModuleDefinitionContext) interface{} {
	fmt.Println("moduledefi")

	identilist := ctx.AllIDENTIFIER()
	fmt.Println(len(identilist))
	for key, val := range identilist {
		fmt.Println(key, val.GetText())
	}
	return v.VisitChildren(ctx)
}