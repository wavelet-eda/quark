package quark

type AST interface {
	Start() ObjectPosition
	End() ObjectPosition
}

type Package interface {
	AST
	programNode()
}

type Decl interface {
	AST
	declNode()
}

type ImportDecl interface {
	AST
	importDeclNode()
}

type Stmt interface {
	AST
	stmtDecl()
}

type Assignment interface {
	AST
	assignmentNode()
}

type Name interface {
	AST
	nameNode()
}

type ClockExpr interface {
	AST
	clockExprNode()
}

type Expr interface {
	AST
	exprNode()
}

type ArraySlice interface {
	AST
	arraySliceNode()
}

type TypeExpr interface {
	AST
	typeExprNode()
}

type Branch interface {
	AST
	branchNode()
}

type Pattern interface {
	AST
	patternNode()
}

type ParameterDef interface {
	AST
	parameterDefNode()
}

type ReturnList interface {
	AST
	returnList()
}

type ArgumentList interface {
	AST
	argumentListNode()
}

type StructDef interface {
	AST
	structDefNode()
}

type TraitImpl interface {
	AST
	traitImplNode()
}

type FuncDecl interface {
	AST
	funcDeclNode()
}

type ModuleDecl interface {
	AST
	moduleDeclNode()
}

type Annotation interface {
	AST
	annotationNode()
}

type Literal interface {
	AST
	literalNode()
}

type InnerConcat interface {
	AST
	innerConcatNode()
}



