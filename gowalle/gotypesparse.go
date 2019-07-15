package main

import (
	"./gotypes"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"io/ioutil"
	"strings"
)

type GoTypesVisitor struct {
	gotypes.BaseGOTYPESVisitor
	Defs GoDefs
}

type GoDefs struct {
	Structlists []Struct
	Funclists   []Func
}

type Func struct {
	Name     string
	RetType  string
	Arglists []Argitem
}

type Argitem struct {
	Name      string
	Type      string
	Islist    bool
	Ispointer bool
}

func (v *GoTypesVisitor) visitRule(node antlr.RuleNode) interface{} {
	return node.Accept(v)
}

func (v *GoTypesVisitor) VisitProg(ctx *gotypes.ProgContext) interface{} {
	for _, value := range ctx.AllExpression() {
		v.visitRule(value)
	}

	for _, val := range v.Defs.Funclists {
		for arginx, value := range val.Arglists {
			if len(value.Type) == 0 {
				for temp := arginx + 1; temp < len(val.Arglists); temp++ {
					if len(val.Arglists[temp].Type) != 0 {
						val.Arglists[arginx].Type = val.Arglists[temp].Type
						val.Arglists[arginx].Ispointer = val.Arglists[temp].Ispointer
						val.Arglists[arginx].Islist = val.Arglists[temp].Islist
						break
					}
				}
			}
		}
	}
	return nil
}

func (v *GoTypesVisitor) VisitExpression(ctx *gotypes.ExpressionContext) interface{} {
	if ctx.Structdef() != nil {
		v.Defs.Structlists = append(v.Defs.Structlists, v.visitRule(ctx.Structdef()).(Struct))
	}
	if ctx.Funcdef() != nil {
		v.Defs.Funclists = append(v.Defs.Funclists, v.visitRule(ctx.Funcdef()).(Func))
	}
	return v.Defs
}

func (v *GoTypesVisitor) VisitFuncdef(ctx *gotypes.FuncdefContext) interface{} {
	item := Func{}
	item.Name = ctx.IDENTIFIER(0).GetText()
	item.RetType = ctx.IDENTIFIER(1).GetText()
	for _, value := range ctx.AllArgs() {
		item.Arglists = append(item.Arglists, v.visitRule(value).(Argitem))
	}
	return item
}

func (v *GoTypesVisitor) VisitArgs(ctx *gotypes.ArgsContext) interface{} {
	item := Argitem{}
	itemlist := ctx.AllIDENTIFIER()
	if len(itemlist) == 1 {
		item.Name = ctx.IDENTIFIER(0).GetText()
	} else if len(itemlist) == 2 {
		item.Name = ctx.IDENTIFIER(0).GetText()
		item.Type = ctx.IDENTIFIER(1).GetText()
	}
	if ctx.BRACKETS() != nil {
		item.Islist = true
	}
	if ctx.POINTER() != nil {
		item.Ispointer = true
	}
	return item
}

type Struct struct {
	Name     string
	ItemList []StructItem
}

func (v *GoTypesVisitor) VisitStructdef(ctx *gotypes.StructdefContext) interface{} {
	one := Struct{}
	one.Name = ctx.IDENTIFIER().GetText()
	for _, value := range ctx.AllStructitem() {
		one.ItemList = append(one.ItemList, v.visitRule(value).(StructItem))
	}
	return one
}

type StructItem struct {
	Name      string
	Type      string
	Islist    bool
	Ispointer bool
	Tags      []Tagitem
}

func (v *GoTypesVisitor) VisitStructitem(ctx *gotypes.StructitemContext) interface{} {
	item := StructItem{}
	ids := ctx.AllIDENTIFIER()
	if len(ids) == 2 {
		item.Name = ctx.IDENTIFIER(0).GetText()
		item.Type = ctx.IDENTIFIER(1).GetText()
	} else if len(ids) == 1 {
		item.Name = ctx.IDENTIFIER(0).GetText()
		item.Type = ctx.IDENTIFIER(0).GetText()
	}
	if ctx.POINTER() != nil {
		item.Ispointer = true
	}

	if ctx.BRACKETS() != nil {
		item.Islist = true
	}

	if ctx.Tag() != nil {
		item.Tags = v.visitRule(ctx.Tag()).([]Tagitem)
	}
	return item
}

func (v *GoTypesVisitor) VisitTag(ctx *gotypes.TagContext) interface{} {
	var tagitemlist []Tagitem
	for _, value := range ctx.AllTagitem() {
		tagitem := v.visitRule(value).(Tagitem)
		tagitemlist = append(tagitemlist, tagitem)
	}
	return tagitemlist
}

type Tagitem struct {
	Name  string
	Value string
}

func (v *GoTypesVisitor) VisitTagitem(ctx *gotypes.TagitemContext) interface{} {
	return Tagitem{ctx.IDENTIFIER().GetText(), strings.Trim(ctx.CSTRING().GetText(), `"`)}
}

func Parsestruct(strdef string) GoDefs {
	is := antlr.NewInputStream(strdef)

	// Create the Lexer
	lexer := gotypes.NewGOTYPESLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := gotypes.NewGOTYPESParser(tokens)
	p.AddErrorListener(antlr.ConsoleErrorListenerINSTANCE)

	v := &GoTypesVisitor{}
	v.visitRule(p.Prog())
	return v.Defs
}

func init() {
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
	}
	v := Parsestruct(string(data))
	fmt.Println(v)
	/*for _, value := range v {
		fmt.Println(value)
	}*/
}
