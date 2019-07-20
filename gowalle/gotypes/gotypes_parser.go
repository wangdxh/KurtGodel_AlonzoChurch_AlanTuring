// Code generated from ./g4/GOTYPES.g4 by ANTLR 4.7.1. DO NOT EDIT.

package gotypes // GOTYPES
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 20, 114,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 6, 2, 22, 10, 2, 13, 2, 14,
	2, 23, 3, 3, 3, 3, 5, 3, 28, 10, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 3,
	5, 3, 35, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 41, 10, 4, 3, 4, 3, 4, 7,
	4, 45, 10, 4, 12, 4, 14, 4, 48, 11, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5,
	5, 55, 10, 5, 3, 5, 5, 5, 58, 10, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 5, 6,
	64, 10, 6, 3, 6, 5, 6, 67, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 7, 7, 78, 10, 7, 12, 7, 14, 7, 81, 11, 7, 6, 7, 83, 10,
	7, 13, 7, 14, 7, 84, 3, 7, 3, 7, 3, 8, 5, 8, 90, 10, 8, 3, 8, 5, 8, 93,
	10, 8, 3, 8, 5, 8, 96, 10, 8, 3, 8, 3, 8, 5, 8, 100, 10, 8, 3, 9, 3, 9,
	6, 9, 104, 10, 9, 13, 9, 14, 9, 105, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 2, 2, 11, 2, 4, 6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 123, 2, 21,
	3, 2, 2, 2, 4, 34, 3, 2, 2, 2, 6, 36, 3, 2, 2, 2, 8, 52, 3, 2, 2, 2, 10,
	63, 3, 2, 2, 2, 12, 70, 3, 2, 2, 2, 14, 89, 3, 2, 2, 2, 16, 101, 3, 2,
	2, 2, 18, 109, 3, 2, 2, 2, 20, 22, 5, 4, 3, 2, 21, 20, 3, 2, 2, 2, 22,
	23, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 3, 3, 2, 2,
	2, 25, 27, 5, 12, 7, 2, 26, 28, 7, 19, 2, 2, 27, 26, 3, 2, 2, 2, 27, 28,
	3, 2, 2, 2, 28, 35, 3, 2, 2, 2, 29, 31, 5, 6, 4, 2, 30, 32, 7, 19, 2, 2,
	31, 30, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33, 35, 7,
	19, 2, 2, 34, 25, 3, 2, 2, 2, 34, 29, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35,
	5, 3, 2, 2, 2, 36, 37, 7, 3, 2, 2, 37, 38, 7, 18, 2, 2, 38, 40, 7, 4, 2,
	2, 39, 41, 5, 8, 5, 2, 40, 39, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 46,
	3, 2, 2, 2, 42, 43, 7, 5, 2, 2, 43, 45, 5, 8, 5, 2, 44, 42, 3, 2, 2, 2,
	45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3,
	2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 50, 7, 6, 2, 2, 50, 51, 5, 10, 6, 2, 51,
	7, 3, 2, 2, 2, 52, 54, 7, 18, 2, 2, 53, 55, 7, 14, 2, 2, 54, 53, 3, 2,
	2, 2, 54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 58, 7, 13, 2, 2, 57,
	56, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59, 61, 7, 18,
	2, 2, 60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 64,
	7, 14, 2, 2, 63, 62, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2,
	65, 67, 7, 13, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3,
	2, 2, 2, 68, 69, 7, 18, 2, 2, 69, 11, 3, 2, 2, 2, 70, 71, 7, 7, 2, 2, 71,
	72, 7, 18, 2, 2, 72, 73, 7, 8, 2, 2, 73, 74, 7, 9, 2, 2, 74, 82, 7, 19,
	2, 2, 75, 79, 5, 14, 8, 2, 76, 78, 7, 19, 2, 2, 77, 76, 3, 2, 2, 2, 78,
	81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 83, 3, 2, 2,
	2, 81, 79, 3, 2, 2, 2, 82, 75, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 82,
	3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 87, 7, 10, 2, 2,
	87, 13, 3, 2, 2, 2, 88, 90, 7, 18, 2, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 93, 7, 14, 2, 2, 92, 91, 3, 2, 2, 2, 92,
	93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 96, 7, 13, 2, 2, 95, 94, 3, 2,
	2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 7, 18, 2, 2, 98,
	100, 5, 16, 9, 2, 99, 98, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 15, 3,
	2, 2, 2, 101, 103, 7, 11, 2, 2, 102, 104, 5, 18, 10, 2, 103, 102, 3, 2,
	2, 2, 104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 108, 7, 11, 2, 2, 108, 17, 3, 2, 2, 2, 109,
	110, 7, 18, 2, 2, 110, 111, 7, 12, 2, 2, 111, 112, 7, 16, 2, 2, 112, 19,
	3, 2, 2, 2, 20, 23, 27, 31, 34, 40, 46, 54, 57, 60, 63, 66, 79, 84, 89,
	92, 95, 99, 105,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "','", "')'", "'type'", "'struct'", "'{'", "'}'",
	"'`'", "':'", "'*'", "", "", "", "'''",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "POINTER", "BRACKETS", "WS",
	"CSTRING", "APOSTROPHE", "IDENTIFIER", "NEWLINE", "COMMENT",
}

var ruleNames = []string{
	"prog", "expression", "funcdef", "args", "rettyps", "structdef", "structitem",
	"tag", "tagitem",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GOTYPESParser struct {
	*antlr.BaseParser
}

func NewGOTYPESParser(input antlr.TokenStream) *GOTYPESParser {
	this := new(GOTYPESParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GOTYPES.g4"

	return this
}

// GOTYPESParser tokens.
const (
	GOTYPESParserEOF        = antlr.TokenEOF
	GOTYPESParserT__0       = 1
	GOTYPESParserT__1       = 2
	GOTYPESParserT__2       = 3
	GOTYPESParserT__3       = 4
	GOTYPESParserT__4       = 5
	GOTYPESParserT__5       = 6
	GOTYPESParserT__6       = 7
	GOTYPESParserT__7       = 8
	GOTYPESParserT__8       = 9
	GOTYPESParserT__9       = 10
	GOTYPESParserPOINTER    = 11
	GOTYPESParserBRACKETS   = 12
	GOTYPESParserWS         = 13
	GOTYPESParserCSTRING    = 14
	GOTYPESParserAPOSTROPHE = 15
	GOTYPESParserIDENTIFIER = 16
	GOTYPESParserNEWLINE    = 17
	GOTYPESParserCOMMENT    = 18
)

// GOTYPESParser rules.
const (
	GOTYPESParserRULE_prog       = 0
	GOTYPESParserRULE_expression = 1
	GOTYPESParserRULE_funcdef    = 2
	GOTYPESParserRULE_args       = 3
	GOTYPESParserRULE_rettyps    = 4
	GOTYPESParserRULE_structdef  = 5
	GOTYPESParserRULE_structitem = 6
	GOTYPESParserRULE_tag        = 7
	GOTYPESParserRULE_tagitem    = 8
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ProgContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GOTYPESParserRULE_prog)
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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GOTYPESParserT__0)|(1<<GOTYPESParserT__4)|(1<<GOTYPESParserNEWLINE))) != 0) {
		{
			p.SetState(18)
			p.Expression()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

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
	p.RuleIndex = GOTYPESParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Structdef() IStructdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructdefContext)
}

func (s *ExpressionContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserNEWLINE, 0)
}

func (s *ExpressionContext) Funcdef() IFuncdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncdefContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GOTYPESParserRULE_expression)

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

	switch p.GetTokenStream().LA(1) {
	case GOTYPESParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Structdef()
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(24)
				p.Match(GOTYPESParserNEWLINE)
			}

		}

	case GOTYPESParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Funcdef()
		}
		p.SetState(29)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(28)
				p.Match(GOTYPESParserNEWLINE)
			}

		}

	case GOTYPESParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(31)
			p.Match(GOTYPESParserNEWLINE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncdefContext is an interface to support dynamic dispatch.
type IFuncdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncdefContext differentiates from other interfaces.
	IsFuncdefContext()
}

type FuncdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncdefContext() *FuncdefContext {
	var p = new(FuncdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_funcdef
	return p
}

func (*FuncdefContext) IsFuncdefContext() {}

func NewFuncdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncdefContext {
	var p = new(FuncdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_funcdef

	return p
}

func (s *FuncdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncdefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, 0)
}

func (s *FuncdefContext) Rettyps() IRettypsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRettypsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRettypsContext)
}

func (s *FuncdefContext) AllArgs() []IArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgsContext)(nil)).Elem())
	var tst = make([]IArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgsContext)
		}
	}

	return tst
}

func (s *FuncdefContext) Args(i int) IArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgsContext)
}

func (s *FuncdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitFuncdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Funcdef() (localctx IFuncdefContext) {
	localctx = NewFuncdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GOTYPESParserRULE_funcdef)
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
		p.Match(GOTYPESParserT__0)
	}
	{
		p.SetState(35)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(36)
		p.Match(GOTYPESParserT__1)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(37)
			p.Args()
		}

	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GOTYPESParserT__2 {
		{
			p.SetState(40)
			p.Match(GOTYPESParserT__2)
		}
		{
			p.SetState(41)
			p.Args()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(GOTYPESParserT__3)
	}
	{
		p.SetState(48)
		p.Rettyps()
	}

	return localctx
}

// IArgsContext is an interface to support dynamic dispatch.
type IArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgsContext differentiates from other interfaces.
	IsArgsContext()
}

type ArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgsContext() *ArgsContext {
	var p = new(ArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_args
	return p
}

func (*ArgsContext) IsArgsContext() {}

func NewArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgsContext {
	var p = new(ArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_args

	return p
}

func (s *ArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgsContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GOTYPESParserIDENTIFIER)
}

func (s *ArgsContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, i)
}

func (s *ArgsContext) BRACKETS() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserBRACKETS, 0)
}

func (s *ArgsContext) POINTER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserPOINTER, 0)
}

func (s *ArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Args() (localctx IArgsContext) {
	localctx = NewArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GOTYPESParserRULE_args)
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
		p.SetState(50)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserBRACKETS {
		{
			p.SetState(51)
			p.Match(GOTYPESParserBRACKETS)
		}

	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserPOINTER {
		{
			p.SetState(54)
			p.Match(GOTYPESParserPOINTER)
		}

	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(57)
			p.Match(GOTYPESParserIDENTIFIER)
		}

	}

	return localctx
}

// IRettypsContext is an interface to support dynamic dispatch.
type IRettypsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRettypsContext differentiates from other interfaces.
	IsRettypsContext()
}

type RettypsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRettypsContext() *RettypsContext {
	var p = new(RettypsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_rettyps
	return p
}

func (*RettypsContext) IsRettypsContext() {}

func NewRettypsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RettypsContext {
	var p = new(RettypsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_rettyps

	return p
}

func (s *RettypsContext) GetParser() antlr.Parser { return s.parser }

func (s *RettypsContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, 0)
}

func (s *RettypsContext) BRACKETS() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserBRACKETS, 0)
}

func (s *RettypsContext) POINTER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserPOINTER, 0)
}

func (s *RettypsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RettypsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RettypsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitRettyps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Rettyps() (localctx IRettypsContext) {
	localctx = NewRettypsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GOTYPESParserRULE_rettyps)
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserBRACKETS {
		{
			p.SetState(60)
			p.Match(GOTYPESParserBRACKETS)
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserPOINTER {
		{
			p.SetState(63)
			p.Match(GOTYPESParserPOINTER)
		}

	}
	{
		p.SetState(66)
		p.Match(GOTYPESParserIDENTIFIER)
	}

	return localctx
}

// IStructdefContext is an interface to support dynamic dispatch.
type IStructdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructdefContext differentiates from other interfaces.
	IsStructdefContext()
}

type StructdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructdefContext() *StructdefContext {
	var p = new(StructdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_structdef
	return p
}

func (*StructdefContext) IsStructdefContext() {}

func NewStructdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructdefContext {
	var p = new(StructdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_structdef

	return p
}

func (s *StructdefContext) GetParser() antlr.Parser { return s.parser }

func (s *StructdefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, 0)
}

func (s *StructdefContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(GOTYPESParserNEWLINE)
}

func (s *StructdefContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(GOTYPESParserNEWLINE, i)
}

func (s *StructdefContext) AllStructitem() []IStructitemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructitemContext)(nil)).Elem())
	var tst = make([]IStructitemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructitemContext)
		}
	}

	return tst
}

func (s *StructdefContext) Structitem(i int) IStructitemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructitemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructitemContext)
}

func (s *StructdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructdefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitStructdef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Structdef() (localctx IStructdefContext) {
	localctx = NewStructdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GOTYPESParserRULE_structdef)
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
		p.SetState(68)
		p.Match(GOTYPESParserT__4)
	}
	{
		p.SetState(69)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(70)
		p.Match(GOTYPESParserT__5)
	}
	{
		p.SetState(71)
		p.Match(GOTYPESParserT__6)
	}
	{
		p.SetState(72)
		p.Match(GOTYPESParserNEWLINE)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GOTYPESParserPOINTER)|(1<<GOTYPESParserBRACKETS)|(1<<GOTYPESParserIDENTIFIER))) != 0) {
		{
			p.SetState(73)
			p.Structitem()
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GOTYPESParserNEWLINE {
			{
				p.SetState(74)
				p.Match(GOTYPESParserNEWLINE)
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(84)
		p.Match(GOTYPESParserT__7)
	}

	return localctx
}

// IStructitemContext is an interface to support dynamic dispatch.
type IStructitemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructitemContext differentiates from other interfaces.
	IsStructitemContext()
}

type StructitemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructitemContext() *StructitemContext {
	var p = new(StructitemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_structitem
	return p
}

func (*StructitemContext) IsStructitemContext() {}

func NewStructitemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructitemContext {
	var p = new(StructitemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_structitem

	return p
}

func (s *StructitemContext) GetParser() antlr.Parser { return s.parser }

func (s *StructitemContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GOTYPESParserIDENTIFIER)
}

func (s *StructitemContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, i)
}

func (s *StructitemContext) BRACKETS() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserBRACKETS, 0)
}

func (s *StructitemContext) POINTER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserPOINTER, 0)
}

func (s *StructitemContext) Tag() ITagContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITagContext)
}

func (s *StructitemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructitemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructitemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitStructitem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Structitem() (localctx IStructitemContext) {
	localctx = NewStructitemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GOTYPESParserRULE_structitem)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(86)
			p.Match(GOTYPESParserIDENTIFIER)
		}

	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserBRACKETS {
		{
			p.SetState(89)
			p.Match(GOTYPESParserBRACKETS)
		}

	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserPOINTER {
		{
			p.SetState(92)
			p.Match(GOTYPESParserPOINTER)
		}

	}
	{
		p.SetState(95)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserT__8 {
		{
			p.SetState(96)
			p.Tag()
		}

	}

	return localctx
}

// ITagContext is an interface to support dynamic dispatch.
type ITagContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagContext differentiates from other interfaces.
	IsTagContext()
}

type TagContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagContext() *TagContext {
	var p = new(TagContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_tag
	return p
}

func (*TagContext) IsTagContext() {}

func NewTagContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagContext {
	var p = new(TagContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_tag

	return p
}

func (s *TagContext) GetParser() antlr.Parser { return s.parser }

func (s *TagContext) AllTagitem() []ITagitemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITagitemContext)(nil)).Elem())
	var tst = make([]ITagitemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITagitemContext)
		}
	}

	return tst
}

func (s *TagContext) Tagitem(i int) ITagitemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITagitemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITagitemContext)
}

func (s *TagContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitTag(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Tag() (localctx ITagContext) {
	localctx = NewTagContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GOTYPESParserRULE_tag)
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
		p.SetState(99)
		p.Match(GOTYPESParserT__8)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(100)
			p.Tagitem()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
		p.Match(GOTYPESParserT__8)
	}

	return localctx
}

// ITagitemContext is an interface to support dynamic dispatch.
type ITagitemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTagitemContext differentiates from other interfaces.
	IsTagitemContext()
}

type TagitemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTagitemContext() *TagitemContext {
	var p = new(TagitemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GOTYPESParserRULE_tagitem
	return p
}

func (*TagitemContext) IsTagitemContext() {}

func NewTagitemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TagitemContext {
	var p = new(TagitemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GOTYPESParserRULE_tagitem

	return p
}

func (s *TagitemContext) GetParser() antlr.Parser { return s.parser }

func (s *TagitemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, 0)
}

func (s *TagitemContext) CSTRING() antlr.TerminalNode {
	return s.GetToken(GOTYPESParserCSTRING, 0)
}

func (s *TagitemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagitemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TagitemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GOTYPESVisitor:
		return t.VisitTagitem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GOTYPESParser) Tagitem() (localctx ITagitemContext) {
	localctx = NewTagitemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GOTYPESParserRULE_tagitem)

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
		p.SetState(107)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(108)
		p.Match(GOTYPESParserT__9)
	}
	{
		p.SetState(109)
		p.Match(GOTYPESParserCSTRING)
	}

	return localctx
}
