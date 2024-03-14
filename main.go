package main

import (
	"MiniGolang/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	// Test Lexer
	input, _ := antlr.NewFileStream( /*os.Args[1]*/ "test.txt")
	inst := parser.Newexpr_lexer(input)
	tokens := antlr.NewCommonTokenStream(inst, 0)

	tokens.Fill()

	allTokens := tokens.GetAllTokens()

	for _, token := range allTokens {

		if token.GetTokenType() == antlr.TokenEOF {
			println("EOF")
			continue

		}
		tokenLiteralName := inst.SymbolicNames[token.GetTokenType()]
		println(tokenLiteralName, " ", token.GetText())

	}

}
