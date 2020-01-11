//Home of the parser and lexer implementations.
package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/wwerst/wavelet/wavelet-go/src/quark"
)

//Visitor implementation which converts parse tree from Antlr into
//proper quark.AST.
type ParseTreeConverter struct {
	BaseQuarkParserVisitor
	file *quark.QuarkFile
}

func NewParseTreeConverter(file *quark.QuarkFile) *ParseTreeConverter {
	return &ParseTreeConverter{file: file}
}

func (ptc *ParseTreeConverter) VisitQuarkpackage(ctx *QuarkpackageContext) interface{} {
	var rawImportDecls = ctx.AllImportdecl()
	var rawDecls = ctx.AllDecl()

	var importDecls = make([]quark.ImportDecl, len(rawImportDecls))
	var decls = make([]quark.Decl, len(rawDecls))

	for index, importDecl := range rawImportDecls {
		importDecls[index] = ptc.visitImportDecl(importDecl)
	}

	for index, decl := range rawDecls {
		decls[index] = ptc.visitDecl(decl)
	}

	return quark.Package{
		Imports: importDecls,
		Symbols: decls,
	}
}

func (ptc *ParseTreeConverter) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(ptc)
}

func (ptc *ParseTreeConverter) visitImportDecl(rawImportDecl IImportdeclContext) quark.ImportDecl {
	return ptc.Visit(rawImportDecl).(quark.ImportDecl)
}

func (ptc *ParseTreeConverter) visitDecl(rawDecl IDeclContext) quark.Decl {
	return ptc.Visit(rawDecl).(quark.Decl)
}

func (ptc *ParseTreeConverter) VisitSingleImport(ctx *SingleImportContext) interface{} {
	println("visiting single import")
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	return &quark.SingleImport{
		GenericImport: quark.GenericImport{PackageName:quarkName},
	}
}

func (ptc *ParseTreeConverter) VisitWildcardImport(ctx *WildcardImportContext) interface{} {
	quarkName := ptc.Visit(ctx.Name()).(quark.Name)
	return &quark.WildcardImport{
		GenericImport: quark.GenericImport{PackageName:quarkName},
	}
}

func (ptc *ParseTreeConverter) VisitRealname(ctx *RealnameContext) interface{} { //quark.Name
	pos := quark.NewObjectPosition(ptc.file, ctx.REAL_NAME().GetSymbol().GetTokenIndex())
	return quark.NewRealName(ctx.REAL_NAME().GetText(), pos)
}

func (ptc *ParseTreeConverter) VisitQualifiedName(ctx *QualifiedNameContext) interface{} { //quark.QualifiedName
	text := ctx.AllREAL_NAME()
	nameParts := make([]*quark.RealName, len(text))
	for index, node := range text {
		pos := quark.NewObjectPosition(ptc.file, node.GetSymbol().GetTokenIndex())
		text := node.GetText()
		nameParts[index] = quark.NewRealName(text, pos)
	}
	return &quark.QualifiedName{
		Parts: nameParts,
	}
}

func (ptc *ParseTreeConverter) VisitLiteral(ctx *LiteralContext) interface{} {
	pos := quark.NewObjectPosition(ptc.file, ctx.GetStart().GetTokenIndex())
	text := ctx.INTEGRAL().GetText()
	return quark.NewLiteral(text, pos)
}