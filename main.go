package main

import (
	"MiniGolang/checker"
	"MiniGolang/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {

	//routes.Run()
	testParser()
	//testLexer()
}

func testParser() {
	input, _ := antlr.NewFileStream( /*os.Args[1]*/ "example1.minigo")
	lexer := parser.Newexpr_lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.Newexpr_parser(stream)
	errorListener := &parser.CustomErrorListener{}
	p.AddErrorListener(errorListener)

	tree := p.Root()
	c := checker.NewChecker(errorListener)
	c.Visit(tree)

	allError := c.ErrorListener.Errors
	if allError != nil {
		for _, err := range allError {
			println(err.Message)
		}
		return

	}

	//c.Table.Print()

	//print errorlist
	//for _, err := range c.ErrorList {
	//	println(err)
	//
	//}

}

func testLexer() {
	input, _ := antlr.NewFileStream( /*os.Args[1]*/ "example1.minigo")
	inst := parser.Newexpr_lexer(input)
	tokens := antlr.NewCommonTokenStream(inst, 0)

	tokens.Fill()
	//get Token recognition error

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
