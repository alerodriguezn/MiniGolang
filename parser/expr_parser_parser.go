// Code generated from C:/Users/navar/GolandProjects/MiniGolang/parser/expr_parser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // expr_parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type expr_parser struct {
	*antlr.BaseParser
}

var Expr_parserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expr_parserParserInit() {
	staticData := &Expr_parserParserStaticData
	staticData.LiteralNames = []string{
		"", "'package'", "'var'", "'type'", "'func'", "'struct'", "'append'",
		"'len'", "'cap'", "'print'", "'println'", "'return'", "'break'", "'continue'",
		"'if'", "'else'", "'for'", "'switch'", "'case'", "'default'", "'='",
		"':='", "';'", "':'", "'.'", "'('", "')'", "']'", "'['", "'{'", "'}'",
		"','", "'+'", "'-'", "'*'", "'/'", "'%'", "'<<'", "'>>'", "'&'", "'&^'",
		"'|'", "'^'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'",
		"'||'", "'!'", "'++'", "'--'", "'+='", "'&='", "'-='", "'|='", "'*='",
		"'^='", "'<<='", "'>>='", "'&^='", "'%='", "'/='",
	}
	staticData.SymbolicNames = []string{
		"", "PACKAGE", "VAR", "TYPE", "FUNCTION", "STRUCT", "APPEND", "LENGHT",
		"CAP", "PRINT", "PRINTLN", "RETURN", "BREAK", "CONTINUE", "IF", "ELSE",
		"FOR", "SWITCH", "CASE", "DEFAULT", "ASSIGN", "SHORT_DEC", "SEMICOLON",
		"COLON", "DOT", "LPAREN", "RPAREN", "RBRACK", "LBRACK", "LBRACE", "RBRACE",
		"COMMA", "ADD", "SUB", "MULT", "DIV", "MOD", "LSHIFT", "RSHIFT", "AND",
		"ANDNOT", "PIPE", "CARET", "EQUALS", "NOT_EQ", "LESS_THAN", "LESS_THAN_OR_EQUALS",
		"GREATER_THAN", "GREATER_THAN_OR_EQUALS", "LOG_AND", "LOG_OR", "LOG_NOT",
		"INCREMENT", "DECREMENT", "PLUS_ASSIGN", "AND_ASSIGN", "MINUS_ASSIGN",
		"OR_ASSIGN", "TIMES_ASSIGN", "XOR_ASSIGN", "LEFT_SHIFT_ASSIGN", "RIGHT_SHIFT_ASSIGN",
		"AND_NOT_ASSIGN", "MOD_ASSIGN", "DIVIDE_ASSIGN", "INT_LIT", "IDENTIFIER",
		"FLOAT_LIT", "RUNE_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"STRING_LIT", "WS",
	}
	staticData.RuleNames = []string{
		"root", "topDeclarationList", "variableDecl", "innerVarDecls", "singleVarDecl",
		"singleVarDeclNoExps", "typeDecl", "innerTypeDecls", "singleTypeDecl",
		"funcDecl", "funcFrontDecl", "funcArgDecls", "declType", "sliceDeclType",
		"arrayDeclType", "structDeclType", "structMemDecls", "identifierList",
		"expression", "expressionList", "primaryExpression", "operand", "literal",
		"index", "arguments", "selector", "appendExpression", "lengthExpression",
		"capExpression", "statementList", "block", "statement", "simpleStatement",
		"assignmentStatement", "ifStatement", "loop", "switch", "expressionCaseClauseList",
		"expressionCaseClause", "expressionSwitchCase",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 72, 521, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 1, 1, 1, 1, 1, 5, 1, 89, 8, 1, 10, 1, 12, 1, 92, 9, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3,
		2, 108, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 115, 8, 3, 10, 3, 12,
		3, 118, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 130, 8, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 149, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 156, 8, 7, 10, 7, 12, 7, 159, 9, 7, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10,
		193, 8, 10, 1, 11, 1, 11, 1, 11, 5, 11, 198, 8, 11, 10, 11, 12, 11, 201,
		9, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 3, 12, 214, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 3, 15, 228, 8, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 237, 8, 16, 10, 16, 12, 16,
		240, 9, 16, 1, 17, 1, 17, 1, 17, 5, 17, 245, 8, 17, 10, 17, 12, 17, 248,
		9, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 254, 8, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 259, 8, 18, 10, 18, 12, 18, 262, 9, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 267, 8, 19, 10, 19, 12, 19, 270, 9, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 3, 20, 277, 8, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		5, 20, 285, 8, 20, 10, 20, 12, 20, 288, 9, 20, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 296, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		3, 22, 303, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 311,
		8, 24, 10, 24, 12, 24, 314, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 5, 29, 339, 8, 29, 10,
		29, 12, 29, 342, 9, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		3, 31, 351, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 358, 8, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 364, 8, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3,
		31, 390, 8, 31, 1, 32, 1, 32, 3, 32, 394, 8, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 3, 32, 401, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 3, 33, 411, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34,
		1, 34, 3, 34, 451, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 3, 35, 474, 8, 35, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 3, 36, 502, 8, 36, 1, 37, 1, 37, 1, 37, 5, 37, 507, 8, 37, 10,
		37, 12, 37, 510, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39,
		3, 39, 519, 8, 39, 1, 39, 0, 2, 36, 40, 40, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 0, 4, 3, 0, 32,
		33, 42, 42, 51, 51, 1, 0, 32, 50, 1, 0, 52, 53, 1, 0, 54, 64, 552, 0, 80,
		1, 0, 0, 0, 2, 90, 1, 0, 0, 0, 4, 107, 1, 0, 0, 0, 6, 109, 1, 0, 0, 0,
		8, 129, 1, 0, 0, 0, 10, 131, 1, 0, 0, 0, 12, 148, 1, 0, 0, 0, 14, 150,
		1, 0, 0, 0, 16, 160, 1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 192, 1, 0, 0,
		0, 22, 199, 1, 0, 0, 0, 24, 213, 1, 0, 0, 0, 26, 215, 1, 0, 0, 0, 28, 219,
		1, 0, 0, 0, 30, 224, 1, 0, 0, 0, 32, 231, 1, 0, 0, 0, 34, 241, 1, 0, 0,
		0, 36, 253, 1, 0, 0, 0, 38, 263, 1, 0, 0, 0, 40, 276, 1, 0, 0, 0, 42, 295,
		1, 0, 0, 0, 44, 302, 1, 0, 0, 0, 46, 304, 1, 0, 0, 0, 48, 308, 1, 0, 0,
		0, 50, 317, 1, 0, 0, 0, 52, 320, 1, 0, 0, 0, 54, 327, 1, 0, 0, 0, 56, 332,
		1, 0, 0, 0, 58, 340, 1, 0, 0, 0, 60, 343, 1, 0, 0, 0, 62, 389, 1, 0, 0,
		0, 64, 400, 1, 0, 0, 0, 66, 410, 1, 0, 0, 0, 68, 450, 1, 0, 0, 0, 70, 473,
		1, 0, 0, 0, 72, 501, 1, 0, 0, 0, 74, 508, 1, 0, 0, 0, 76, 511, 1, 0, 0,
		0, 78, 518, 1, 0, 0, 0, 80, 81, 5, 1, 0, 0, 81, 82, 5, 66, 0, 0, 82, 83,
		5, 22, 0, 0, 83, 84, 3, 2, 1, 0, 84, 1, 1, 0, 0, 0, 85, 89, 3, 4, 2, 0,
		86, 89, 3, 12, 6, 0, 87, 89, 3, 18, 9, 0, 88, 85, 1, 0, 0, 0, 88, 86, 1,
		0, 0, 0, 88, 87, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90,
		91, 1, 0, 0, 0, 91, 3, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 93, 94, 5, 2, 0,
		0, 94, 95, 3, 8, 4, 0, 95, 96, 5, 22, 0, 0, 96, 108, 1, 0, 0, 0, 97, 98,
		5, 2, 0, 0, 98, 99, 5, 25, 0, 0, 99, 100, 3, 6, 3, 0, 100, 101, 5, 26,
		0, 0, 101, 102, 5, 22, 0, 0, 102, 108, 1, 0, 0, 0, 103, 104, 5, 2, 0, 0,
		104, 105, 5, 25, 0, 0, 105, 106, 5, 26, 0, 0, 106, 108, 5, 22, 0, 0, 107,
		93, 1, 0, 0, 0, 107, 97, 1, 0, 0, 0, 107, 103, 1, 0, 0, 0, 108, 5, 1, 0,
		0, 0, 109, 110, 3, 8, 4, 0, 110, 116, 5, 22, 0, 0, 111, 112, 3, 8, 4, 0,
		112, 113, 5, 22, 0, 0, 113, 115, 1, 0, 0, 0, 114, 111, 1, 0, 0, 0, 115,
		118, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 116, 117, 1, 0, 0, 0, 117, 7, 1,
		0, 0, 0, 118, 116, 1, 0, 0, 0, 119, 120, 3, 34, 17, 0, 120, 121, 3, 24,
		12, 0, 121, 122, 5, 20, 0, 0, 122, 123, 3, 38, 19, 0, 123, 130, 1, 0, 0,
		0, 124, 125, 3, 34, 17, 0, 125, 126, 5, 20, 0, 0, 126, 127, 3, 38, 19,
		0, 127, 130, 1, 0, 0, 0, 128, 130, 3, 10, 5, 0, 129, 119, 1, 0, 0, 0, 129,
		124, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 9, 1, 0, 0, 0, 131, 132, 3,
		34, 17, 0, 132, 133, 3, 24, 12, 0, 133, 11, 1, 0, 0, 0, 134, 135, 5, 3,
		0, 0, 135, 136, 3, 16, 8, 0, 136, 137, 5, 22, 0, 0, 137, 149, 1, 0, 0,
		0, 138, 139, 5, 3, 0, 0, 139, 140, 5, 25, 0, 0, 140, 141, 3, 14, 7, 0,
		141, 142, 5, 26, 0, 0, 142, 143, 5, 22, 0, 0, 143, 149, 1, 0, 0, 0, 144,
		145, 5, 3, 0, 0, 145, 146, 5, 25, 0, 0, 146, 147, 5, 26, 0, 0, 147, 149,
		5, 22, 0, 0, 148, 134, 1, 0, 0, 0, 148, 138, 1, 0, 0, 0, 148, 144, 1, 0,
		0, 0, 149, 13, 1, 0, 0, 0, 150, 151, 3, 16, 8, 0, 151, 157, 5, 22, 0, 0,
		152, 153, 3, 16, 8, 0, 153, 154, 5, 22, 0, 0, 154, 156, 1, 0, 0, 0, 155,
		152, 1, 0, 0, 0, 156, 159, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158,
		1, 0, 0, 0, 158, 15, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 66,
		0, 0, 161, 162, 3, 24, 12, 0, 162, 17, 1, 0, 0, 0, 163, 164, 3, 20, 10,
		0, 164, 165, 3, 60, 30, 0, 165, 166, 5, 22, 0, 0, 166, 19, 1, 0, 0, 0,
		167, 168, 5, 4, 0, 0, 168, 169, 5, 66, 0, 0, 169, 170, 5, 25, 0, 0, 170,
		171, 3, 22, 11, 0, 171, 172, 5, 26, 0, 0, 172, 173, 3, 24, 12, 0, 173,
		193, 1, 0, 0, 0, 174, 175, 5, 4, 0, 0, 175, 176, 5, 66, 0, 0, 176, 177,
		5, 25, 0, 0, 177, 178, 3, 22, 11, 0, 178, 179, 5, 26, 0, 0, 179, 193, 1,
		0, 0, 0, 180, 181, 5, 4, 0, 0, 181, 182, 5, 66, 0, 0, 182, 183, 5, 25,
		0, 0, 183, 184, 5, 26, 0, 0, 184, 193, 3, 24, 12, 0, 185, 186, 5, 4, 0,
		0, 186, 187, 5, 66, 0, 0, 187, 188, 5, 25, 0, 0, 188, 193, 5, 26, 0, 0,
		189, 190, 5, 4, 0, 0, 190, 191, 5, 66, 0, 0, 191, 193, 3, 24, 12, 0, 192,
		167, 1, 0, 0, 0, 192, 174, 1, 0, 0, 0, 192, 180, 1, 0, 0, 0, 192, 185,
		1, 0, 0, 0, 192, 189, 1, 0, 0, 0, 193, 21, 1, 0, 0, 0, 194, 195, 3, 10,
		5, 0, 195, 196, 5, 31, 0, 0, 196, 198, 1, 0, 0, 0, 197, 194, 1, 0, 0, 0,
		198, 201, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200,
		202, 1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 203, 3, 10, 5, 0, 203, 23,
		1, 0, 0, 0, 204, 205, 5, 25, 0, 0, 205, 206, 3, 24, 12, 0, 206, 207, 5,
		26, 0, 0, 207, 214, 1, 0, 0, 0, 208, 214, 5, 66, 0, 0, 209, 214, 3, 26,
		13, 0, 210, 214, 3, 28, 14, 0, 211, 214, 3, 30, 15, 0, 212, 214, 1, 0,
		0, 0, 213, 204, 1, 0, 0, 0, 213, 208, 1, 0, 0, 0, 213, 209, 1, 0, 0, 0,
		213, 210, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 212, 1, 0, 0, 0, 214,
		25, 1, 0, 0, 0, 215, 216, 5, 28, 0, 0, 216, 217, 5, 27, 0, 0, 217, 218,
		3, 24, 12, 0, 218, 27, 1, 0, 0, 0, 219, 220, 5, 28, 0, 0, 220, 221, 5,
		65, 0, 0, 221, 222, 5, 27, 0, 0, 222, 223, 3, 24, 12, 0, 223, 29, 1, 0,
		0, 0, 224, 225, 5, 5, 0, 0, 225, 227, 5, 29, 0, 0, 226, 228, 3, 32, 16,
		0, 227, 226, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229,
		230, 5, 30, 0, 0, 230, 31, 1, 0, 0, 0, 231, 232, 3, 10, 5, 0, 232, 238,
		5, 22, 0, 0, 233, 234, 3, 10, 5, 0, 234, 235, 5, 22, 0, 0, 235, 237, 1,
		0, 0, 0, 236, 233, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0,
		0, 238, 239, 1, 0, 0, 0, 239, 33, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241,
		246, 5, 66, 0, 0, 242, 243, 5, 31, 0, 0, 243, 245, 5, 66, 0, 0, 244, 242,
		1, 0, 0, 0, 245, 248, 1, 0, 0, 0, 246, 244, 1, 0, 0, 0, 246, 247, 1, 0,
		0, 0, 247, 35, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 249, 250, 6, 18, -1, 0,
		250, 254, 3, 40, 20, 0, 251, 252, 7, 0, 0, 0, 252, 254, 3, 36, 18, 1, 253,
		249, 1, 0, 0, 0, 253, 251, 1, 0, 0, 0, 254, 260, 1, 0, 0, 0, 255, 256,
		10, 2, 0, 0, 256, 257, 7, 1, 0, 0, 257, 259, 3, 36, 18, 3, 258, 255, 1,
		0, 0, 0, 259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0,
		0, 261, 37, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 268, 3, 36, 18, 0, 264,
		265, 5, 31, 0, 0, 265, 267, 3, 36, 18, 0, 266, 264, 1, 0, 0, 0, 267, 270,
		1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269, 39, 1, 0,
		0, 0, 270, 268, 1, 0, 0, 0, 271, 272, 6, 20, -1, 0, 272, 277, 3, 42, 21,
		0, 273, 277, 3, 52, 26, 0, 274, 277, 3, 54, 27, 0, 275, 277, 3, 56, 28,
		0, 276, 271, 1, 0, 0, 0, 276, 273, 1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276,
		275, 1, 0, 0, 0, 277, 286, 1, 0, 0, 0, 278, 279, 10, 6, 0, 0, 279, 285,
		3, 50, 25, 0, 280, 281, 10, 5, 0, 0, 281, 285, 3, 46, 23, 0, 282, 283,
		10, 4, 0, 0, 283, 285, 3, 48, 24, 0, 284, 278, 1, 0, 0, 0, 284, 280, 1,
		0, 0, 0, 284, 282, 1, 0, 0, 0, 285, 288, 1, 0, 0, 0, 286, 284, 1, 0, 0,
		0, 286, 287, 1, 0, 0, 0, 287, 41, 1, 0, 0, 0, 288, 286, 1, 0, 0, 0, 289,
		296, 3, 44, 22, 0, 290, 296, 5, 66, 0, 0, 291, 292, 5, 25, 0, 0, 292, 293,
		3, 36, 18, 0, 293, 294, 5, 26, 0, 0, 294, 296, 1, 0, 0, 0, 295, 289, 1,
		0, 0, 0, 295, 290, 1, 0, 0, 0, 295, 291, 1, 0, 0, 0, 296, 43, 1, 0, 0,
		0, 297, 303, 5, 65, 0, 0, 298, 303, 5, 67, 0, 0, 299, 303, 5, 69, 0, 0,
		300, 303, 5, 70, 0, 0, 301, 303, 5, 68, 0, 0, 302, 297, 1, 0, 0, 0, 302,
		298, 1, 0, 0, 0, 302, 299, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 302, 301,
		1, 0, 0, 0, 303, 45, 1, 0, 0, 0, 304, 305, 5, 28, 0, 0, 305, 306, 3, 36,
		18, 0, 306, 307, 5, 27, 0, 0, 307, 47, 1, 0, 0, 0, 308, 312, 5, 25, 0,
		0, 309, 311, 3, 38, 19, 0, 310, 309, 1, 0, 0, 0, 311, 314, 1, 0, 0, 0,
		312, 310, 1, 0, 0, 0, 312, 313, 1, 0, 0, 0, 313, 315, 1, 0, 0, 0, 314,
		312, 1, 0, 0, 0, 315, 316, 5, 26, 0, 0, 316, 49, 1, 0, 0, 0, 317, 318,
		5, 24, 0, 0, 318, 319, 5, 66, 0, 0, 319, 51, 1, 0, 0, 0, 320, 321, 5, 6,
		0, 0, 321, 322, 5, 25, 0, 0, 322, 323, 3, 36, 18, 0, 323, 324, 5, 31, 0,
		0, 324, 325, 3, 36, 18, 0, 325, 326, 5, 26, 0, 0, 326, 53, 1, 0, 0, 0,
		327, 328, 5, 7, 0, 0, 328, 329, 5, 25, 0, 0, 329, 330, 3, 36, 18, 0, 330,
		331, 5, 26, 0, 0, 331, 55, 1, 0, 0, 0, 332, 333, 5, 8, 0, 0, 333, 334,
		5, 25, 0, 0, 334, 335, 3, 36, 18, 0, 335, 336, 5, 26, 0, 0, 336, 57, 1,
		0, 0, 0, 337, 339, 3, 62, 31, 0, 338, 337, 1, 0, 0, 0, 339, 342, 1, 0,
		0, 0, 340, 338, 1, 0, 0, 0, 340, 341, 1, 0, 0, 0, 341, 59, 1, 0, 0, 0,
		342, 340, 1, 0, 0, 0, 343, 344, 5, 29, 0, 0, 344, 345, 3, 58, 29, 0, 345,
		346, 5, 30, 0, 0, 346, 61, 1, 0, 0, 0, 347, 348, 5, 9, 0, 0, 348, 350,
		5, 25, 0, 0, 349, 351, 3, 38, 19, 0, 350, 349, 1, 0, 0, 0, 350, 351, 1,
		0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 353, 5, 26, 0, 0, 353, 390, 5, 22,
		0, 0, 354, 355, 5, 10, 0, 0, 355, 357, 5, 25, 0, 0, 356, 358, 3, 38, 19,
		0, 357, 356, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359,
		360, 5, 26, 0, 0, 360, 390, 5, 22, 0, 0, 361, 363, 5, 11, 0, 0, 362, 364,
		3, 36, 18, 0, 363, 362, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 365, 1,
		0, 0, 0, 365, 390, 5, 22, 0, 0, 366, 367, 5, 12, 0, 0, 367, 390, 5, 22,
		0, 0, 368, 369, 5, 13, 0, 0, 369, 390, 5, 22, 0, 0, 370, 371, 3, 64, 32,
		0, 371, 372, 5, 22, 0, 0, 372, 390, 1, 0, 0, 0, 373, 374, 3, 60, 30, 0,
		374, 375, 5, 22, 0, 0, 375, 390, 1, 0, 0, 0, 376, 377, 3, 72, 36, 0, 377,
		378, 5, 22, 0, 0, 378, 390, 1, 0, 0, 0, 379, 380, 3, 68, 34, 0, 380, 381,
		5, 22, 0, 0, 381, 390, 1, 0, 0, 0, 382, 383, 3, 70, 35, 0, 383, 384, 5,
		22, 0, 0, 384, 390, 1, 0, 0, 0, 385, 386, 3, 12, 6, 0, 386, 387, 5, 22,
		0, 0, 387, 390, 1, 0, 0, 0, 388, 390, 3, 4, 2, 0, 389, 347, 1, 0, 0, 0,
		389, 354, 1, 0, 0, 0, 389, 361, 1, 0, 0, 0, 389, 366, 1, 0, 0, 0, 389,
		368, 1, 0, 0, 0, 389, 370, 1, 0, 0, 0, 389, 373, 1, 0, 0, 0, 389, 376,
		1, 0, 0, 0, 389, 379, 1, 0, 0, 0, 389, 382, 1, 0, 0, 0, 389, 385, 1, 0,
		0, 0, 389, 388, 1, 0, 0, 0, 390, 63, 1, 0, 0, 0, 391, 393, 3, 36, 18, 0,
		392, 394, 7, 2, 0, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0, 394,
		401, 1, 0, 0, 0, 395, 401, 3, 66, 33, 0, 396, 397, 3, 38, 19, 0, 397, 398,
		5, 21, 0, 0, 398, 399, 3, 38, 19, 0, 399, 401, 1, 0, 0, 0, 400, 391, 1,
		0, 0, 0, 400, 395, 1, 0, 0, 0, 400, 396, 1, 0, 0, 0, 401, 65, 1, 0, 0,
		0, 402, 403, 3, 38, 19, 0, 403, 404, 5, 20, 0, 0, 404, 405, 3, 38, 19,
		0, 405, 411, 1, 0, 0, 0, 406, 407, 3, 36, 18, 0, 407, 408, 7, 3, 0, 0,
		408, 409, 3, 36, 18, 0, 409, 411, 1, 0, 0, 0, 410, 402, 1, 0, 0, 0, 410,
		406, 1, 0, 0, 0, 411, 67, 1, 0, 0, 0, 412, 413, 5, 14, 0, 0, 413, 414,
		3, 36, 18, 0, 414, 415, 3, 60, 30, 0, 415, 451, 1, 0, 0, 0, 416, 417, 5,
		14, 0, 0, 417, 418, 3, 36, 18, 0, 418, 419, 3, 60, 30, 0, 419, 420, 5,
		15, 0, 0, 420, 421, 3, 68, 34, 0, 421, 451, 1, 0, 0, 0, 422, 423, 5, 14,
		0, 0, 423, 424, 3, 36, 18, 0, 424, 425, 3, 60, 30, 0, 425, 426, 5, 15,
		0, 0, 426, 427, 3, 60, 30, 0, 427, 451, 1, 0, 0, 0, 428, 429, 5, 14, 0,
		0, 429, 430, 3, 64, 32, 0, 430, 431, 5, 22, 0, 0, 431, 432, 3, 36, 18,
		0, 432, 433, 3, 60, 30, 0, 433, 451, 1, 0, 0, 0, 434, 435, 5, 14, 0, 0,
		435, 436, 3, 64, 32, 0, 436, 437, 5, 22, 0, 0, 437, 438, 3, 36, 18, 0,
		438, 439, 3, 60, 30, 0, 439, 440, 5, 15, 0, 0, 440, 441, 3, 68, 34, 0,
		441, 451, 1, 0, 0, 0, 442, 443, 5, 14, 0, 0, 443, 444, 3, 64, 32, 0, 444,
		445, 5, 22, 0, 0, 445, 446, 3, 36, 18, 0, 446, 447, 3, 60, 30, 0, 447,
		448, 5, 15, 0, 0, 448, 449, 3, 60, 30, 0, 449, 451, 1, 0, 0, 0, 450, 412,
		1, 0, 0, 0, 450, 416, 1, 0, 0, 0, 450, 422, 1, 0, 0, 0, 450, 428, 1, 0,
		0, 0, 450, 434, 1, 0, 0, 0, 450, 442, 1, 0, 0, 0, 451, 69, 1, 0, 0, 0,
		452, 453, 5, 16, 0, 0, 453, 474, 3, 60, 30, 0, 454, 455, 5, 16, 0, 0, 455,
		456, 3, 36, 18, 0, 456, 457, 3, 60, 30, 0, 457, 474, 1, 0, 0, 0, 458, 459,
		5, 16, 0, 0, 459, 460, 3, 64, 32, 0, 460, 461, 5, 22, 0, 0, 461, 462, 3,
		36, 18, 0, 462, 463, 5, 22, 0, 0, 463, 464, 3, 64, 32, 0, 464, 465, 3,
		60, 30, 0, 465, 474, 1, 0, 0, 0, 466, 467, 5, 16, 0, 0, 467, 468, 3, 64,
		32, 0, 468, 469, 5, 22, 0, 0, 469, 470, 5, 22, 0, 0, 470, 471, 3, 64, 32,
		0, 471, 472, 3, 60, 30, 0, 472, 474, 1, 0, 0, 0, 473, 452, 1, 0, 0, 0,
		473, 454, 1, 0, 0, 0, 473, 458, 1, 0, 0, 0, 473, 466, 1, 0, 0, 0, 474,
		71, 1, 0, 0, 0, 475, 476, 5, 17, 0, 0, 476, 477, 3, 64, 32, 0, 477, 478,
		5, 22, 0, 0, 478, 479, 3, 36, 18, 0, 479, 480, 5, 29, 0, 0, 480, 481, 3,
		74, 37, 0, 481, 482, 5, 30, 0, 0, 482, 502, 1, 0, 0, 0, 483, 484, 5, 17,
		0, 0, 484, 485, 3, 36, 18, 0, 485, 486, 5, 29, 0, 0, 486, 487, 3, 74, 37,
		0, 487, 488, 5, 30, 0, 0, 488, 502, 1, 0, 0, 0, 489, 490, 5, 17, 0, 0,
		490, 491, 3, 64, 32, 0, 491, 492, 5, 22, 0, 0, 492, 493, 5, 29, 0, 0, 493,
		494, 3, 74, 37, 0, 494, 495, 5, 30, 0, 0, 495, 502, 1, 0, 0, 0, 496, 497,
		5, 17, 0, 0, 497, 498, 5, 29, 0, 0, 498, 499, 3, 74, 37, 0, 499, 500, 5,
		30, 0, 0, 500, 502, 1, 0, 0, 0, 501, 475, 1, 0, 0, 0, 501, 483, 1, 0, 0,
		0, 501, 489, 1, 0, 0, 0, 501, 496, 1, 0, 0, 0, 502, 73, 1, 0, 0, 0, 503,
		504, 3, 76, 38, 0, 504, 505, 3, 74, 37, 0, 505, 507, 1, 0, 0, 0, 506, 503,
		1, 0, 0, 0, 507, 510, 1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 508, 509, 1, 0,
		0, 0, 509, 75, 1, 0, 0, 0, 510, 508, 1, 0, 0, 0, 511, 512, 3, 78, 39, 0,
		512, 513, 5, 23, 0, 0, 513, 514, 3, 58, 29, 0, 514, 77, 1, 0, 0, 0, 515,
		516, 5, 18, 0, 0, 516, 519, 3, 38, 19, 0, 517, 519, 5, 19, 0, 0, 518, 515,
		1, 0, 0, 0, 518, 517, 1, 0, 0, 0, 519, 79, 1, 0, 0, 0, 35, 88, 90, 107,
		116, 129, 148, 157, 192, 199, 213, 227, 238, 246, 253, 260, 268, 276, 284,
		286, 295, 302, 312, 340, 350, 357, 363, 389, 393, 400, 410, 450, 473, 501,
		508, 518,
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

// expr_parserInit initializes any static state used to implement expr_parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newexpr_parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Expr_parserInit() {
	staticData := &Expr_parserParserStaticData
	staticData.once.Do(expr_parserParserInit)
}

// Newexpr_parser produces a new parser instance for the optional input antlr.TokenStream.
func Newexpr_parser(input antlr.TokenStream) *expr_parser {
	Expr_parserInit()
	this := new(expr_parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Expr_parserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "expr_parser.g4"

	return this
}

// expr_parser tokens.
const (
	expr_parserEOF                    = antlr.TokenEOF
	expr_parserPACKAGE                = 1
	expr_parserVAR                    = 2
	expr_parserTYPE                   = 3
	expr_parserFUNCTION               = 4
	expr_parserSTRUCT                 = 5
	expr_parserAPPEND                 = 6
	expr_parserLENGHT                 = 7
	expr_parserCAP                    = 8
	expr_parserPRINT                  = 9
	expr_parserPRINTLN                = 10
	expr_parserRETURN                 = 11
	expr_parserBREAK                  = 12
	expr_parserCONTINUE               = 13
	expr_parserIF                     = 14
	expr_parserELSE                   = 15
	expr_parserFOR                    = 16
	expr_parserSWITCH                 = 17
	expr_parserCASE                   = 18
	expr_parserDEFAULT                = 19
	expr_parserASSIGN                 = 20
	expr_parserSHORT_DEC              = 21
	expr_parserSEMICOLON              = 22
	expr_parserCOLON                  = 23
	expr_parserDOT                    = 24
	expr_parserLPAREN                 = 25
	expr_parserRPAREN                 = 26
	expr_parserRBRACK                 = 27
	expr_parserLBRACK                 = 28
	expr_parserLBRACE                 = 29
	expr_parserRBRACE                 = 30
	expr_parserCOMMA                  = 31
	expr_parserADD                    = 32
	expr_parserSUB                    = 33
	expr_parserMULT                   = 34
	expr_parserDIV                    = 35
	expr_parserMOD                    = 36
	expr_parserLSHIFT                 = 37
	expr_parserRSHIFT                 = 38
	expr_parserAND                    = 39
	expr_parserANDNOT                 = 40
	expr_parserPIPE                   = 41
	expr_parserCARET                  = 42
	expr_parserEQUALS                 = 43
	expr_parserNOT_EQ                 = 44
	expr_parserLESS_THAN              = 45
	expr_parserLESS_THAN_OR_EQUALS    = 46
	expr_parserGREATER_THAN           = 47
	expr_parserGREATER_THAN_OR_EQUALS = 48
	expr_parserLOG_AND                = 49
	expr_parserLOG_OR                 = 50
	expr_parserLOG_NOT                = 51
	expr_parserINCREMENT              = 52
	expr_parserDECREMENT              = 53
	expr_parserPLUS_ASSIGN            = 54
	expr_parserAND_ASSIGN             = 55
	expr_parserMINUS_ASSIGN           = 56
	expr_parserOR_ASSIGN              = 57
	expr_parserTIMES_ASSIGN           = 58
	expr_parserXOR_ASSIGN             = 59
	expr_parserLEFT_SHIFT_ASSIGN      = 60
	expr_parserRIGHT_SHIFT_ASSIGN     = 61
	expr_parserAND_NOT_ASSIGN         = 62
	expr_parserMOD_ASSIGN             = 63
	expr_parserDIVIDE_ASSIGN          = 64
	expr_parserINT_LIT                = 65
	expr_parserIDENTIFIER             = 66
	expr_parserFLOAT_LIT              = 67
	expr_parserRUNE_LIT               = 68
	expr_parserRAW_STRING_LIT         = 69
	expr_parserINTERPRETED_STRING_LIT = 70
	expr_parserSTRING_LIT             = 71
	expr_parserWS                     = 72
)

// expr_parser rules.
const (
	expr_parserRULE_root                     = 0
	expr_parserRULE_topDeclarationList       = 1
	expr_parserRULE_variableDecl             = 2
	expr_parserRULE_innerVarDecls            = 3
	expr_parserRULE_singleVarDecl            = 4
	expr_parserRULE_singleVarDeclNoExps      = 5
	expr_parserRULE_typeDecl                 = 6
	expr_parserRULE_innerTypeDecls           = 7
	expr_parserRULE_singleTypeDecl           = 8
	expr_parserRULE_funcDecl                 = 9
	expr_parserRULE_funcFrontDecl            = 10
	expr_parserRULE_funcArgDecls             = 11
	expr_parserRULE_declType                 = 12
	expr_parserRULE_sliceDeclType            = 13
	expr_parserRULE_arrayDeclType            = 14
	expr_parserRULE_structDeclType           = 15
	expr_parserRULE_structMemDecls           = 16
	expr_parserRULE_identifierList           = 17
	expr_parserRULE_expression               = 18
	expr_parserRULE_expressionList           = 19
	expr_parserRULE_primaryExpression        = 20
	expr_parserRULE_operand                  = 21
	expr_parserRULE_literal                  = 22
	expr_parserRULE_index                    = 23
	expr_parserRULE_arguments                = 24
	expr_parserRULE_selector                 = 25
	expr_parserRULE_appendExpression         = 26
	expr_parserRULE_lengthExpression         = 27
	expr_parserRULE_capExpression            = 28
	expr_parserRULE_statementList            = 29
	expr_parserRULE_block                    = 30
	expr_parserRULE_statement                = 31
	expr_parserRULE_simpleStatement          = 32
	expr_parserRULE_assignmentStatement      = 33
	expr_parserRULE_ifStatement              = 34
	expr_parserRULE_loop                     = 35
	expr_parserRULE_switch                   = 36
	expr_parserRULE_expressionCaseClauseList = 37
	expr_parserRULE_expressionCaseClause     = 38
	expr_parserRULE_expressionSwitchCase     = 39
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	TopDeclarationList() ITopDeclarationListContext

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_root
	return p
}

func InitEmptyRootContext(p *RootContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_root
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(expr_parserPACKAGE, 0)
}

func (s *RootContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *RootContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *RootContext) TopDeclarationList() ITopDeclarationListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopDeclarationListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITopDeclarationListContext)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitRoot(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Root() (localctx IRootContext) {
	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, expr_parserRULE_root)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(expr_parserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.Match(expr_parserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(expr_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.TopDeclarationList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITopDeclarationListContext is an interface to support dynamic dispatch.
type ITopDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVariableDecl() []IVariableDeclContext
	VariableDecl(i int) IVariableDeclContext
	AllTypeDecl() []ITypeDeclContext
	TypeDecl(i int) ITypeDeclContext
	AllFuncDecl() []IFuncDeclContext
	FuncDecl(i int) IFuncDeclContext

	// IsTopDeclarationListContext differentiates from other interfaces.
	IsTopDeclarationListContext()
}

type TopDeclarationListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopDeclarationListContext() *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_topDeclarationList
	return p
}

func InitEmptyTopDeclarationListContext(p *TopDeclarationListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_topDeclarationList
}

func (*TopDeclarationListContext) IsTopDeclarationListContext() {}

func NewTopDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopDeclarationListContext {
	var p = new(TopDeclarationListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_topDeclarationList

	return p
}

func (s *TopDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *TopDeclarationListContext) AllVariableDecl() []IVariableDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVariableDeclContext); ok {
			len++
		}
	}

	tst := make([]IVariableDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVariableDeclContext); ok {
			tst[i] = t.(IVariableDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) VariableDecl(i int) IVariableDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *TopDeclarationListContext) AllTypeDecl() []ITypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ITypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDeclContext); ok {
			tst[i] = t.(ITypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) TypeDecl(i int) ITypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TopDeclarationListContext) AllFuncDecl() []IFuncDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFuncDeclContext); ok {
			len++
		}
	}

	tst := make([]IFuncDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFuncDeclContext); ok {
			tst[i] = t.(IFuncDeclContext)
			i++
		}
	}

	return tst
}

func (s *TopDeclarationListContext) FuncDecl(i int) IFuncDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncDeclContext)
}

func (s *TopDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopDeclarationListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitTopDeclarationList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) TopDeclarationList() (localctx ITopDeclarationListContext) {
	localctx = NewTopDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, expr_parserRULE_topDeclarationList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&28) != 0 {
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case expr_parserVAR:
			{
				p.SetState(85)
				p.VariableDecl()
			}

		case expr_parserTYPE:
			{
				p.SetState(86)
				p.TypeDecl()
			}

		case expr_parserFUNCTION:
			{
				p.SetState(87)
				p.FuncDecl()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableDeclContext is an interface to support dynamic dispatch.
type IVariableDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVariableDeclContext differentiates from other interfaces.
	IsVariableDeclContext()
}

type VariableDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclContext() *VariableDeclContext {
	var p = new(VariableDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_variableDecl
	return p
}

func InitEmptyVariableDeclContext(p *VariableDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_variableDecl
}

func (*VariableDeclContext) IsVariableDeclContext() {}

func NewVariableDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclContext {
	var p = new(VariableDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_variableDecl

	return p
}

func (s *VariableDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclContext) CopyAll(ctx *VariableDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VariableDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MVarDeclContext struct {
	VariableDeclContext
}

func NewMVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MVarDeclContext {
	var p = new(MVarDeclContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *MVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MVarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(expr_parserVAR, 0)
}

func (s *MVarDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *MVarDeclContext) InnerVarDecls() IInnerVarDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerVarDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerVarDeclsContext)
}

func (s *MVarDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *MVarDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *MVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitMVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type VoidVarDeclContext struct {
	VariableDeclContext
}

func NewVoidVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidVarDeclContext {
	var p = new(VoidVarDeclContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VoidVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidVarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(expr_parserVAR, 0)
}

func (s *VoidVarDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *VoidVarDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *VoidVarDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *VoidVarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVoidVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclContext struct {
	VariableDeclContext
}

func NewVarDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclContext {
	var p = new(VarDeclContext)

	InitEmptyVariableDeclContext(&p.VariableDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*VariableDeclContext))

	return p
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) VAR() antlr.TerminalNode {
	return s.GetToken(expr_parserVAR, 0)
}

func (s *VarDeclContext) SingleVarDecl() ISingleVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *VarDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) VariableDecl() (localctx IVariableDeclContext) {
	localctx = NewVariableDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, expr_parserRULE_variableDecl)
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(expr_parserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(94)
			p.SingleVarDecl()
		}
		{
			p.SetState(95)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.Match(expr_parserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(98)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.InnerVarDecls()
		}
		{
			p.SetState(100)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVoidVarDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(103)
			p.Match(expr_parserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerVarDeclsContext is an interface to support dynamic dispatch.
type IInnerVarDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDecl() []ISingleVarDeclContext
	SingleVarDecl(i int) ISingleVarDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsInnerVarDeclsContext differentiates from other interfaces.
	IsInnerVarDeclsContext()
}

type InnerVarDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerVarDeclsContext() *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_innerVarDecls
	return p
}

func InitEmptyInnerVarDeclsContext(p *InnerVarDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_innerVarDecls
}

func (*InnerVarDeclsContext) IsInnerVarDeclsContext() {}

func NewInnerVarDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerVarDeclsContext {
	var p = new(InnerVarDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_innerVarDecls

	return p
}

func (s *InnerVarDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerVarDeclsContext) AllSingleVarDecl() []ISingleVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclContext); ok {
			tst[i] = t.(ISingleVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerVarDeclsContext) SingleVarDecl(i int) ISingleVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclContext)
}

func (s *InnerVarDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(expr_parserSEMICOLON)
}

func (s *InnerVarDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, i)
}

func (s *InnerVarDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerVarDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerVarDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitInnerVarDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) InnerVarDecls() (localctx IInnerVarDeclsContext) {
	localctx = NewInnerVarDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, expr_parserRULE_innerVarDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.SingleVarDecl()
	}
	{
		p.SetState(110)
		p.Match(expr_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == expr_parserIDENTIFIER {
		{
			p.SetState(111)
			p.SingleVarDecl()
		}
		{
			p.SetState(112)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclContext is an interface to support dynamic dispatch.
type ISingleVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSingleVarDeclContext differentiates from other interfaces.
	IsSingleVarDeclContext()
}

type SingleVarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclContext() *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleVarDecl
	return p
}

func InitEmptySingleVarDeclContext(p *SingleVarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleVarDecl
}

func (*SingleVarDeclContext) IsSingleVarDeclContext() {}

func NewSingleVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclContext {
	var p = new(SingleVarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_singleVarDecl

	return p
}

func (s *SingleVarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclContext) CopyAll(ctx *SingleVarDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SingleVarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarDeclWithoutTypeContext struct {
	SingleVarDeclContext
}

func NewVarDeclWithoutTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclWithoutTypeContext {
	var p = new(VarDeclWithoutTypeContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *VarDeclWithoutTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclWithoutTypeContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *VarDeclWithoutTypeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserASSIGN, 0)
}

func (s *VarDeclWithoutTypeContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *VarDeclWithoutTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVarDeclWithoutType(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclWithTypeContext struct {
	SingleVarDeclContext
}

func NewVarDeclWithTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclWithTypeContext {
	var p = new(VarDeclWithTypeContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *VarDeclWithTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclWithTypeContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *VarDeclWithTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *VarDeclWithTypeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserASSIGN, 0)
}

func (s *VarDeclWithTypeContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *VarDeclWithTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVarDeclWithType(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclNoExpContext struct {
	SingleVarDeclContext
}

func NewVarDeclNoExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclNoExpContext {
	var p = new(VarDeclNoExpContext)

	InitEmptySingleVarDeclContext(&p.SingleVarDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*SingleVarDeclContext))

	return p
}

func (s *VarDeclNoExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclNoExpContext) SingleVarDeclNoExps() ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *VarDeclNoExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVarDeclNoExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) SingleVarDecl() (localctx ISingleVarDeclContext) {
	localctx = NewSingleVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, expr_parserRULE_singleVarDecl)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDeclWithTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.IdentifierList()
		}
		{
			p.SetState(120)
			p.DeclType()
		}
		{
			p.SetState(121)
			p.Match(expr_parserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.ExpressionList()
		}

	case 2:
		localctx = NewVarDeclWithoutTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.IdentifierList()
		}
		{
			p.SetState(125)
			p.Match(expr_parserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.ExpressionList()
		}

	case 3:
		localctx = NewVarDeclNoExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
			p.SingleVarDeclNoExps()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleVarDeclNoExpsContext is an interface to support dynamic dispatch.
type ISingleVarDeclNoExpsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierList() IIdentifierListContext
	DeclType() IDeclTypeContext

	// IsSingleVarDeclNoExpsContext differentiates from other interfaces.
	IsSingleVarDeclNoExpsContext()
}

type SingleVarDeclNoExpsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleVarDeclNoExpsContext() *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleVarDeclNoExps
	return p
}

func InitEmptySingleVarDeclNoExpsContext(p *SingleVarDeclNoExpsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleVarDeclNoExps
}

func (*SingleVarDeclNoExpsContext) IsSingleVarDeclNoExpsContext() {}

func NewSingleVarDeclNoExpsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleVarDeclNoExpsContext {
	var p = new(SingleVarDeclNoExpsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_singleVarDeclNoExps

	return p
}

func (s *SingleVarDeclNoExpsContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleVarDeclNoExpsContext) IdentifierList() IIdentifierListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierListContext)
}

func (s *SingleVarDeclNoExpsContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleVarDeclNoExpsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleVarDeclNoExpsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleVarDeclNoExpsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSingleVarDeclNoExps(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) SingleVarDeclNoExps() (localctx ISingleVarDeclNoExpsContext) {
	localctx = NewSingleVarDeclNoExpsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, expr_parserRULE_singleVarDeclNoExps)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.IdentifierList()
	}
	{
		p.SetState(132)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_typeDecl
	return p
}

func InitEmptyTypeDeclContext(p *TypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_typeDecl
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) CopyAll(ctx *TypeDeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VoidTypeDeclContext struct {
	TypeDeclContext
}

func NewVoidTypeDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidTypeDeclContext {
	var p = new(VoidTypeDeclContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *VoidTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidTypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(expr_parserTYPE, 0)
}

func (s *VoidTypeDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *VoidTypeDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *VoidTypeDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *VoidTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVoidTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type MultiTypeDeclContext struct {
	TypeDeclContext
}

func NewMultiTypeDeclContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiTypeDeclContext {
	var p = new(MultiTypeDeclContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *MultiTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiTypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(expr_parserTYPE, 0)
}

func (s *MultiTypeDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *MultiTypeDeclContext) InnerTypeDecls() IInnerTypeDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInnerTypeDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInnerTypeDeclsContext)
}

func (s *MultiTypeDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *MultiTypeDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *MultiTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitMultiTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDecContext struct {
	TypeDeclContext
}

func NewTypeDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDecContext {
	var p = new(TypeDecContext)

	InitEmptyTypeDeclContext(&p.TypeDeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeDeclContext))

	return p
}

func (s *TypeDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDecContext) TYPE() antlr.TerminalNode {
	return s.GetToken(expr_parserTYPE, 0)
}

func (s *TypeDecContext) SingleTypeDecl() ISingleTypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *TypeDecContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *TypeDecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitTypeDec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, expr_parserRULE_typeDecl)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeDecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(134)
			p.Match(expr_parserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.SingleTypeDecl()
		}
		{
			p.SetState(136)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewMultiTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(expr_parserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.InnerTypeDecls()
		}
		{
			p.SetState(141)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVoidTypeDeclContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Match(expr_parserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInnerTypeDeclsContext is an interface to support dynamic dispatch.
type IInnerTypeDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleTypeDecl() []ISingleTypeDeclContext
	SingleTypeDecl(i int) ISingleTypeDeclContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsInnerTypeDeclsContext differentiates from other interfaces.
	IsInnerTypeDeclsContext()
}

type InnerTypeDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInnerTypeDeclsContext() *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_innerTypeDecls
	return p
}

func InitEmptyInnerTypeDeclsContext(p *InnerTypeDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_innerTypeDecls
}

func (*InnerTypeDeclsContext) IsInnerTypeDeclsContext() {}

func NewInnerTypeDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InnerTypeDeclsContext {
	var p = new(InnerTypeDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_innerTypeDecls

	return p
}

func (s *InnerTypeDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *InnerTypeDeclsContext) AllSingleTypeDecl() []ISingleTypeDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			len++
		}
	}

	tst := make([]ISingleTypeDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleTypeDeclContext); ok {
			tst[i] = t.(ISingleTypeDeclContext)
			i++
		}
	}

	return tst
}

func (s *InnerTypeDeclsContext) SingleTypeDecl(i int) ISingleTypeDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleTypeDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleTypeDeclContext)
}

func (s *InnerTypeDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(expr_parserSEMICOLON)
}

func (s *InnerTypeDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, i)
}

func (s *InnerTypeDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InnerTypeDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InnerTypeDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitInnerTypeDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) InnerTypeDecls() (localctx IInnerTypeDeclsContext) {
	localctx = NewInnerTypeDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, expr_parserRULE_innerTypeDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.SingleTypeDecl()
	}
	{
		p.SetState(151)
		p.Match(expr_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == expr_parserIDENTIFIER {
		{
			p.SetState(152)
			p.SingleTypeDecl()
		}
		{
			p.SetState(153)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISingleTypeDeclContext is an interface to support dynamic dispatch.
type ISingleTypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsSingleTypeDeclContext differentiates from other interfaces.
	IsSingleTypeDeclContext()
}

type SingleTypeDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTypeDeclContext() *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleTypeDecl
	return p
}

func InitEmptySingleTypeDeclContext(p *SingleTypeDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_singleTypeDecl
}

func (*SingleTypeDeclContext) IsSingleTypeDeclContext() {}

func NewSingleTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTypeDeclContext {
	var p = new(SingleTypeDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_singleTypeDecl

	return p
}

func (s *SingleTypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTypeDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *SingleTypeDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SingleTypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleTypeDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSingleTypeDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) SingleTypeDecl() (localctx ISingleTypeDeclContext) {
	localctx = NewSingleTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, expr_parserRULE_singleTypeDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(expr_parserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncDeclContext is an interface to support dynamic dispatch.
type IFuncDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncFrontDecl() IFuncFrontDeclContext
	Block() IBlockContext
	SEMICOLON() antlr.TerminalNode

	// IsFuncDeclContext differentiates from other interfaces.
	IsFuncDeclContext()
}

type FuncDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDeclContext() *FuncDeclContext {
	var p = new(FuncDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcDecl
	return p
}

func InitEmptyFuncDeclContext(p *FuncDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcDecl
}

func (*FuncDeclContext) IsFuncDeclContext() {}

func NewFuncDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDeclContext {
	var p = new(FuncDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_funcDecl

	return p
}

func (s *FuncDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDeclContext) FuncFrontDecl() IFuncFrontDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncFrontDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncFrontDeclContext)
}

func (s *FuncDeclContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *FuncDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitFuncDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) FuncDecl() (localctx IFuncDeclContext) {
	localctx = NewFuncDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, expr_parserRULE_funcDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.FuncFrontDecl()
	}
	{
		p.SetState(164)
		p.Block()
	}
	{
		p.SetState(165)
		p.Match(expr_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncFrontDeclContext is an interface to support dynamic dispatch.
type IFuncFrontDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	FuncArgDecls() IFuncArgDeclsContext
	RPAREN() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsFuncFrontDeclContext differentiates from other interfaces.
	IsFuncFrontDeclContext()
}

type FuncFrontDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncFrontDeclContext() *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcFrontDecl
	return p
}

func InitEmptyFuncFrontDeclContext(p *FuncFrontDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcFrontDecl
}

func (*FuncFrontDeclContext) IsFuncFrontDeclContext() {}

func NewFuncFrontDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncFrontDeclContext {
	var p = new(FuncFrontDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_funcFrontDecl

	return p
}

func (s *FuncFrontDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncFrontDeclContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(expr_parserFUNCTION, 0)
}

func (s *FuncFrontDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *FuncFrontDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *FuncFrontDeclContext) FuncArgDecls() IFuncArgDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncArgDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncArgDeclsContext)
}

func (s *FuncFrontDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *FuncFrontDeclContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *FuncFrontDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncFrontDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncFrontDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitFuncFrontDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) FuncFrontDecl() (localctx IFuncFrontDeclContext) {
	localctx = NewFuncFrontDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, expr_parserRULE_funcFrontDecl)
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(167)
			p.Match(expr_parserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.FuncArgDecls()
		}
		{
			p.SetState(171)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.DeclType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
			p.Match(expr_parserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(176)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.FuncArgDecls()
		}
		{
			p.SetState(178)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.Match(expr_parserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(182)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.DeclType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(185)
			p.Match(expr_parserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(189)
			p.Match(expr_parserFUNCTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.DeclType()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncArgDeclsContext is an interface to support dynamic dispatch.
type IFuncArgDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFuncArgDeclsContext differentiates from other interfaces.
	IsFuncArgDeclsContext()
}

type FuncArgDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgDeclsContext() *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcArgDecls
	return p
}

func InitEmptyFuncArgDeclsContext(p *FuncArgDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_funcArgDecls
}

func (*FuncArgDeclsContext) IsFuncArgDeclsContext() {}

func NewFuncArgDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgDeclsContext {
	var p = new(FuncArgDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_funcArgDecls

	return p
}

func (s *FuncArgDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *FuncArgDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *FuncArgDeclsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expr_parserCOMMA)
}

func (s *FuncArgDeclsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserCOMMA, i)
}

func (s *FuncArgDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitFuncArgDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) FuncArgDecls() (localctx IFuncArgDeclsContext) {
	localctx = NewFuncArgDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, expr_parserRULE_funcArgDecls)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(194)
				p.SingleVarDeclNoExps()
			}
			{
				p.SetState(195)
				p.Match(expr_parserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.SingleVarDeclNoExps()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclTypeContext is an interface to support dynamic dispatch.
type IDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	DeclType() IDeclTypeContext
	RPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	SliceDeclType() ISliceDeclTypeContext
	ArrayDeclType() IArrayDeclTypeContext
	StructDeclType() IStructDeclTypeContext

	// IsDeclTypeContext differentiates from other interfaces.
	IsDeclTypeContext()
}

type DeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_declType
	return p
}

func InitEmptyDeclTypeContext(p *DeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_declType
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *DeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *DeclTypeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *DeclTypeContext) SliceDeclType() ISliceDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISliceDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISliceDeclTypeContext)
}

func (s *DeclTypeContext) ArrayDeclType() IArrayDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayDeclTypeContext)
}

func (s *DeclTypeContext) StructDeclType() IStructDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclTypeContext)
}

func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, expr_parserRULE_declType)
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(204)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.DeclType()
		}
		{
			p.SetState(206)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.SliceDeclType()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.ArrayDeclType()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(211)
			p.StructDeclType()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISliceDeclTypeContext is an interface to support dynamic dispatch.
type ISliceDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsSliceDeclTypeContext differentiates from other interfaces.
	IsSliceDeclTypeContext()
}

type SliceDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySliceDeclTypeContext() *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_sliceDeclType
	return p
}

func InitEmptySliceDeclTypeContext(p *SliceDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_sliceDeclType
}

func (*SliceDeclTypeContext) IsSliceDeclTypeContext() {}

func NewSliceDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SliceDeclTypeContext {
	var p = new(SliceDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_sliceDeclType

	return p
}

func (s *SliceDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SliceDeclTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACK, 0)
}

func (s *SliceDeclTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACK, 0)
}

func (s *SliceDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *SliceDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SliceDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSliceDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) SliceDeclType() (localctx ISliceDeclTypeContext) {
	localctx = NewSliceDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, expr_parserRULE_sliceDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(expr_parserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(expr_parserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayDeclTypeContext is an interface to support dynamic dispatch.
type IArrayDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode
	RBRACK() antlr.TerminalNode
	DeclType() IDeclTypeContext

	// IsArrayDeclTypeContext differentiates from other interfaces.
	IsArrayDeclTypeContext()
}

type ArrayDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayDeclTypeContext() *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_arrayDeclType
	return p
}

func InitEmptyArrayDeclTypeContext(p *ArrayDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_arrayDeclType
}

func (*ArrayDeclTypeContext) IsArrayDeclTypeContext() {}

func NewArrayDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayDeclTypeContext {
	var p = new(ArrayDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_arrayDeclType

	return p
}

func (s *ArrayDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayDeclTypeContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACK, 0)
}

func (s *ArrayDeclTypeContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserINT_LIT, 0)
}

func (s *ArrayDeclTypeContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACK, 0)
}

func (s *ArrayDeclTypeContext) DeclType() IDeclTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *ArrayDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitArrayDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) ArrayDeclType() (localctx IArrayDeclTypeContext) {
	localctx = NewArrayDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, expr_parserRULE_arrayDeclType)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(expr_parserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(expr_parserINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(expr_parserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
		p.DeclType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructDeclTypeContext is an interface to support dynamic dispatch.
type IStructDeclTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRUCT() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	StructMemDecls() IStructMemDeclsContext

	// IsStructDeclTypeContext differentiates from other interfaces.
	IsStructDeclTypeContext()
}

type StructDeclTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclTypeContext() *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_structDeclType
	return p
}

func InitEmptyStructDeclTypeContext(p *StructDeclTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_structDeclType
}

func (*StructDeclTypeContext) IsStructDeclTypeContext() {}

func NewStructDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclTypeContext {
	var p = new(StructDeclTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_structDeclType

	return p
}

func (s *StructDeclTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclTypeContext) STRUCT() antlr.TerminalNode {
	return s.GetToken(expr_parserSTRUCT, 0)
}

func (s *StructDeclTypeContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *StructDeclTypeContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *StructDeclTypeContext) StructMemDecls() IStructMemDeclsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructMemDeclsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructMemDeclsContext)
}

func (s *StructDeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitStructDeclType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) StructDeclType() (localctx IStructDeclTypeContext) {
	localctx = NewStructDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, expr_parserRULE_structDeclType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(expr_parserSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(expr_parserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == expr_parserIDENTIFIER {
		{
			p.SetState(226)
			p.StructMemDecls()
		}

	}
	{
		p.SetState(229)
		p.Match(expr_parserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStructMemDeclsContext is an interface to support dynamic dispatch.
type IStructMemDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext
	SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsStructMemDeclsContext differentiates from other interfaces.
	IsStructMemDeclsContext()
}

type StructMemDeclsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructMemDeclsContext() *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_structMemDecls
	return p
}

func InitEmptyStructMemDeclsContext(p *StructMemDeclsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_structMemDecls
}

func (*StructMemDeclsContext) IsStructMemDeclsContext() {}

func NewStructMemDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructMemDeclsContext {
	var p = new(StructMemDeclsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_structMemDecls

	return p
}

func (s *StructMemDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *StructMemDeclsContext) AllSingleVarDeclNoExps() []ISingleVarDeclNoExpsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			len++
		}
	}

	tst := make([]ISingleVarDeclNoExpsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			tst[i] = t.(ISingleVarDeclNoExpsContext)
			i++
		}
	}

	return tst
}

func (s *StructMemDeclsContext) SingleVarDeclNoExps(i int) ISingleVarDeclNoExpsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleVarDeclNoExpsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleVarDeclNoExpsContext)
}

func (s *StructMemDeclsContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(expr_parserSEMICOLON)
}

func (s *StructMemDeclsContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, i)
}

func (s *StructMemDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructMemDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructMemDeclsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitStructMemDecls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) StructMemDecls() (localctx IStructMemDeclsContext) {
	localctx = NewStructMemDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, expr_parserRULE_structMemDecls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.SingleVarDeclNoExps()
	}
	{
		p.SetState(232)
		p.Match(expr_parserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == expr_parserIDENTIFIER {
		{
			p.SetState(233)
			p.SingleVarDeclNoExps()
		}
		{
			p.SetState(234)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierListContext is an interface to support dynamic dispatch.
type IIdentifierListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdentifierListContext differentiates from other interfaces.
	IsIdentifierListContext()
}

type IdentifierListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierListContext() *IdentifierListContext {
	var p = new(IdentifierListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_identifierList
	return p
}

func InitEmptyIdentifierListContext(p *IdentifierListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_identifierList
}

func (*IdentifierListContext) IsIdentifierListContext() {}

func NewIdentifierListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierListContext {
	var p = new(IdentifierListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_identifierList

	return p
}

func (s *IdentifierListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierListContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(expr_parserIDENTIFIER)
}

func (s *IdentifierListContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, i)
}

func (s *IdentifierListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expr_parserCOMMA)
}

func (s *IdentifierListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserCOMMA, i)
}

func (s *IdentifierListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIdentifierList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) IdentifierList() (localctx IIdentifierListContext) {
	localctx = NewIdentifierListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, expr_parserRULE_identifierList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(expr_parserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(242)
				p.Match(expr_parserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(243)
				p.Match(expr_parserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpUnaryContext struct {
	ExpressionContext
}

func NewExpUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpUnaryContext {
	var p = new(ExpUnaryContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpUnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpUnaryContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpUnaryContext) ADD() antlr.TerminalNode {
	return s.GetToken(expr_parserADD, 0)
}

func (s *ExpUnaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(expr_parserSUB, 0)
}

func (s *ExpUnaryContext) LOG_NOT() antlr.TerminalNode {
	return s.GetToken(expr_parserLOG_NOT, 0)
}

func (s *ExpUnaryContext) CARET() antlr.TerminalNode {
	return s.GetToken(expr_parserCARET, 0)
}

func (s *ExpUnaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpUnary(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpPrimaryExpContext struct {
	ExpressionContext
}

func NewExpPrimaryExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpPrimaryExpContext {
	var p = new(ExpPrimaryExpContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpPrimaryExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpPrimaryExpContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpPrimaryExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpPrimaryExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpBinaryContext struct {
	ExpressionContext
}

func NewExpBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpBinaryContext {
	var p = new(ExpBinaryContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ExpBinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpBinaryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpBinaryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpBinaryContext) MULT() antlr.TerminalNode {
	return s.GetToken(expr_parserMULT, 0)
}

func (s *ExpBinaryContext) DIV() antlr.TerminalNode {
	return s.GetToken(expr_parserDIV, 0)
}

func (s *ExpBinaryContext) MOD() antlr.TerminalNode {
	return s.GetToken(expr_parserMOD, 0)
}

func (s *ExpBinaryContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(expr_parserLSHIFT, 0)
}

func (s *ExpBinaryContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(expr_parserRSHIFT, 0)
}

func (s *ExpBinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(expr_parserAND, 0)
}

func (s *ExpBinaryContext) ANDNOT() antlr.TerminalNode {
	return s.GetToken(expr_parserANDNOT, 0)
}

func (s *ExpBinaryContext) ADD() antlr.TerminalNode {
	return s.GetToken(expr_parserADD, 0)
}

func (s *ExpBinaryContext) SUB() antlr.TerminalNode {
	return s.GetToken(expr_parserSUB, 0)
}

func (s *ExpBinaryContext) PIPE() antlr.TerminalNode {
	return s.GetToken(expr_parserPIPE, 0)
}

func (s *ExpBinaryContext) CARET() antlr.TerminalNode {
	return s.GetToken(expr_parserCARET, 0)
}

func (s *ExpBinaryContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(expr_parserEQUALS, 0)
}

func (s *ExpBinaryContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(expr_parserNOT_EQ, 0)
}

func (s *ExpBinaryContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(expr_parserLESS_THAN, 0)
}

func (s *ExpBinaryContext) LESS_THAN_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(expr_parserLESS_THAN_OR_EQUALS, 0)
}

func (s *ExpBinaryContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(expr_parserGREATER_THAN, 0)
}

func (s *ExpBinaryContext) GREATER_THAN_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(expr_parserGREATER_THAN_OR_EQUALS, 0)
}

func (s *ExpBinaryContext) LOG_AND() antlr.TerminalNode {
	return s.GetToken(expr_parserLOG_AND, 0)
}

func (s *ExpBinaryContext) LOG_OR() antlr.TerminalNode {
	return s.GetToken(expr_parserLOG_OR, 0)
}

func (s *ExpBinaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpBinary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *expr_parser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 36
	p.EnterRecursionRule(localctx, 36, expr_parserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserAPPEND, expr_parserLENGHT, expr_parserCAP, expr_parserLPAREN, expr_parserINT_LIT, expr_parserIDENTIFIER, expr_parserFLOAT_LIT, expr_parserRUNE_LIT, expr_parserRAW_STRING_LIT, expr_parserINTERPRETED_STRING_LIT:
		localctx = NewExpPrimaryExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(250)
			p.primaryExpression(0)
		}

	case expr_parserADD, expr_parserSUB, expr_parserCARET, expr_parserLOG_NOT:
		localctx = NewExpUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(251)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256210745098240) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(252)
			p.expression(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpBinaryContext(p, NewExpressionContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, expr_parserRULE_expression)
			p.SetState(255)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(256)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251795518717952) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(257)
				p.expression(3)
			}

		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(expr_parserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, expr_parserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.expression(0)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == expr_parserCOMMA {
		{
			p.SetState(264)
			p.Match(expr_parserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(265)
			p.expression(0)
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) CopyAll(ctx *PrimaryExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CapExpContext struct {
	PrimaryExpressionContext
}

func NewCapExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapExpContext {
	var p = new(CapExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *CapExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpContext) CapExpression() ICapExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICapExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICapExpressionContext)
}

func (s *CapExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitCapExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AppendExpContext struct {
	PrimaryExpressionContext
}

func NewAppendExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AppendExpContext {
	var p = new(AppendExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *AppendExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpContext) AppendExpression() IAppendExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAppendExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAppendExpressionContext)
}

func (s *AppendExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitAppendExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type LenExpContext struct {
	PrimaryExpressionContext
}

func NewLenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LenExpContext {
	var p = new(LenExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *LenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LenExpContext) LengthExpression() ILengthExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILengthExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILengthExpressionContext)
}

func (s *LenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitLenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectorExpContext struct {
	PrimaryExpressionContext
}

func NewSelectorExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectorExpContext {
	var p = new(SelectorExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *SelectorExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorExpContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *SelectorExpContext) Selector() ISelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *SelectorExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSelectorExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type FuncCallContext struct {
	PrimaryExpressionContext
}

func NewFuncCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallContext {
	var p = new(FuncCallContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *FuncCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *FuncCallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FuncCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitFuncCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExpContext struct {
	PrimaryExpressionContext
}

func NewOpExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExpContext {
	var p = new(OpExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *OpExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExpContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *OpExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitOpExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExpContext struct {
	PrimaryExpressionContext
}

func NewIndexExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpContext {
	var p = new(IndexExpContext)

	InitEmptyPrimaryExpressionContext(&p.PrimaryExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PrimaryExpressionContext))

	return p
}

func (s *IndexExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *IndexExpContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *IndexExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIndexExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *expr_parser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, expr_parserRULE_primaryExpression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserLPAREN, expr_parserINT_LIT, expr_parserIDENTIFIER, expr_parserFLOAT_LIT, expr_parserRUNE_LIT, expr_parserRAW_STRING_LIT, expr_parserINTERPRETED_STRING_LIT:
		localctx = NewOpExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(272)
			p.Operand()
		}

	case expr_parserAPPEND:
		localctx = NewAppendExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(273)
			p.AppendExpression()
		}

	case expr_parserLENGHT:
		localctx = NewLenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(274)
			p.LengthExpression()
		}

	case expr_parserCAP:
		localctx = NewCapExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.CapExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSelectorExpContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expr_parserRULE_primaryExpression)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(279)
					p.Selector()
				}

			case 2:
				localctx = NewIndexExpContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expr_parserRULE_primaryExpression)
				p.SetState(280)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(281)
					p.Index()
				}

			case 3:
				localctx = NewFuncCallContext(p, NewPrimaryExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, expr_parserRULE_primaryExpression)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					p.Arguments()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(288)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Literal() ILiteralContext
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_operand
	return p
}

func InitEmptyOperandContext(p *OperandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_operand
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *OperandContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Operand() (localctx IOperandContext) {
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, expr_parserRULE_operand)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserINT_LIT, expr_parserFLOAT_LIT, expr_parserRUNE_LIT, expr_parserRAW_STRING_LIT, expr_parserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.Literal()
		}

	case expr_parserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(290)
			p.Match(expr_parserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserLPAREN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(291)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(292)
			p.expression(0)
		}
		{
			p.SetState(293)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyAll(ctx *LiteralContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FloatLitContext struct {
	LiteralContext
}

func NewFloatLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLitContext {
	var p = new(FloatLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserFLOAT_LIT, 0)
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type RuneLitContext struct {
	LiteralContext
}

func NewRuneLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuneLitContext {
	var p = new(RuneLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RuneLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuneLitContext) RUNE_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserRUNE_LIT, 0)
}

func (s *RuneLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitRuneLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntLitContext struct {
	LiteralContext
}

func NewIntLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntLitContext {
	var p = new(IntLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserINT_LIT, 0)
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type InterpretedStringLitContext struct {
	LiteralContext
}

func NewInterpretedStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterpretedStringLitContext {
	var p = new(InterpretedStringLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *InterpretedStringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterpretedStringLitContext) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserINTERPRETED_STRING_LIT, 0)
}

func (s *InterpretedStringLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitInterpretedStringLit(s)

	default:
		return t.VisitChildren(s)
	}
}

type RawStringLitContext struct {
	LiteralContext
}

func NewRawStringLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RawStringLitContext {
	var p = new(RawStringLitContext)

	InitEmptyLiteralContext(&p.LiteralContext)
	p.parser = parser
	p.CopyAll(ctx.(*LiteralContext))

	return p
}

func (s *RawStringLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawStringLitContext) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(expr_parserRAW_STRING_LIT, 0)
}

func (s *RawStringLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitRawStringLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, expr_parserRULE_literal)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserINT_LIT:
		localctx = NewIntLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(expr_parserINT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserFLOAT_LIT:
		localctx = NewFloatLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Match(expr_parserFLOAT_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserRAW_STRING_LIT:
		localctx = NewRawStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.Match(expr_parserRAW_STRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserINTERPRETED_STRING_LIT:
		localctx = NewInterpretedStringLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(300)
			p.Match(expr_parserINTERPRETED_STRING_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserRUNE_LIT:
		localctx = NewRuneLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(301)
			p.Match(expr_parserRUNE_LIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACK() antlr.TerminalNode
	Expression() IExpressionContext
	RBRACK() antlr.TerminalNode

	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACK, 0)
}

func (s *IndexContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACK, 0)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IndexContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIndex(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, expr_parserRULE_index)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(304)
		p.Match(expr_parserLBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(305)
		p.expression(0)
	}
	{
		p.SetState(306)
		p.Match(expr_parserRBRACK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpressionList() []IExpressionListContext
	ExpressionList(i int) IExpressionListContext

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *ArgumentsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *ArgumentsContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, expr_parserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(expr_parserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256210778653120) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(309)
			p.ExpressionList()
		}

		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(315)
		p.Match(expr_parserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_selector
	return p
}

func InitEmptySelectorContext(p *SelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_selector
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) DOT() antlr.TerminalNode {
	return s.GetToken(expr_parserDOT, 0)
}

func (s *SelectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(expr_parserIDENTIFIER, 0)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, expr_parserRULE_selector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(expr_parserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(expr_parserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAppendExpressionContext is an interface to support dynamic dispatch.
type IAppendExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	APPEND() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	COMMA() antlr.TerminalNode
	RPAREN() antlr.TerminalNode

	// IsAppendExpressionContext differentiates from other interfaces.
	IsAppendExpressionContext()
}

type AppendExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAppendExpressionContext() *AppendExpressionContext {
	var p = new(AppendExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_appendExpression
	return p
}

func InitEmptyAppendExpressionContext(p *AppendExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_appendExpression
}

func (*AppendExpressionContext) IsAppendExpressionContext() {}

func NewAppendExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AppendExpressionContext {
	var p = new(AppendExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_appendExpression

	return p
}

func (s *AppendExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AppendExpressionContext) APPEND() antlr.TerminalNode {
	return s.GetToken(expr_parserAPPEND, 0)
}

func (s *AppendExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *AppendExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AppendExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AppendExpressionContext) COMMA() antlr.TerminalNode {
	return s.GetToken(expr_parserCOMMA, 0)
}

func (s *AppendExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *AppendExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AppendExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AppendExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitAppendExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) AppendExpression() (localctx IAppendExpressionContext) {
	localctx = NewAppendExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, expr_parserRULE_appendExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(expr_parserAPPEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(expr_parserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.expression(0)
	}
	{
		p.SetState(323)
		p.Match(expr_parserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.expression(0)
	}
	{
		p.SetState(325)
		p.Match(expr_parserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILengthExpressionContext is an interface to support dynamic dispatch.
type ILengthExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LENGHT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsLengthExpressionContext differentiates from other interfaces.
	IsLengthExpressionContext()
}

type LengthExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLengthExpressionContext() *LengthExpressionContext {
	var p = new(LengthExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_lengthExpression
	return p
}

func InitEmptyLengthExpressionContext(p *LengthExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_lengthExpression
}

func (*LengthExpressionContext) IsLengthExpressionContext() {}

func NewLengthExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LengthExpressionContext {
	var p = new(LengthExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_lengthExpression

	return p
}

func (s *LengthExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *LengthExpressionContext) LENGHT() antlr.TerminalNode {
	return s.GetToken(expr_parserLENGHT, 0)
}

func (s *LengthExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *LengthExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LengthExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *LengthExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LengthExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitLengthExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) LengthExpression() (localctx ILengthExpressionContext) {
	localctx = NewLengthExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, expr_parserRULE_lengthExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(expr_parserLENGHT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Match(expr_parserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(329)
		p.expression(0)
	}
	{
		p.SetState(330)
		p.Match(expr_parserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICapExpressionContext is an interface to support dynamic dispatch.
type ICapExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CAP() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsCapExpressionContext differentiates from other interfaces.
	IsCapExpressionContext()
}

type CapExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapExpressionContext() *CapExpressionContext {
	var p = new(CapExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_capExpression
	return p
}

func InitEmptyCapExpressionContext(p *CapExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_capExpression
}

func (*CapExpressionContext) IsCapExpressionContext() {}

func NewCapExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapExpressionContext {
	var p = new(CapExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_capExpression

	return p
}

func (s *CapExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *CapExpressionContext) CAP() antlr.TerminalNode {
	return s.GetToken(expr_parserCAP, 0)
}

func (s *CapExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *CapExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CapExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *CapExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitCapExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) CapExpression() (localctx ICapExpressionContext) {
	localctx = NewCapExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, expr_parserRULE_capExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(expr_parserCAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(333)
		p.Match(expr_parserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(334)
		p.expression(0)
	}
	{
		p.SetState(335)
		p.Match(expr_parserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_statementList
	return p
}

func InitEmptyStatementListContext(p *StatementListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_statementList
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, expr_parserRULE_statementList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256211315752908) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
		{
			p.SetState(337)
			p.Statement()
		}

		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	StatementList() IStatementListContext
	RBRACE() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, expr_parserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(343)
		p.Match(expr_parserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(344)
		p.StatementList()
	}
	{
		p.SetState(345)
		p.Match(expr_parserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContinueStContext struct {
	StatementContext
}

func NewContinueStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStContext {
	var p = new(ContinueStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ContinueStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(expr_parserCONTINUE, 0)
}

func (s *ContinueStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *ContinueStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitContinueSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStatContext struct {
	StatementContext
}

func NewIfStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatContext {
	var p = new(IfStatContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *IfStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpleStContext struct {
	StatementContext
}

func NewSimpleStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleStContext {
	var p = new(SimpleStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SimpleStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpleStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *SimpleStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSimpleSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStContext struct {
	StatementContext
}

func NewReturnStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStContext {
	var p = new(ReturnStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ReturnStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStContext) RETURN() antlr.TerminalNode {
	return s.GetToken(expr_parserRETURN, 0)
}

func (s *ReturnStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *ReturnStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitReturnSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintStContext struct {
	StatementContext
}

func NewPrintStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintStContext {
	var p = new(PrintStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStContext) PRINT() antlr.TerminalNode {
	return s.GetToken(expr_parserPRINT, 0)
}

func (s *PrintStContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *PrintStContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *PrintStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *PrintStContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitPrintSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type TypeDeclStContext struct {
	StatementContext
}

func NewTypeDeclStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeDeclStContext {
	var p = new(TypeDeclStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *TypeDeclStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclStContext) TypeDecl() ITypeDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TypeDeclStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *TypeDeclStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitTypeDeclSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BlockStContext struct {
	StatementContext
}

func NewBlockStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockStContext {
	var p = new(BlockStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BlockStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *BlockStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitBlockSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakStContext struct {
	StatementContext
}

func NewBreakStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStContext {
	var p = new(BreakStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BreakStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStContext) BREAK() antlr.TerminalNode {
	return s.GetToken(expr_parserBREAK, 0)
}

func (s *BreakStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *BreakStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitBreakSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type VarDeclStContext struct {
	StatementContext
}

func NewVarDeclStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDeclStContext {
	var p = new(VarDeclStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *VarDeclStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclStContext) VariableDecl() IVariableDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclContext)
}

func (s *VarDeclStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitVarDeclSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type LoopStContext struct {
	StatementContext
}

func NewLoopStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopStContext {
	var p = new(LoopStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *LoopStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopStContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *LoopStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *LoopStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitLoopSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrintlnStContext struct {
	StatementContext
}

func NewPrintlnStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintlnStContext {
	var p = new(PrintlnStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *PrintlnStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnStContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(expr_parserPRINTLN, 0)
}

func (s *PrintlnStContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserLPAREN, 0)
}

func (s *PrintlnStContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(expr_parserRPAREN, 0)
}

func (s *PrintlnStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *PrintlnStContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *PrintlnStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitPrintlnSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStContext struct {
	StatementContext
}

func NewSwitchStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStContext {
	var p = new(SwitchStContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *SwitchStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStContext) Switch_() ISwitchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchContext)
}

func (s *SwitchStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *SwitchStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSwitchSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, expr_parserRULE_statement)
	var _la int

	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserPRINT:
		localctx = NewPrintStContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(347)
			p.Match(expr_parserPRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(350)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256210778653120) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(349)
				p.ExpressionList()
			}

		}
		{
			p.SetState(352)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserPRINTLN:
		localctx = NewPrintlnStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(expr_parserPRINTLN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			p.Match(expr_parserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256210778653120) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(356)
				p.ExpressionList()
			}

		}
		{
			p.SetState(359)
			p.Match(expr_parserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserRETURN:
		localctx = NewReturnStContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(361)
			p.Match(expr_parserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2256210778653120) != 0) || ((int64((_la-65)) & ^0x3f) == 0 && ((int64(1)<<(_la-65))&63) != 0) {
			{
				p.SetState(362)
				p.expression(0)
			}

		}
		{
			p.SetState(365)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserBREAK:
		localctx = NewBreakStContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(366)
			p.Match(expr_parserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserCONTINUE:
		localctx = NewContinueStContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(368)
			p.Match(expr_parserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(369)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserAPPEND, expr_parserLENGHT, expr_parserCAP, expr_parserLPAREN, expr_parserADD, expr_parserSUB, expr_parserCARET, expr_parserLOG_NOT, expr_parserINT_LIT, expr_parserIDENTIFIER, expr_parserFLOAT_LIT, expr_parserRUNE_LIT, expr_parserRAW_STRING_LIT, expr_parserINTERPRETED_STRING_LIT:
		localctx = NewSimpleStContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(370)
			p.SimpleStatement()
		}
		{
			p.SetState(371)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserLBRACE:
		localctx = NewBlockStContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(373)
			p.Block()
		}
		{
			p.SetState(374)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserSWITCH:
		localctx = NewSwitchStContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(376)
			p.Switch_()
		}
		{
			p.SetState(377)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserIF:
		localctx = NewIfStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(379)
			p.IfStatement()
		}
		{
			p.SetState(380)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserFOR:
		localctx = NewLoopStContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(382)
			p.Loop()
		}
		{
			p.SetState(383)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserTYPE:
		localctx = NewTypeDeclStContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(385)
			p.TypeDecl()
		}
		{
			p.SetState(386)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case expr_parserVAR:
		localctx = NewVarDeclStContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(388)
			p.VariableDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStatementContext is an interface to support dynamic dispatch.
type ISimpleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimpleStatementContext differentiates from other interfaces.
	IsSimpleStatementContext()
}

type SimpleStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStatementContext() *SimpleStatementContext {
	var p = new(SimpleStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_simpleStatement
	return p
}

func InitEmptySimpleStatementContext(p *SimpleStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_simpleStatement
}

func (*SimpleStatementContext) IsSimpleStatementContext() {}

func NewSimpleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStatementContext {
	var p = new(SimpleStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_simpleStatement

	return p
}

func (s *SimpleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStatementContext) CopyAll(ctx *SimpleStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SimpleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignStContext struct {
	SimpleStatementContext
}

func NewAssignStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStContext {
	var p = new(AssignStContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *AssignStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStContext) AssignmentStatement() IAssignmentStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentStatementContext)
}

func (s *AssignStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitAssignSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ShortDecStContext struct {
	SimpleStatementContext
}

func NewShortDecStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShortDecStContext {
	var p = new(ShortDecStContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ShortDecStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShortDecStContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *ShortDecStContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ShortDecStContext) SHORT_DEC() antlr.TerminalNode {
	return s.GetToken(expr_parserSHORT_DEC, 0)
}

func (s *ShortDecStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitShortDecSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpStContext struct {
	SimpleStatementContext
}

func NewExpStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpStContext {
	var p = new(ExpStContext)

	InitEmptySimpleStatementContext(&p.SimpleStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*SimpleStatementContext))

	return p
}

func (s *ExpStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpStContext) INCREMENT() antlr.TerminalNode {
	return s.GetToken(expr_parserINCREMENT, 0)
}

func (s *ExpStContext) DECREMENT() antlr.TerminalNode {
	return s.GetToken(expr_parserDECREMENT, 0)
}

func (s *ExpStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) SimpleStatement() (localctx ISimpleStatementContext) {
	localctx = NewSimpleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, expr_parserRULE_simpleStatement)
	var _la int

	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpStContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(391)
			p.expression(0)
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == expr_parserINCREMENT || _la == expr_parserDECREMENT {
			{
				p.SetState(392)
				_la = p.GetTokenStream().LA(1)

				if !(_la == expr_parserINCREMENT || _la == expr_parserDECREMENT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case 2:
		localctx = NewAssignStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.AssignmentStatement()
		}

	case 3:
		localctx = NewShortDecStContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.ExpressionList()
		}
		{
			p.SetState(397)
			p.Match(expr_parserSHORT_DEC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.ExpressionList()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentStatementContext is an interface to support dynamic dispatch.
type IAssignmentStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssignmentStatementContext differentiates from other interfaces.
	IsAssignmentStatementContext()
}

type AssignmentStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentStatementContext() *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_assignmentStatement
	return p
}

func InitEmptyAssignmentStatementContext(p *AssignmentStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_assignmentStatement
}

func (*AssignmentStatementContext) IsAssignmentStatementContext() {}

func NewAssignmentStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_assignmentStatement

	return p
}

func (s *AssignmentStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentStatementContext) CopyAll(ctx *AssignmentStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignStatContext struct {
	AssignmentStatementContext
}

func NewAssignStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignStatContext {
	var p = new(AssignStatContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *AssignStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignStatContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *AssignStatContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AssignStatContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserASSIGN, 0)
}

func (s *AssignStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitAssignStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type OtAssignStContext struct {
	AssignmentStatementContext
}

func NewOtAssignStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OtAssignStContext {
	var p = new(OtAssignStContext)

	InitEmptyAssignmentStatementContext(&p.AssignmentStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssignmentStatementContext))

	return p
}

func (s *OtAssignStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtAssignStContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OtAssignStContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OtAssignStContext) PLUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserPLUS_ASSIGN, 0)
}

func (s *OtAssignStContext) AND_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserAND_ASSIGN, 0)
}

func (s *OtAssignStContext) MINUS_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserMINUS_ASSIGN, 0)
}

func (s *OtAssignStContext) OR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserOR_ASSIGN, 0)
}

func (s *OtAssignStContext) TIMES_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserTIMES_ASSIGN, 0)
}

func (s *OtAssignStContext) XOR_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserXOR_ASSIGN, 0)
}

func (s *OtAssignStContext) LEFT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserLEFT_SHIFT_ASSIGN, 0)
}

func (s *OtAssignStContext) RIGHT_SHIFT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserRIGHT_SHIFT_ASSIGN, 0)
}

func (s *OtAssignStContext) AND_NOT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserAND_NOT_ASSIGN, 0)
}

func (s *OtAssignStContext) MOD_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserMOD_ASSIGN, 0)
}

func (s *OtAssignStContext) DIVIDE_ASSIGN() antlr.TerminalNode {
	return s.GetToken(expr_parserDIVIDE_ASSIGN, 0)
}

func (s *OtAssignStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitOtAssignSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) AssignmentStatement() (localctx IAssignmentStatementContext) {
	localctx = NewAssignmentStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, expr_parserRULE_assignmentStatement)
	var _la int

	p.SetState(410)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAssignStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.ExpressionList()
		}
		{
			p.SetState(403)
			p.Match(expr_parserASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.ExpressionList()
		}

	case 2:
		localctx = NewOtAssignStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.expression(0)
		}
		{
			p.SetState(407)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-54)) & ^0x3f) == 0 && ((int64(1)<<(_la-54))&2047) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(408)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) CopyAll(ctx *IfStatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfSimpleStContext struct {
	IfStatementContext
}

func NewIfSimpleStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleStContext {
	var p = new(IfSimpleStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfSimpleStContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *IfSimpleStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfSimpleSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseIfStContext struct {
	IfStatementContext
}

func NewIfElseIfStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseIfStContext {
	var p = new(IfElseIfStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseIfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseIfStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfElseIfStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseIfStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseIfStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(expr_parserELSE, 0)
}

func (s *IfElseIfStContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfElseIfStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfElseIfSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseIfStContext struct {
	IfStatementContext
}

func NewIfSimpleElseIfStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseIfStContext {
	var p = new(IfSimpleElseIfStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseIfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseIfStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfSimpleElseIfStContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleElseIfStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *IfSimpleElseIfStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleElseIfStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleElseIfStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(expr_parserELSE, 0)
}

func (s *IfSimpleElseIfStContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfSimpleElseIfStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfSimpleElseIfSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfElseBlockStContext struct {
	IfStatementContext
}

func NewIfElseBlockStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfElseBlockStContext {
	var p = new(IfElseBlockStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfElseBlockStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseBlockStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfElseBlockStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfElseBlockStContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfElseBlockStContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfElseBlockStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(expr_parserELSE, 0)
}

func (s *IfElseBlockStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfElseBlockSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfSimpleElseBlockStContext struct {
	IfStatementContext
}

func NewIfSimpleElseBlockStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfSimpleElseBlockStContext {
	var p = new(IfSimpleElseBlockStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfSimpleElseBlockStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfSimpleElseBlockStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfSimpleElseBlockStContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *IfSimpleElseBlockStContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *IfSimpleElseBlockStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfSimpleElseBlockStContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfSimpleElseBlockStContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfSimpleElseBlockStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(expr_parserELSE, 0)
}

func (s *IfSimpleElseBlockStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfSimpleElseBlockSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfStContext struct {
	IfStatementContext
}

func NewIfStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStContext {
	var p = new(IfStContext)

	InitEmptyIfStatementContext(&p.IfStatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*IfStatementContext))

	return p
}

func (s *IfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStContext) IF() antlr.TerminalNode {
	return s.GetToken(expr_parserIF, 0)
}

func (s *IfStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitIfSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, expr_parserRULE_ifStatement)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIfStContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(413)
			p.expression(0)
		}
		{
			p.SetState(414)
			p.Block()
		}

	case 2:
		localctx = NewIfElseIfStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.expression(0)
		}
		{
			p.SetState(418)
			p.Block()
		}
		{
			p.SetState(419)
			p.Match(expr_parserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.IfStatement()
		}

	case 3:
		localctx = NewIfElseBlockStContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(422)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(423)
			p.expression(0)
		}
		{
			p.SetState(424)
			p.Block()
		}
		{
			p.SetState(425)
			p.Match(expr_parserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(426)
			p.Block()
		}

	case 4:
		localctx = NewIfSimpleStContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(428)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(429)
			p.SimpleStatement()
		}
		{
			p.SetState(430)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(431)
			p.expression(0)
		}
		{
			p.SetState(432)
			p.Block()
		}

	case 5:
		localctx = NewIfSimpleElseIfStContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(434)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(435)
			p.SimpleStatement()
		}
		{
			p.SetState(436)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(437)
			p.expression(0)
		}
		{
			p.SetState(438)
			p.Block()
		}
		{
			p.SetState(439)
			p.Match(expr_parserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(440)
			p.IfStatement()
		}

	case 6:
		localctx = NewIfSimpleElseBlockStContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(442)
			p.Match(expr_parserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(443)
			p.SimpleStatement()
		}
		{
			p.SetState(444)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(445)
			p.expression(0)
		}
		{
			p.SetState(446)
			p.Block()
		}
		{
			p.SetState(447)
			p.Match(expr_parserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(448)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) CopyAll(ctx *LoopContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForExprStContext struct {
	LoopContext
}

func NewForExprStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForExprStContext {
	var p = new(ForExprStContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForExprStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForExprStContext) FOR() antlr.TerminalNode {
	return s.GetToken(expr_parserFOR, 0)
}

func (s *ForExprStContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForExprStContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *ForExprStContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(expr_parserSEMICOLON)
}

func (s *ForExprStContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, i)
}

func (s *ForExprStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForExprStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForExprStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitForExprSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForStContext struct {
	LoopContext
}

func NewForStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStContext {
	var p = new(ForStContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *ForStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStContext) FOR() antlr.TerminalNode {
	return s.GetToken(expr_parserFOR, 0)
}

func (s *ForStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitForSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type OtherForStContext struct {
	LoopContext
}

func NewOtherForStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OtherForStContext {
	var p = new(OtherForStContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *OtherForStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherForStContext) FOR() antlr.TerminalNode {
	return s.GetToken(expr_parserFOR, 0)
}

func (s *OtherForStContext) AllSimpleStatement() []ISimpleStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStatementContext); ok {
			tst[i] = t.(ISimpleStatementContext)
			i++
		}
	}

	return tst
}

func (s *OtherForStContext) SimpleStatement(i int) ISimpleStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *OtherForStContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(expr_parserSEMICOLON)
}

func (s *OtherForStContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, i)
}

func (s *OtherForStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *OtherForStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitOtherForSt(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileExprStContext struct {
	LoopContext
}

func NewWhileExprStContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileExprStContext {
	var p = new(WhileExprStContext)

	InitEmptyLoopContext(&p.LoopContext)
	p.parser = parser
	p.CopyAll(ctx.(*LoopContext))

	return p
}

func (s *WhileExprStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprStContext) FOR() antlr.TerminalNode {
	return s.GetToken(expr_parserFOR, 0)
}

func (s *WhileExprStContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileExprStContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileExprStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitWhileExprSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, expr_parserRULE_loop)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForStContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Match(expr_parserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(453)
			p.Block()
		}

	case 2:
		localctx = NewWhileExprStContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(454)
			p.Match(expr_parserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(455)
			p.expression(0)
		}
		{
			p.SetState(456)
			p.Block()
		}

	case 3:
		localctx = NewForExprStContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(458)
			p.Match(expr_parserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(459)
			p.SimpleStatement()
		}
		{
			p.SetState(460)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(461)
			p.expression(0)
		}
		{
			p.SetState(462)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(463)
			p.SimpleStatement()
		}
		{
			p.SetState(464)
			p.Block()
		}

	case 4:
		localctx = NewOtherForStContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(466)
			p.Match(expr_parserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(467)
			p.SimpleStatement()
		}
		{
			p.SetState(468)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(469)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.SimpleStatement()
		}
		{
			p.SetState(471)
			p.Block()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchContext is an interface to support dynamic dispatch.
type ISwitchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitchContext differentiates from other interfaces.
	IsSwitchContext()
}

type SwitchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchContext() *SwitchContext {
	var p = new(SwitchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_switch
	return p
}

func InitEmptySwitchContext(p *SwitchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_switch
}

func (*SwitchContext) IsSwitchContext() {}

func NewSwitchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchContext {
	var p = new(SwitchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_switch

	return p
}

func (s *SwitchContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchContext) CopyAll(ctx *SwitchContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpSwitchContext struct {
	SwitchContext
}

func NewExpSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpSwitchContext {
	var p = new(ExpSwitchContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *ExpSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpSwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(expr_parserSWITCH, 0)
}

func (s *ExpSwitchContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpSwitchContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *ExpSwitchContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpSwitchContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *ExpSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpStSwitchNoExpContext struct {
	SwitchContext
}

func NewSimpStSwitchNoExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpStSwitchNoExpContext {
	var p = new(SimpStSwitchNoExpContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SimpStSwitchNoExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpStSwitchNoExpContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(expr_parserSWITCH, 0)
}

func (s *SimpStSwitchNoExpContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpStSwitchNoExpContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *SimpStSwitchNoExpContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *SimpStSwitchNoExpContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SimpStSwitchNoExpContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *SimpStSwitchNoExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSimpStSwitchNoExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type SimpStSwitchContext struct {
	SwitchContext
}

func NewSimpStSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpStSwitchContext {
	var p = new(SimpStSwitchContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SimpStSwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpStSwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(expr_parserSWITCH, 0)
}

func (s *SimpStSwitchContext) SimpleStatement() ISimpleStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStatementContext)
}

func (s *SimpStSwitchContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(expr_parserSEMICOLON, 0)
}

func (s *SimpStSwitchContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpStSwitchContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *SimpStSwitchContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SimpStSwitchContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *SimpStSwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSimpStSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchDefaultContext struct {
	SwitchContext
}

func NewSwitchDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchDefaultContext {
	var p = new(SwitchDefaultContext)

	InitEmptySwitchContext(&p.SwitchContext)
	p.parser = parser
	p.CopyAll(ctx.(*SwitchContext))

	return p
}

func (s *SwitchDefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchDefaultContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(expr_parserSWITCH, 0)
}

func (s *SwitchDefaultContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserLBRACE, 0)
}

func (s *SwitchDefaultContext) ExpressionCaseClauseList() IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *SwitchDefaultContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(expr_parserRBRACE, 0)
}

func (s *SwitchDefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitSwitchDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) Switch_() (localctx ISwitchContext) {
	localctx = NewSwitchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, expr_parserRULE_switch)
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpStSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.Match(expr_parserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(476)
			p.SimpleStatement()
		}
		{
			p.SetState(477)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(478)
			p.expression(0)
		}
		{
			p.SetState(479)
			p.Match(expr_parserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(480)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(481)
			p.Match(expr_parserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExpSwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(483)
			p.Match(expr_parserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.expression(0)
		}
		{
			p.SetState(485)
			p.Match(expr_parserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(487)
			p.Match(expr_parserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewSimpStSwitchNoExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(489)
			p.Match(expr_parserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.SimpleStatement()
		}
		{
			p.SetState(491)
			p.Match(expr_parserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(492)
			p.Match(expr_parserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(494)
			p.Match(expr_parserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewSwitchDefaultContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(496)
			p.Match(expr_parserSWITCH)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
			p.Match(expr_parserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(498)
			p.ExpressionCaseClauseList()
		}
		{
			p.SetState(499)
			p.Match(expr_parserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseListContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionCaseClause() []IExpressionCaseClauseContext
	ExpressionCaseClause(i int) IExpressionCaseClauseContext
	AllExpressionCaseClauseList() []IExpressionCaseClauseListContext
	ExpressionCaseClauseList(i int) IExpressionCaseClauseListContext

	// IsExpressionCaseClauseListContext differentiates from other interfaces.
	IsExpressionCaseClauseListContext()
}

type ExpressionCaseClauseListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseListContext() *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionCaseClauseList
	return p
}

func InitEmptyExpressionCaseClauseListContext(p *ExpressionCaseClauseListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionCaseClauseList
}

func (*ExpressionCaseClauseListContext) IsExpressionCaseClauseListContext() {}

func NewExpressionCaseClauseListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseListContext {
	var p = new(ExpressionCaseClauseListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_expressionCaseClauseList

	return p
}

func (s *ExpressionCaseClauseListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseListContext) AllExpressionCaseClause() []IExpressionCaseClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			len++
		}
	}

	tst := make([]IExpressionCaseClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionCaseClauseContext); ok {
			tst[i] = t.(IExpressionCaseClauseContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionCaseClauseListContext) ExpressionCaseClause(i int) IExpressionCaseClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseContext)
}

func (s *ExpressionCaseClauseListContext) AllExpressionCaseClauseList() []IExpressionCaseClauseListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionCaseClauseListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionCaseClauseListContext); ok {
			tst[i] = t.(IExpressionCaseClauseListContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionCaseClauseListContext) ExpressionCaseClauseList(i int) IExpressionCaseClauseListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionCaseClauseListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionCaseClauseListContext)
}

func (s *ExpressionCaseClauseListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpressionCaseClauseList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) ExpressionCaseClauseList() (localctx IExpressionCaseClauseListContext) {
	localctx = NewExpressionCaseClauseListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, expr_parserRULE_expressionCaseClauseList)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(508)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(503)
				p.ExpressionCaseClause()
			}
			{
				p.SetState(504)
				p.ExpressionCaseClauseList()
			}

		}
		p.SetState(510)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionCaseClauseContext is an interface to support dynamic dispatch.
type IExpressionCaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionSwitchCase() IExpressionSwitchCaseContext
	COLON() antlr.TerminalNode
	StatementList() IStatementListContext

	// IsExpressionCaseClauseContext differentiates from other interfaces.
	IsExpressionCaseClauseContext()
}

type ExpressionCaseClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionCaseClauseContext() *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionCaseClause
	return p
}

func InitEmptyExpressionCaseClauseContext(p *ExpressionCaseClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionCaseClause
}

func (*ExpressionCaseClauseContext) IsExpressionCaseClauseContext() {}

func NewExpressionCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionCaseClauseContext {
	var p = new(ExpressionCaseClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_expressionCaseClause

	return p
}

func (s *ExpressionCaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionCaseClauseContext) ExpressionSwitchCase() IExpressionSwitchCaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSwitchCaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSwitchCaseContext)
}

func (s *ExpressionCaseClauseContext) COLON() antlr.TerminalNode {
	return s.GetToken(expr_parserCOLON, 0)
}

func (s *ExpressionCaseClauseContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *ExpressionCaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionCaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionCaseClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitExpressionCaseClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) ExpressionCaseClause() (localctx IExpressionCaseClauseContext) {
	localctx = NewExpressionCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, expr_parserRULE_expressionCaseClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.ExpressionSwitchCase()
	}
	{
		p.SetState(512)
		p.Match(expr_parserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(513)
		p.StatementList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionSwitchCaseContext is an interface to support dynamic dispatch.
type IExpressionSwitchCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSwitchCaseContext differentiates from other interfaces.
	IsExpressionSwitchCaseContext()
}

type ExpressionSwitchCaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSwitchCaseContext() *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionSwitchCase
	return p
}

func InitEmptyExpressionSwitchCaseContext(p *ExpressionSwitchCaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = expr_parserRULE_expressionSwitchCase
}

func (*ExpressionSwitchCaseContext) IsExpressionSwitchCaseContext() {}

func NewExpressionSwitchCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSwitchCaseContext {
	var p = new(ExpressionSwitchCaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = expr_parserRULE_expressionSwitchCase

	return p
}

func (s *ExpressionSwitchCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSwitchCaseContext) CopyAll(ctx *ExpressionSwitchCaseContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionSwitchCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSwitchCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CaseExpContext struct {
	ExpressionSwitchCaseContext
}

func NewCaseExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseExpContext {
	var p = new(CaseExpContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *CaseExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseExpContext) CASE() antlr.TerminalNode {
	return s.GetToken(expr_parserCASE, 0)
}

func (s *CaseExpContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *CaseExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitCaseExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefaultExpContext struct {
	ExpressionSwitchCaseContext
}

func NewDefaultExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultExpContext {
	var p = new(DefaultExpContext)

	InitEmptyExpressionSwitchCaseContext(&p.ExpressionSwitchCaseContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionSwitchCaseContext))

	return p
}

func (s *DefaultExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultExpContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(expr_parserDEFAULT, 0)
}

func (s *DefaultExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case expr_parserVisitor:
		return t.VisitDefaultExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *expr_parser) ExpressionSwitchCase() (localctx IExpressionSwitchCaseContext) {
	localctx = NewExpressionSwitchCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, expr_parserRULE_expressionSwitchCase)
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case expr_parserCASE:
		localctx = NewCaseExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(515)
			p.Match(expr_parserCASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(516)
			p.ExpressionList()
		}

	case expr_parserDEFAULT:
		localctx = NewDefaultExpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(517)
			p.Match(expr_parserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *expr_parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 18:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 20:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *expr_parser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *expr_parser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
