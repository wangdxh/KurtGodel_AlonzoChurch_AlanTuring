// Code generated from ./PREFIX.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PREFIX

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PREFIXParser.
type PREFIXVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PREFIXParser#expstring.
	VisitExpstring(ctx *ExpstringContext) interface{}

	// Visit a parse tree produced by PREFIXParser#expsignle.
	VisitExpsignle(ctx *ExpsignleContext) interface{}

	// Visit a parse tree produced by PREFIXParser#express.
	VisitExpress(ctx *ExpressContext) interface{}

	// Visit a parse tree produced by PREFIXParser#operate.
	VisitOperate(ctx *OperateContext) interface{}

	// Visit a parse tree produced by PREFIXParser#test.
	VisitTest(ctx *TestContext) interface{}
}
