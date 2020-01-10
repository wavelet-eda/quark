//Home of the parser and lexer implementations.
package parser

//Visitor implementation which converts parse tree from Antlr into
//proper quark.AST.
type ParseTreeConverter struct {
	BaseQuarkParserVisitor
}