// Code generated from ./PREFIX.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PREFIX

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasePREFIXVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasePREFIXVisitor) VisitExpstring(ctx *ExpstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePREFIXVisitor) VisitExpsignle(ctx *ExpsignleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePREFIXVisitor) VisitExpress(ctx *ExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePREFIXVisitor) VisitOperate(ctx *OperateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasePREFIXVisitor) VisitTest(ctx *TestContext) interface{} {
	return v.VisitChildren(ctx)
}
