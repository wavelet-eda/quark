//Main application package. The main function and application glue code
//belong in this package. Most other logic should be in a specific package.
package main

import (
	"fmt"
	"github.com/wavelet-eda/quark/src/parser"
	"github.com/wavelet-eda/quark/src/quark"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage: quarkc <file>")
	}
	quarkParser := parser.NewFileParser(os.Args[1])
	var quarkPackage *quark.Package = quarkParser.GetPackageAST()


	fmt.Printf("Wow, that file had %d imports!\n", len(quarkPackage.Imports))
	fmt.Printf("And it had %d decls\n", len(quarkPackage.Symbols))
}
