// Code generated from C:/Users/navar/GolandProjects/MiniGolang/parser/expr_lexer.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type expr_lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Expr_lexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expr_lexerLexerInit() {
	staticData := &Expr_lexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'package'", "'var'", "'type'", "'func'", "'struct'", "'append'",
		"'len'", "'cap'", "'print'", "'return'", "'break'", "'continue'", "'if'",
		"'else'", "'for'", "'switch'", "'case'", "'default'", "'='", "':='",
		"';'", "':'", "'.'", "'('", "')'", "']'", "'['", "'{'", "'}'", "'+'",
		"'-'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'&^'", "'|'", "'^'",
		"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'",
		"'++'", "'--'", "'+='", "'&='", "'-='", "'|='", "'*='", "'^='", "'<<='",
		"'>>='", "'&^='", "'%='", "'/='",
	}
	staticData.SymbolicNames = []string{
		"", "PACKAGE", "VAR", "TYPE", "FUNCTION", "STRUCT", "APPEND", "LENGHT",
		"CAP", "PRINT", "RETURN", "BREAK", "CONTINUE", "IF", "ELSE", "FOR",
		"SWITH", "CASE", "DEFAULT", "ASSIGN", "SHORT_DEC", "SEMICOLON", "COLON",
		"DOT", "LPAREN", "DPAREN", "RBRACK", "LBRACK", "LBRACE", "RBRACE", "ADD",
		"SUB", "MULT", "DIV", "MOD", "LSHIFT", "RSHIFT", "AND", "ANDNOT", "PIPE",
		"CARET", "EQUALS", "NOT_EQ", "LESS_THAN", "LESS_THAN_OR_EQUALS", "GREATER_THAN",
		"GREATER_THAN_OR_EQUALS", "LOG_AND", "LOG_OR", "LOG_NOT", "INCREMENT",
		"DECREMENT", "PLUS_ASSIGN", "AND_ASSIGN", "MINUS_ASSIGN", "OR_ASSIGN",
		"TIMES_ASSIGN", "XOR_ASSIGN", "LEFT_SHIFT_ASSIGN", "RIGHT_SHIFT_ASSIGN",
		"AND_NOT_ASSIGN", "MOD_ASSIGN", "DIVIDE_ASSIGN",
	}
	staticData.RuleNames = []string{
		"PACKAGE", "VAR", "TYPE", "FUNCTION", "STRUCT", "APPEND", "LENGHT",
		"CAP", "PRINT", "RETURN", "BREAK", "CONTINUE", "IF", "ELSE", "FOR",
		"SWITH", "CASE", "DEFAULT", "ASSIGN", "SHORT_DEC", "SEMICOLON", "COLON",
		"DOT", "LPAREN", "DPAREN", "RBRACK", "LBRACK", "LBRACE", "RBRACE", "ADD",
		"SUB", "MULT", "DIV", "MOD", "LSHIFT", "RSHIFT", "AND", "ANDNOT", "PIPE",
		"CARET", "EQUALS", "NOT_EQ", "LESS_THAN", "LESS_THAN_OR_EQUALS", "GREATER_THAN",
		"GREATER_THAN_OR_EQUALS", "LOG_AND", "LOG_OR", "LOG_NOT", "INCREMENT",
		"DECREMENT", "PLUS_ASSIGN", "AND_ASSIGN", "MINUS_ASSIGN", "OR_ASSIGN",
		"TIMES_ASSIGN", "XOR_ASSIGN", "LEFT_SHIFT_ASSIGN", "RIGHT_SHIFT_ASSIGN",
		"AND_NOT_ASSIGN", "MOD_ASSIGN", "DIVIDE_ASSIGN",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 62, 343, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1,
		32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1,
		41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52,
		1, 52, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1, 55, 1, 55, 1,
		56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 58,
		1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 0,
		0, 62, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46,
		93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109,
		55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 1, 0,
		0, 342, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61,
		1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0,
		69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0,
		0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0,
		0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0,
		0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1,
		0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0,
		107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0,
		0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121,
		1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 1, 125, 1, 0, 0, 0, 3, 133, 1, 0, 0, 0,
		5, 137, 1, 0, 0, 0, 7, 142, 1, 0, 0, 0, 9, 147, 1, 0, 0, 0, 11, 154, 1,
		0, 0, 0, 13, 161, 1, 0, 0, 0, 15, 165, 1, 0, 0, 0, 17, 169, 1, 0, 0, 0,
		19, 175, 1, 0, 0, 0, 21, 182, 1, 0, 0, 0, 23, 188, 1, 0, 0, 0, 25, 197,
		1, 0, 0, 0, 27, 200, 1, 0, 0, 0, 29, 205, 1, 0, 0, 0, 31, 209, 1, 0, 0,
		0, 33, 216, 1, 0, 0, 0, 35, 221, 1, 0, 0, 0, 37, 229, 1, 0, 0, 0, 39, 231,
		1, 0, 0, 0, 41, 234, 1, 0, 0, 0, 43, 236, 1, 0, 0, 0, 45, 238, 1, 0, 0,
		0, 47, 240, 1, 0, 0, 0, 49, 242, 1, 0, 0, 0, 51, 244, 1, 0, 0, 0, 53, 246,
		1, 0, 0, 0, 55, 248, 1, 0, 0, 0, 57, 250, 1, 0, 0, 0, 59, 252, 1, 0, 0,
		0, 61, 254, 1, 0, 0, 0, 63, 256, 1, 0, 0, 0, 65, 258, 1, 0, 0, 0, 67, 260,
		1, 0, 0, 0, 69, 262, 1, 0, 0, 0, 71, 265, 1, 0, 0, 0, 73, 268, 1, 0, 0,
		0, 75, 270, 1, 0, 0, 0, 77, 273, 1, 0, 0, 0, 79, 275, 1, 0, 0, 0, 81, 277,
		1, 0, 0, 0, 83, 280, 1, 0, 0, 0, 85, 283, 1, 0, 0, 0, 87, 285, 1, 0, 0,
		0, 89, 288, 1, 0, 0, 0, 91, 290, 1, 0, 0, 0, 93, 293, 1, 0, 0, 0, 95, 296,
		1, 0, 0, 0, 97, 299, 1, 0, 0, 0, 99, 301, 1, 0, 0, 0, 101, 304, 1, 0, 0,
		0, 103, 307, 1, 0, 0, 0, 105, 310, 1, 0, 0, 0, 107, 313, 1, 0, 0, 0, 109,
		316, 1, 0, 0, 0, 111, 319, 1, 0, 0, 0, 113, 322, 1, 0, 0, 0, 115, 325,
		1, 0, 0, 0, 117, 329, 1, 0, 0, 0, 119, 333, 1, 0, 0, 0, 121, 337, 1, 0,
		0, 0, 123, 340, 1, 0, 0, 0, 125, 126, 5, 112, 0, 0, 126, 127, 5, 97, 0,
		0, 127, 128, 5, 99, 0, 0, 128, 129, 5, 107, 0, 0, 129, 130, 5, 97, 0, 0,
		130, 131, 5, 103, 0, 0, 131, 132, 5, 101, 0, 0, 132, 2, 1, 0, 0, 0, 133,
		134, 5, 118, 0, 0, 134, 135, 5, 97, 0, 0, 135, 136, 5, 114, 0, 0, 136,
		4, 1, 0, 0, 0, 137, 138, 5, 116, 0, 0, 138, 139, 5, 121, 0, 0, 139, 140,
		5, 112, 0, 0, 140, 141, 5, 101, 0, 0, 141, 6, 1, 0, 0, 0, 142, 143, 5,
		102, 0, 0, 143, 144, 5, 117, 0, 0, 144, 145, 5, 110, 0, 0, 145, 146, 5,
		99, 0, 0, 146, 8, 1, 0, 0, 0, 147, 148, 5, 115, 0, 0, 148, 149, 5, 116,
		0, 0, 149, 150, 5, 114, 0, 0, 150, 151, 5, 117, 0, 0, 151, 152, 5, 99,
		0, 0, 152, 153, 5, 116, 0, 0, 153, 10, 1, 0, 0, 0, 154, 155, 5, 97, 0,
		0, 155, 156, 5, 112, 0, 0, 156, 157, 5, 112, 0, 0, 157, 158, 5, 101, 0,
		0, 158, 159, 5, 110, 0, 0, 159, 160, 5, 100, 0, 0, 160, 12, 1, 0, 0, 0,
		161, 162, 5, 108, 0, 0, 162, 163, 5, 101, 0, 0, 163, 164, 5, 110, 0, 0,
		164, 14, 1, 0, 0, 0, 165, 166, 5, 99, 0, 0, 166, 167, 5, 97, 0, 0, 167,
		168, 5, 112, 0, 0, 168, 16, 1, 0, 0, 0, 169, 170, 5, 112, 0, 0, 170, 171,
		5, 114, 0, 0, 171, 172, 5, 105, 0, 0, 172, 173, 5, 110, 0, 0, 173, 174,
		5, 116, 0, 0, 174, 18, 1, 0, 0, 0, 175, 176, 5, 114, 0, 0, 176, 177, 5,
		101, 0, 0, 177, 178, 5, 116, 0, 0, 178, 179, 5, 117, 0, 0, 179, 180, 5,
		114, 0, 0, 180, 181, 5, 110, 0, 0, 181, 20, 1, 0, 0, 0, 182, 183, 5, 98,
		0, 0, 183, 184, 5, 114, 0, 0, 184, 185, 5, 101, 0, 0, 185, 186, 5, 97,
		0, 0, 186, 187, 5, 107, 0, 0, 187, 22, 1, 0, 0, 0, 188, 189, 5, 99, 0,
		0, 189, 190, 5, 111, 0, 0, 190, 191, 5, 110, 0, 0, 191, 192, 5, 116, 0,
		0, 192, 193, 5, 105, 0, 0, 193, 194, 5, 110, 0, 0, 194, 195, 5, 117, 0,
		0, 195, 196, 5, 101, 0, 0, 196, 24, 1, 0, 0, 0, 197, 198, 5, 105, 0, 0,
		198, 199, 5, 102, 0, 0, 199, 26, 1, 0, 0, 0, 200, 201, 5, 101, 0, 0, 201,
		202, 5, 108, 0, 0, 202, 203, 5, 115, 0, 0, 203, 204, 5, 101, 0, 0, 204,
		28, 1, 0, 0, 0, 205, 206, 5, 102, 0, 0, 206, 207, 5, 111, 0, 0, 207, 208,
		5, 114, 0, 0, 208, 30, 1, 0, 0, 0, 209, 210, 5, 115, 0, 0, 210, 211, 5,
		119, 0, 0, 211, 212, 5, 105, 0, 0, 212, 213, 5, 116, 0, 0, 213, 214, 5,
		99, 0, 0, 214, 215, 5, 104, 0, 0, 215, 32, 1, 0, 0, 0, 216, 217, 5, 99,
		0, 0, 217, 218, 5, 97, 0, 0, 218, 219, 5, 115, 0, 0, 219, 220, 5, 101,
		0, 0, 220, 34, 1, 0, 0, 0, 221, 222, 5, 100, 0, 0, 222, 223, 5, 101, 0,
		0, 223, 224, 5, 102, 0, 0, 224, 225, 5, 97, 0, 0, 225, 226, 5, 117, 0,
		0, 226, 227, 5, 108, 0, 0, 227, 228, 5, 116, 0, 0, 228, 36, 1, 0, 0, 0,
		229, 230, 5, 61, 0, 0, 230, 38, 1, 0, 0, 0, 231, 232, 5, 58, 0, 0, 232,
		233, 5, 61, 0, 0, 233, 40, 1, 0, 0, 0, 234, 235, 5, 59, 0, 0, 235, 42,
		1, 0, 0, 0, 236, 237, 5, 58, 0, 0, 237, 44, 1, 0, 0, 0, 238, 239, 5, 46,
		0, 0, 239, 46, 1, 0, 0, 0, 240, 241, 5, 40, 0, 0, 241, 48, 1, 0, 0, 0,
		242, 243, 5, 41, 0, 0, 243, 50, 1, 0, 0, 0, 244, 245, 5, 93, 0, 0, 245,
		52, 1, 0, 0, 0, 246, 247, 5, 91, 0, 0, 247, 54, 1, 0, 0, 0, 248, 249, 5,
		123, 0, 0, 249, 56, 1, 0, 0, 0, 250, 251, 5, 125, 0, 0, 251, 58, 1, 0,
		0, 0, 252, 253, 5, 43, 0, 0, 253, 60, 1, 0, 0, 0, 254, 255, 5, 45, 0, 0,
		255, 62, 1, 0, 0, 0, 256, 257, 5, 42, 0, 0, 257, 64, 1, 0, 0, 0, 258, 259,
		5, 47, 0, 0, 259, 66, 1, 0, 0, 0, 260, 261, 5, 37, 0, 0, 261, 68, 1, 0,
		0, 0, 262, 263, 5, 60, 0, 0, 263, 264, 5, 60, 0, 0, 264, 70, 1, 0, 0, 0,
		265, 266, 5, 62, 0, 0, 266, 267, 5, 62, 0, 0, 267, 72, 1, 0, 0, 0, 268,
		269, 5, 38, 0, 0, 269, 74, 1, 0, 0, 0, 270, 271, 5, 38, 0, 0, 271, 272,
		5, 94, 0, 0, 272, 76, 1, 0, 0, 0, 273, 274, 5, 124, 0, 0, 274, 78, 1, 0,
		0, 0, 275, 276, 5, 94, 0, 0, 276, 80, 1, 0, 0, 0, 277, 278, 5, 61, 0, 0,
		278, 279, 5, 61, 0, 0, 279, 82, 1, 0, 0, 0, 280, 281, 5, 33, 0, 0, 281,
		282, 5, 61, 0, 0, 282, 84, 1, 0, 0, 0, 283, 284, 5, 60, 0, 0, 284, 86,
		1, 0, 0, 0, 285, 286, 5, 60, 0, 0, 286, 287, 5, 61, 0, 0, 287, 88, 1, 0,
		0, 0, 288, 289, 5, 62, 0, 0, 289, 90, 1, 0, 0, 0, 290, 291, 5, 62, 0, 0,
		291, 292, 5, 61, 0, 0, 292, 92, 1, 0, 0, 0, 293, 294, 5, 38, 0, 0, 294,
		295, 5, 38, 0, 0, 295, 94, 1, 0, 0, 0, 296, 297, 5, 124, 0, 0, 297, 298,
		5, 124, 0, 0, 298, 96, 1, 0, 0, 0, 299, 300, 5, 33, 0, 0, 300, 98, 1, 0,
		0, 0, 301, 302, 5, 43, 0, 0, 302, 303, 5, 43, 0, 0, 303, 100, 1, 0, 0,
		0, 304, 305, 5, 45, 0, 0, 305, 306, 5, 45, 0, 0, 306, 102, 1, 0, 0, 0,
		307, 308, 5, 43, 0, 0, 308, 309, 5, 61, 0, 0, 309, 104, 1, 0, 0, 0, 310,
		311, 5, 38, 0, 0, 311, 312, 5, 61, 0, 0, 312, 106, 1, 0, 0, 0, 313, 314,
		5, 45, 0, 0, 314, 315, 5, 61, 0, 0, 315, 108, 1, 0, 0, 0, 316, 317, 5,
		124, 0, 0, 317, 318, 5, 61, 0, 0, 318, 110, 1, 0, 0, 0, 319, 320, 5, 42,
		0, 0, 320, 321, 5, 61, 0, 0, 321, 112, 1, 0, 0, 0, 322, 323, 5, 94, 0,
		0, 323, 324, 5, 61, 0, 0, 324, 114, 1, 0, 0, 0, 325, 326, 5, 60, 0, 0,
		326, 327, 5, 60, 0, 0, 327, 328, 5, 61, 0, 0, 328, 116, 1, 0, 0, 0, 329,
		330, 5, 62, 0, 0, 330, 331, 5, 62, 0, 0, 331, 332, 5, 61, 0, 0, 332, 118,
		1, 0, 0, 0, 333, 334, 5, 38, 0, 0, 334, 335, 5, 94, 0, 0, 335, 336, 5,
		61, 0, 0, 336, 120, 1, 0, 0, 0, 337, 338, 5, 37, 0, 0, 338, 339, 5, 61,
		0, 0, 339, 122, 1, 0, 0, 0, 340, 341, 5, 47, 0, 0, 341, 342, 5, 61, 0,
		0, 342, 124, 1, 0, 0, 0, 1, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// expr_lexerInit initializes any static state used to implement expr_lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newexpr_lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Expr_lexerInit() {
	staticData := &Expr_lexerLexerStaticData
	staticData.once.Do(expr_lexerLexerInit)
}

// Newexpr_lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newexpr_lexer(input antlr.CharStream) *expr_lexer {
	Expr_lexerInit()
	l := new(expr_lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Expr_lexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "expr_lexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// expr_lexer tokens.
const (
	expr_lexerPACKAGE                = 1
	expr_lexerVAR                    = 2
	expr_lexerTYPE                   = 3
	expr_lexerFUNCTION               = 4
	expr_lexerSTRUCT                 = 5
	expr_lexerAPPEND                 = 6
	expr_lexerLENGHT                 = 7
	expr_lexerCAP                    = 8
	expr_lexerPRINT                  = 9
	expr_lexerRETURN                 = 10
	expr_lexerBREAK                  = 11
	expr_lexerCONTINUE               = 12
	expr_lexerIF                     = 13
	expr_lexerELSE                   = 14
	expr_lexerFOR                    = 15
	expr_lexerSWITH                  = 16
	expr_lexerCASE                   = 17
	expr_lexerDEFAULT                = 18
	expr_lexerASSIGN                 = 19
	expr_lexerSHORT_DEC              = 20
	expr_lexerSEMICOLON              = 21
	expr_lexerCOLON                  = 22
	expr_lexerDOT                    = 23
	expr_lexerLPAREN                 = 24
	expr_lexerDPAREN                 = 25
	expr_lexerRBRACK                 = 26
	expr_lexerLBRACK                 = 27
	expr_lexerLBRACE                 = 28
	expr_lexerRBRACE                 = 29
	expr_lexerADD                    = 30
	expr_lexerSUB                    = 31
	expr_lexerMULT                   = 32
	expr_lexerDIV                    = 33
	expr_lexerMOD                    = 34
	expr_lexerLSHIFT                 = 35
	expr_lexerRSHIFT                 = 36
	expr_lexerAND                    = 37
	expr_lexerANDNOT                 = 38
	expr_lexerPIPE                   = 39
	expr_lexerCARET                  = 40
	expr_lexerEQUALS                 = 41
	expr_lexerNOT_EQ                 = 42
	expr_lexerLESS_THAN              = 43
	expr_lexerLESS_THAN_OR_EQUALS    = 44
	expr_lexerGREATER_THAN           = 45
	expr_lexerGREATER_THAN_OR_EQUALS = 46
	expr_lexerLOG_AND                = 47
	expr_lexerLOG_OR                 = 48
	expr_lexerLOG_NOT                = 49
	expr_lexerINCREMENT              = 50
	expr_lexerDECREMENT              = 51
	expr_lexerPLUS_ASSIGN            = 52
	expr_lexerAND_ASSIGN             = 53
	expr_lexerMINUS_ASSIGN           = 54
	expr_lexerOR_ASSIGN              = 55
	expr_lexerTIMES_ASSIGN           = 56
	expr_lexerXOR_ASSIGN             = 57
	expr_lexerLEFT_SHIFT_ASSIGN      = 58
	expr_lexerRIGHT_SHIFT_ASSIGN     = 59
	expr_lexerAND_NOT_ASSIGN         = 60
	expr_lexerMOD_ASSIGN             = 61
	expr_lexerDIVIDE_ASSIGN          = 62
)
