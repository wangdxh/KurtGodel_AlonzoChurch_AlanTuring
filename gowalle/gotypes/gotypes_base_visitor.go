// Code generated from ./g4/GOTYPES.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gotypes // GOTYPES
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGOTYPESVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGOTYPESVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitFuncdef(ctx *FuncdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitArgs(ctx *ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitRettyps(ctx *RettypsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitStructdef(ctx *StructdefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitStructitem(ctx *StructitemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitTag(ctx *TagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGOTYPESVisitor) VisitTagitem(ctx *TagitemContext) interface{} {
	return v.VisitChildren(ctx)
}
