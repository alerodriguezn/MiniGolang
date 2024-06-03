package main

import (
	"MiniGolang/llvm"
	"MiniGolang/parser"
	"MiniGolang/routes"
	"encoding/json"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"os"
	"os/exec"
)

func main() {

	routes.Run()
	//testParser()
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
	//c := checker.NewChecker(errorListener)
	ll := llvm.NewEncoder(errorListener)

	mod := ll.GetModule()

	ll.Visit(tree)

	fmt.Println(mod.String())
	f, err := os.Create("module.ll")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer f.Close()

	if _, err := mod.WriteTo(f); err != nil {
		fmt.Println("Error al escribir el módulo:", err)
		return
	}

	cmd := exec.Command("clang", "module.ll", "-o", "module.exe")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error al compilar el módulo:", err)
		return
	}

	fmt.Println("El archivo ejecutable .exe ha sido generado correctamente.")

	//execute the file

	cmd = exec.Command("./module.exe") // Cambia esto por el ejecutable que deseas ejecutar

	// Capturar la salida estándar y la salida de error
	output, err := cmd.CombinedOutput()

	// Crear la estructura para el JSON
	type Output struct {
		Stdout string `json:"stdout"`
	}

	// Crear una instancia de la estructura y asignar la salida del comando
	outputData := Output{
		Stdout: string(output),
	}

	// Crear o abrir el archivo JSON
	file, err := os.Create("salida.json")
	if err != nil {
		fmt.Println("Error creando el archivo JSON:", err)
		return
	}
	defer file.Close()

	// Codificar los datos en formato JSON y escribirlos en el archivo
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Para una mejor legibilidad del JSON
	err = encoder.Encode(outputData)
	if err != nil {
		fmt.Println("Error escribiendo en el archivo JSON:", err)
		return
	}

	//print the output

	fmt.Println("Salida guardada exitosamente en salida.json")

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
