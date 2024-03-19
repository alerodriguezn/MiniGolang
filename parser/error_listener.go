package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

type SyntaxErrorInformation struct {
	Message string `json:"message"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []SyntaxErrorInformation
}

func (cel *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {

	cel.Errors = append(cel.Errors, SyntaxErrorInformation{
		Message: msg,
		Line:    line,
		Column:  column,
	})

}
