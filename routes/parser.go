package routes

import (
	"MiniGolang/llvm"
	"MiniGolang/parser"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
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

		errors, data := parseString(code.Code)

		if errors != nil {
			totalErrors := len(errors)
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Compilation failed! %d Errors found. ⚠️", totalErrors),
				"errors":  errors,
				"output":  data,
			})
			return
		}

		fmt.Println("Output2: ", data)

		c.JSON(200, gin.H{"message": "Compiled successfully! 0 errors found. ✅", "output": data})

	})

}

type Output struct {
	Stdout string `json:"stdout"`
}

func parseString(code string) ([]parser.SyntaxErrorInformation, string) {

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
	//c := checker.NewChecker(errorListener)
	ll := llvm.NewEncoder(errorListener)

	mod := ll.GetModule()

	ll.Visit(tree)

	fmt.Println(mod.String())
	f, err := os.Create("module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
	}
	defer f.Close()

	if _, err := mod.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
	}

	cmd := exec.Command("clang", "module.ll", "-o", "module.exe")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente .")

	//execute the file

	cmd = exec.Command("./module.exe") // Cambia esto por el ejecutable que deseas ejecutar

	// Capturar la salida estándar y la salida de error
	output, err := cmd.CombinedOutput()

	// Crear la estructura para el JSON

	// Crear una instancia de la estructura y asignar la salida del comando
	outputData := Output{
		Stdout: string(output),
	}

	// Crear o abrir el archivo JSON
	file, err := os.Create("salida.json")
	if err != nil {
		fmt.Println("Error creando el archivo JSON:", err)
	}
	defer file.Close()

	// Codificar los datos en formato JSON y escribirlos en el archivo
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Para una mejor legibilidad del JSON
	err = encoder.Encode(outputData)
	if err != nil {
		fmt.Println("Error escribiendo en el archivo JSON:", err)
	}

	//return errors and output with br

	//fmt.Println("Output: ", outputData.Stdout)

	return errorListener.Errors, outputData.Stdout

}
