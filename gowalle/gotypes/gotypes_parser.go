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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 3,
	3, 3, 5, 3, 26, 10, 3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 3, 5, 3, 33, 10,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 39, 10, 4, 3, 4, 3, 4, 7, 4, 43, 10, 4,
	12, 4, 14, 4, 46, 11, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 53, 10, 5,
	3, 5, 5, 5, 56, 10, 5, 3, 5, 5, 5, 59, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 7, 6, 68, 10, 6, 12, 6, 14, 6, 71, 11, 6, 6, 6, 73, 10,
	6, 13, 6, 14, 6, 74, 3, 6, 3, 6, 3, 7, 5, 7, 80, 10, 7, 3, 7, 5, 7, 83,
	10, 7, 3, 7, 5, 7, 86, 10, 7, 3, 7, 3, 7, 5, 7, 90, 10, 7, 3, 8, 3, 8,
	6, 8, 94, 10, 8, 13, 8, 14, 8, 95, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 2, 2, 112, 2, 19, 3, 2,
	2, 2, 4, 32, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 50, 3, 2, 2, 2, 10, 60,
	3, 2, 2, 2, 12, 79, 3, 2, 2, 2, 14, 91, 3, 2, 2, 2, 16, 99, 3, 2, 2, 2,
	18, 20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 19, 3,
	2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 3, 3, 2, 2, 2, 23, 25, 5, 10, 6, 2, 24,
	26, 7, 19, 2, 2, 25, 24, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 33, 3, 2,
	2, 2, 27, 29, 5, 6, 4, 2, 28, 30, 7, 19, 2, 2, 29, 28, 3, 2, 2, 2, 29,
	30, 3, 2, 2, 2, 30, 33, 3, 2, 2, 2, 31, 33, 7, 19, 2, 2, 32, 23, 3, 2,
	2, 2, 32, 27, 3, 2, 2, 2, 32, 31, 3, 2, 2, 2, 33, 5, 3, 2, 2, 2, 34, 35,
	7, 3, 2, 2, 35, 36, 7, 18, 2, 2, 36, 38, 7, 4, 2, 2, 37, 39, 5, 8, 5, 2,
	38, 37, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 44, 3, 2, 2, 2, 40, 41, 7,
	5, 2, 2, 41, 43, 5, 8, 5, 2, 42, 40, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44,
	42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2,
	2, 47, 48, 7, 6, 2, 2, 48, 49, 7, 18, 2, 2, 49, 7, 3, 2, 2, 2, 50, 52,
	7, 18, 2, 2, 51, 53, 7, 14, 2, 2, 52, 51, 3, 2, 2, 2, 52, 53, 3, 2, 2,
	2, 53, 55, 3, 2, 2, 2, 54, 56, 7, 13, 2, 2, 55, 54, 3, 2, 2, 2, 55, 56,
	3, 2, 2, 2, 56, 58, 3, 2, 2, 2, 57, 59, 7, 18, 2, 2, 58, 57, 3, 2, 2, 2,
	58, 59, 3, 2, 2, 2, 59, 9, 3, 2, 2, 2, 60, 61, 7, 7, 2, 2, 61, 62, 7, 18,
	2, 2, 62, 63, 7, 8, 2, 2, 63, 64, 7, 9, 2, 2, 64, 72, 7, 19, 2, 2, 65,
	69, 5, 12, 7, 2, 66, 68, 7, 19, 2, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2,
	2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69,
	3, 2, 2, 2, 72, 65, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2,
	74, 75, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 7, 10, 2, 2, 77, 11, 3,
	2, 2, 2, 78, 80, 7, 18, 2, 2, 79, 78, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80,
	82, 3, 2, 2, 2, 81, 83, 7, 14, 2, 2, 82, 81, 3, 2, 2, 2, 82, 83, 3, 2,
	2, 2, 83, 85, 3, 2, 2, 2, 84, 86, 7, 13, 2, 2, 85, 84, 3, 2, 2, 2, 85,
	86, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 89, 7, 18, 2, 2, 88, 90, 5, 14,
	8, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 13, 3, 2, 2, 2, 91, 93,
	7, 11, 2, 2, 92, 94, 5, 16, 9, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2,
	2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98,
	7, 11, 2, 2, 98, 15, 3, 2, 2, 2, 99, 100, 7, 18, 2, 2, 100, 101, 7, 12,
	2, 2, 101, 102, 7, 16, 2, 2, 102, 17, 3, 2, 2, 2, 18, 21, 25, 29, 32, 38,
	44, 52, 55, 58, 69, 74, 79, 82, 85, 89, 95,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "','", "')'", "'type'", "'struct'", "'{'", "'}'",
	"'`'", "':'", "'*'", "", "", "", "'''",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "POINTER", "BRACKETS", "WS",
	"CSTRING", "APOSTROPHE", "IDENTIFIER", "NEWLINE", "LINE_COMMENT", "COMMENT",
}

var ruleNames = []string{
	"prog", "expression", "funcdef", "args", "structdef", "structitem", "tag",
	"tagitem",
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
	GOTYPESParserEOF          = antlr.TokenEOF
	GOTYPESParserT__0         = 1
	GOTYPESParserT__1         = 2
	GOTYPESParserT__2         = 3
	GOTYPESParserT__3         = 4
	GOTYPESParserT__4         = 5
	GOTYPESParserT__5         = 6
	GOTYPESParserT__6         = 7
	GOTYPESParserT__7         = 8
	GOTYPESParserT__8         = 9
	GOTYPESParserT__9         = 10
	GOTYPESParserPOINTER      = 11
	GOTYPESParserBRACKETS     = 12
	GOTYPESParserWS           = 13
	GOTYPESParserCSTRING      = 14
	GOTYPESParserAPOSTROPHE   = 15
	GOTYPESParserIDENTIFIER   = 16
	GOTYPESParserNEWLINE      = 17
	GOTYPESParserLINE_COMMENT = 18
	GOTYPESParserCOMMENT      = 19
)

// GOTYPESParser rules.
const (
	GOTYPESParserRULE_prog       = 0
	GOTYPESParserRULE_expression = 1
	GOTYPESParserRULE_funcdef    = 2
	GOTYPESParserRULE_args       = 3
	GOTYPESParserRULE_structdef  = 4
	GOTYPESParserRULE_structitem = 5
	GOTYPESParserRULE_tag        = 6
	GOTYPESParserRULE_tagitem    = 7
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GOTYPESParserT__0)|(1<<GOTYPESParserT__4)|(1<<GOTYPESParserNEWLINE))) != 0) {
		{
			p.SetState(16)
			p.Expression()
		}

		p.SetState(19)
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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GOTYPESParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Structdef()
		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(22)
				p.Match(GOTYPESParserNEWLINE)
			}

		}

	case GOTYPESParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Funcdef()
		}
		p.SetState(27)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(26)
				p.Match(GOTYPESParserNEWLINE)
			}

		}

	case GOTYPESParserNEWLINE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
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

func (s *FuncdefContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(GOTYPESParserIDENTIFIER)
}

func (s *FuncdefContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(GOTYPESParserIDENTIFIER, i)
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
		p.SetState(32)
		p.Match(GOTYPESParserT__0)
	}
	{
		p.SetState(33)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(34)
		p.Match(GOTYPESParserT__1)
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(35)
			p.Args()
		}

	}
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GOTYPESParserT__2 {
		{
			p.SetState(38)
			p.Match(GOTYPESParserT__2)
		}
		{
			p.SetState(39)
			p.Args()
		}

		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.Match(GOTYPESParserT__3)
	}
	{
		p.SetState(46)
		p.Match(GOTYPESParserIDENTIFIER)
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
		p.SetState(48)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserBRACKETS {
		{
			p.SetState(49)
			p.Match(GOTYPESParserBRACKETS)
		}

	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserPOINTER {
		{
			p.SetState(52)
			p.Match(GOTYPESParserPOINTER)
		}

	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(55)
			p.Match(GOTYPESParserIDENTIFIER)
		}

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
	p.EnterRule(localctx, 8, GOTYPESParserRULE_structdef)
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
		p.SetState(58)
		p.Match(GOTYPESParserT__4)
	}
	{
		p.SetState(59)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(60)
		p.Match(GOTYPESParserT__5)
	}
	{
		p.SetState(61)
		p.Match(GOTYPESParserT__6)
	}
	{
		p.SetState(62)
		p.Match(GOTYPESParserNEWLINE)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GOTYPESParserPOINTER)|(1<<GOTYPESParserBRACKETS)|(1<<GOTYPESParserIDENTIFIER))) != 0) {
		{
			p.SetState(63)
			p.Structitem()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == GOTYPESParserNEWLINE {
			{
				p.SetState(64)
				p.Match(GOTYPESParserNEWLINE)
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
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
	p.EnterRule(localctx, 10, GOTYPESParserRULE_structitem)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(76)
			p.Match(GOTYPESParserIDENTIFIER)
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserBRACKETS {
		{
			p.SetState(79)
			p.Match(GOTYPESParserBRACKETS)
		}

	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserPOINTER {
		{
			p.SetState(82)
			p.Match(GOTYPESParserPOINTER)
		}

	}
	{
		p.SetState(85)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GOTYPESParserT__8 {
		{
			p.SetState(86)
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
	p.EnterRule(localctx, 12, GOTYPESParserRULE_tag)
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
		p.SetState(89)
		p.Match(GOTYPESParserT__8)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GOTYPESParserIDENTIFIER {
		{
			p.SetState(90)
			p.Tagitem()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
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
	p.EnterRule(localctx, 14, GOTYPESParserRULE_tagitem)

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
		p.SetState(97)
		p.Match(GOTYPESParserIDENTIFIER)
	}
	{
		p.SetState(98)
		p.Match(GOTYPESParserT__9)
	}
	{
		p.SetState(99)
		p.Match(GOTYPESParserCSTRING)
	}

	return localctx
}
