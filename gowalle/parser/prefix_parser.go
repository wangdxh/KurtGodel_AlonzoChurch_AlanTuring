// Code generated from ./PREFIX.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // PREFIX

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 41, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 7, 2, 28, 10, 2, 12, 2, 14, 2, 31, 11, 2, 3, 2, 3, 2, 5, 2, 35, 10,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 4, 4, 2, 4, 4, 7,
	8, 4, 2, 9, 9, 11, 11, 2, 40, 2, 34, 3, 2, 2, 2, 4, 36, 3, 2, 2, 2, 6,
	38, 3, 2, 2, 2, 8, 35, 7, 11, 2, 2, 9, 10, 7, 3, 2, 2, 10, 11, 7, 4, 2,
	2, 11, 12, 7, 5, 2, 2, 12, 13, 5, 6, 4, 2, 13, 14, 7, 5, 2, 2, 14, 15,
	7, 11, 2, 2, 15, 16, 7, 6, 2, 2, 16, 35, 3, 2, 2, 2, 17, 18, 7, 3, 2, 2,
	18, 19, 5, 4, 3, 2, 19, 20, 7, 5, 2, 2, 20, 21, 5, 6, 4, 2, 21, 22, 7,
	5, 2, 2, 22, 23, 5, 2, 2, 2, 23, 24, 7, 5, 2, 2, 24, 29, 5, 2, 2, 2, 25,
	26, 7, 5, 2, 2, 26, 28, 5, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 31, 3, 2, 2,
	2, 29, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 32, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 32, 33, 7, 6, 2, 2, 33, 35, 3, 2, 2, 2, 34, 8, 3, 2, 2, 2,
	34, 9, 3, 2, 2, 2, 34, 17, 3, 2, 2, 2, 35, 3, 3, 2, 2, 2, 36, 37, 9, 2,
	2, 2, 37, 5, 3, 2, 2, 2, 38, 39, 9, 3, 2, 2, 39, 7, 3, 2, 2, 2, 4, 29,
	34,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "'atom'", "','", "')'", "'and'", "'or'", "'true'", "", "", "'''",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "WS", "CSTRING", "APOSTROPHE",
}

var ruleNames = []string{
	"expression", "operate", "test",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type PREFIXParser struct {
	*antlr.BaseParser
}

func NewPREFIXParser(input antlr.TokenStream) *PREFIXParser {
	this := new(PREFIXParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "PREFIX.g4"

	return this
}

// PREFIXParser tokens.
const (
	PREFIXParserEOF        = antlr.TokenEOF
	PREFIXParserT__0       = 1
	PREFIXParserT__1       = 2
	PREFIXParserT__2       = 3
	PREFIXParserT__3       = 4
	PREFIXParserT__4       = 5
	PREFIXParserT__5       = 6
	PREFIXParserT__6       = 7
	PREFIXParserWS         = 8
	PREFIXParserCSTRING    = 9
	PREFIXParserAPOSTROPHE = 10
)

// PREFIXParser rules.
const (
	PREFIXParserRULE_expression = 0
	PREFIXParserRULE_operate    = 1
	PREFIXParserRULE_test       = 2
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PREFIXParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PREFIXParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpstringContext struct {
	*ExpressionContext
}

func NewExpstringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpstringContext {
	var p = new(ExpstringContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpstringContext) CSTRING() antlr.TerminalNode {
	return s.GetToken(PREFIXParserCSTRING, 0)
}

func (s *ExpstringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PREFIXVisitor:
		return t.VisitExpstring(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressContext struct {
	*ExpressionContext
}

func NewExpressContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressContext {
	var p = new(ExpressContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressContext) Operate() IOperateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperateContext)
}

func (s *ExpressContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExpressContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PREFIXVisitor:
		return t.VisitExpress(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpsignleContext struct {
	*ExpressionContext
}

func NewExpsignleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpsignleContext {
	var p = new(ExpsignleContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpsignleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpsignleContext) Test() ITestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestContext)
}

func (s *ExpsignleContext) CSTRING() antlr.TerminalNode {
	return s.GetToken(PREFIXParserCSTRING, 0)
}

func (s *ExpsignleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PREFIXVisitor:
		return t.VisitExpsignle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PREFIXParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PREFIXParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpstringContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(6)
			p.Match(PREFIXParserCSTRING)
		}

	case 2:
		localctx = NewExpsignleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(7)
			p.Match(PREFIXParserT__0)
		}
		{
			p.SetState(8)
			p.Match(PREFIXParserT__1)
		}
		{
			p.SetState(9)
			p.Match(PREFIXParserT__2)
		}
		{
			p.SetState(10)
			p.Test()
		}
		{
			p.SetState(11)
			p.Match(PREFIXParserT__2)
		}
		{
			p.SetState(12)
			p.Match(PREFIXParserCSTRING)
		}
		{
			p.SetState(13)
			p.Match(PREFIXParserT__3)
		}

	case 3:
		localctx = NewExpressContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(15)
			p.Match(PREFIXParserT__0)
		}
		{
			p.SetState(16)
			p.Operate()
		}
		{
			p.SetState(17)
			p.Match(PREFIXParserT__2)
		}
		{
			p.SetState(18)
			p.Test()
		}
		{
			p.SetState(19)
			p.Match(PREFIXParserT__2)
		}
		{
			p.SetState(20)
			p.Expression()
		}
		{
			p.SetState(21)
			p.Match(PREFIXParserT__2)
		}
		{
			p.SetState(22)
			p.Expression()
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PREFIXParserT__2 {
			{
				p.SetState(23)
				p.Match(PREFIXParserT__2)
			}
			{
				p.SetState(24)
				p.Expression()
			}

			p.SetState(29)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(30)
			p.Match(PREFIXParserT__3)
		}

	}

	return localctx
}

// IOperateContext is an interface to support dynamic dispatch.
type IOperateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperateContext differentiates from other interfaces.
	IsOperateContext()
}

type OperateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperateContext() *OperateContext {
	var p = new(OperateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PREFIXParserRULE_operate
	return p
}

func (*OperateContext) IsOperateContext() {}

func NewOperateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperateContext {
	var p = new(OperateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PREFIXParserRULE_operate

	return p
}

func (s *OperateContext) GetParser() antlr.Parser { return s.parser }
func (s *OperateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PREFIXVisitor:
		return t.VisitOperate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PREFIXParser) Operate() (localctx IOperateContext) {
	localctx = NewOperateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PREFIXParserRULE_operate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PREFIXParserT__1)|(1<<PREFIXParserT__4)|(1<<PREFIXParserT__5))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITestContext is an interface to support dynamic dispatch.
type ITestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestContext differentiates from other interfaces.
	IsTestContext()
}

type TestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestContext() *TestContext {
	var p = new(TestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PREFIXParserRULE_test
	return p
}

func (*TestContext) IsTestContext() {}

func NewTestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestContext {
	var p = new(TestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PREFIXParserRULE_test

	return p
}

func (s *TestContext) GetParser() antlr.Parser { return s.parser }

func (s *TestContext) CSTRING() antlr.TerminalNode {
	return s.GetToken(PREFIXParserCSTRING, 0)
}

func (s *TestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PREFIXVisitor:
		return t.VisitTest(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PREFIXParser) Test() (localctx ITestContext) {
	localctx = NewTestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PREFIXParserRULE_test)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PREFIXParserT__6 || _la == PREFIXParserCSTRING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
