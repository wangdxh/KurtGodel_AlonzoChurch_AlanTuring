// Code generated from ./g4/GOTYPES.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gotypes // GOTYPES
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GOTYPESParser.
type GOTYPESVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GOTYPESParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#funcdef.
	VisitFuncdef(ctx *FuncdefContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#args.
	VisitArgs(ctx *ArgsContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#rettyps.
	VisitRettyps(ctx *RettypsContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#structdef.
	VisitStructdef(ctx *StructdefContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#structitem.
	VisitStructitem(ctx *StructitemContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#tag.
	VisitTag(ctx *TagContext) interface{}

	// Visit a parse tree produced by GOTYPESParser#tagitem.
	VisitTagitem(ctx *TagitemContext) interface{}
}
