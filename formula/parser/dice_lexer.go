// Code generated from /Users/xguzman/Projects/dice/formula/Dice.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 58, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3, 
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 
	3, 8, 3, 9, 6, 9, 41, 10, 9, 13, 9, 14, 9, 42, 3, 10, 3, 10, 6, 10, 47, 
	10, 10, 13, 10, 14, 10, 48, 3, 11, 3, 11, 6, 11, 53, 10, 11, 13, 11, 14, 
	11, 54, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 
	9, 17, 10, 19, 11, 21, 12, 3, 2, 9, 4, 2, 70, 70, 102, 102, 4, 2, 45, 45, 
	47, 47, 4, 2, 11, 12, 15, 15, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 5, 2, 
	50, 59, 67, 92, 99, 124, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 2, 60, 2, 
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 
	2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 31, 3, 2, 2, 2, 13, 33, 3, 
	2, 2, 2, 15, 35, 3, 2, 2, 2, 17, 40, 3, 2, 2, 2, 19, 44, 3, 2, 2, 2, 21, 
	50, 3, 2, 2, 2, 23, 24, 9, 2, 2, 2, 24, 4, 3, 2, 2, 2, 25, 26, 9, 3, 2, 
	2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 42, 2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 
	43, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32, 7, 46, 2, 2, 32, 12, 3, 2, 2, 2, 
	33, 34, 7, 34, 2, 2, 34, 14, 3, 2, 2, 2, 35, 36, 9, 4, 2, 2, 36, 37, 3, 
	2, 2, 2, 37, 38, 8, 8, 2, 2, 38, 16, 3, 2, 2, 2, 39, 41, 9, 5, 2, 2, 40, 
	39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 
	2, 43, 18, 3, 2, 2, 2, 44, 46, 9, 6, 2, 2, 45, 47, 9, 7, 2, 2, 46, 45, 
	3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 
	49, 20, 3, 2, 2, 2, 50, 52, 7, 36, 2, 2, 51, 53, 10, 8, 2, 2, 52, 51, 3, 
	2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 
	56, 3, 2, 2, 2, 56, 57, 7, 36, 2, 2, 57, 22, 3, 2, 2, 2, 6, 2, 42, 48, 
	54, 3, 8, 2, 2,
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
	"", "", "", "'('", "')'", "','", "' '",
}

var lexerSymbolicNames = []string{
	"", "D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer", 
	"Id", "StringLiteral",
}

var lexerRuleNames = []string{
	"D", "SIGN", "LPAREN", "RPAREN", "COMMA", "SPACE", "WS", "Integer", "Id", 
	"StringLiteral",
}

type DiceLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewDiceLexer(input antlr.CharStream) *DiceLexer {

	l := new(DiceLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Dice.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DiceLexer tokens.
const (
	DiceLexerD = 1
	DiceLexerSIGN = 2
	DiceLexerLPAREN = 3
	DiceLexerRPAREN = 4
	DiceLexerCOMMA = 5
	DiceLexerSPACE = 6
	DiceLexerWS = 7
	DiceLexerInteger = 8
	DiceLexerId = 9
	DiceLexerStringLiteral = 10
)

