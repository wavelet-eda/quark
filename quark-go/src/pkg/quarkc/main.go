//Main application package. The main function and application glue code
//belong in this package. Most other logic should be in a specific package.
package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wwerst/wavelet/wavelet-go/src/parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: quarkc <file>")
	}
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewQuarkLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewQuarkParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	parseTreeConverter := parser.NewParseTreeConverter()
	p.Quarkpackage()
	parseTreeConverter.VisitQuarkpackage(p.Quarkpackage().(*parser.QuarkpackageContext))
}
