package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wwerst/wavelet/wavelet-go/src/quark"
)

type Parser struct {
	tokens antlr.TokenStream
	antlrParser *QuarkParser
	antlrLexer *QuarkLexer
}

func NewFileParser(fileName string) *Parser {
	input, _ := antlr.NewFileStream(fileName)
	return initAnltr(input)
}

func NewStringParser(data string) *Parser {
	input := antlr.NewInputStream(data)
	return initAnltr(input)
}

func initAnltr(input antlr.CharStream) *Parser {
	lexer := NewQuarkLexer(input)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewQuarkParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	return &Parser{
		tokens:      tokenStream,
		antlrParser: p,
		antlrLexer:  lexer,
	}
}

func (p *Parser) AddErrorListener(listener antlr.ErrorListener) {
	p.antlrParser.AddErrorListener(listener)
}

func (p *Parser) GetPackageAST() *quark.Package {
	parseTreeConverter := NewParseTreeConverter(&quark.QuarkFile{FileName: nil}) //TODO: fix
	thePackage := p.antlrParser.Quarkpackage().(*QuarkpackageContext)
	return parseTreeConverter.VisitQuarkpackage(thePackage).(*quark.Package)
}


