//Main application package. The main function and application glue code
//belong in this package. Most other logic should be in a specific package.
package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wwerst/wavelet/wavelet-go/src/parser"
	"github.com/wwerst/wavelet/wavelet-go/src/quark"
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

	parseTreeConverter := parser.NewParseTreeConverter(&quark.QuarkFile{&os.Args[1]})
	quarkPackage := parseTreeConverter.VisitQuarkpackage(p.Quarkpackage().(*parser.QuarkpackageContext)).(quark.Package)

	fmt.Printf("Wow, that file had %d imports!\n", len(quarkPackage.Imports))
	fmt.Printf("And it had %d decls\n", len(quarkPackage.Symbols))
}
