// Code generated from ./PREFIX.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 70, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 7, 10, 56, 10, 10, 12, 10, 14, 10,
	59, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5,
	12, 69, 10, 12, 2, 2, 13, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9,
	17, 10, 19, 11, 21, 12, 23, 2, 3, 2, 5, 5, 2, 11, 12, 14, 15, 34, 34, 4,
	2, 36, 36, 94, 94, 8, 2, 36, 36, 100, 100, 104, 104, 112, 112, 116, 116,
	118, 118, 2, 72, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2,
	2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2,
	2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 25, 3, 2,
	2, 2, 5, 27, 3, 2, 2, 2, 7, 32, 3, 2, 2, 2, 9, 34, 3, 2, 2, 2, 11, 36,
	3, 2, 2, 2, 13, 40, 3, 2, 2, 2, 15, 43, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2,
	19, 52, 3, 2, 2, 2, 21, 62, 3, 2, 2, 2, 23, 64, 3, 2, 2, 2, 25, 26, 7,
	42, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 99, 2, 2, 28, 29, 7, 118, 2, 2,
	29, 30, 7, 113, 2, 2, 30, 31, 7, 111, 2, 2, 31, 6, 3, 2, 2, 2, 32, 33,
	7, 46, 2, 2, 33, 8, 3, 2, 2, 2, 34, 35, 7, 43, 2, 2, 35, 10, 3, 2, 2, 2,
	36, 37, 7, 99, 2, 2, 37, 38, 7, 112, 2, 2, 38, 39, 7, 102, 2, 2, 39, 12,
	3, 2, 2, 2, 40, 41, 7, 113, 2, 2, 41, 42, 7, 116, 2, 2, 42, 14, 3, 2, 2,
	2, 43, 44, 7, 118, 2, 2, 44, 45, 7, 116, 2, 2, 45, 46, 7, 119, 2, 2, 46,
	47, 7, 103, 2, 2, 47, 16, 3, 2, 2, 2, 48, 49, 9, 2, 2, 2, 49, 50, 3, 2,
	2, 2, 50, 51, 8, 9, 2, 2, 51, 18, 3, 2, 2, 2, 52, 57, 7, 36, 2, 2, 53,
	56, 5, 23, 12, 2, 54, 56, 10, 3, 2, 2, 55, 53, 3, 2, 2, 2, 55, 54, 3, 2,
	2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60,
	3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60, 61, 7, 36, 2, 2, 61, 20, 3, 2, 2, 2,
	62, 63, 7, 41, 2, 2, 63, 22, 3, 2, 2, 2, 64, 68, 7, 94, 2, 2, 65, 69, 9,
	4, 2, 2, 66, 69, 5, 21, 11, 2, 67, 69, 7, 94, 2, 2, 68, 65, 3, 2, 2, 2,
	68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69, 24, 3, 2, 2, 2, 6, 2, 55, 57,
	68, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "'atom'", "','", "')'", "'and'", "'or'", "'true'", "", "", "'''",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "WS", "CSTRING", "APOSTROPHE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "WS", "CSTRING",
	"APOSTROPHE", "EscapeSequence",
}

type PREFIXLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPREFIXLexer(input antlr.CharStream) *PREFIXLexer {

	l := new(PREFIXLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "PREFIX.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PREFIXLexer tokens.
const (
	PREFIXLexerT__0       = 1
	PREFIXLexerT__1       = 2
	PREFIXLexerT__2       = 3
	PREFIXLexerT__3       = 4
	PREFIXLexerT__4       = 5
	PREFIXLexerT__5       = 6
	PREFIXLexerT__6       = 7
	PREFIXLexerWS         = 8
	PREFIXLexerCSTRING    = 9
	PREFIXLexerAPOSTROPHE = 10
)
