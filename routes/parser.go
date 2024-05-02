package routes

import (
	"MiniGolang/checker"
	"MiniGolang/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/gin-gonic/gin"
)

type Code struct {
	Code string `json:"code"`
}

func parseCode(rg *gin.RouterGroup) {
	parse := rg.Group("/parser")

	parse.POST("/", func(c *gin.Context) {

		var code Code
		if err := c.ShouldBindJSON(&code); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		errors := parseString(code.Code)

		if errors != nil {
			totalErrors := len(errors)
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Compilation failed! %d Errors found. ⚠️", totalErrors),
				"errors":  errors,
			})
			return
		}

		c.JSON(200, gin.H{"message": "Compiled successfully! 0 errors found. ✅"})

	})

}

func parseString(code string) []parser.SyntaxErrorInformation {

	input := antlr.NewInputStream(code)
	lexer := parser.Newexpr_lexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)
	//get tokens
	stream.Fill()
	p := parser.Newexpr_parser(stream)
	errorListener := &parser.CustomErrorListener{}
	p.AddErrorListener(errorListener)
	// Add Type Checker
	tree := p.Root()
	c := checker.NewChecker(errorListener)
	c.Visit(tree)

	if errorListener.Errors == nil {
		return nil
	}

	return errorListener.Errors

}
