package quark

type (
	//AST is the base AST node type for quark. All AST nodes implement this
	//interface.
	AST interface {
		//Start returns the position of the first character which belongs to
		//this AST node.
		Start() *ObjectPosition

		//End returns the position the first character after this AST node.
		End() *ObjectPosition
	}

	//Package is a single Quark file.
	Package struct {
		Imports []ImportDecl
		Symbols []Decl
	}

	//Decl nodes are symbol declarations.
	Decl interface {
		AST
		declNode()
	}

	//ImportDecl nodes are package import declarations.
	ImportDecl interface {
		AST
		importDeclNode()
	}

	//Stmt nodes are statements which implement some behavior.
	Stmt interface {
		AST
		stmtDecl()
	}

	//Assignment nodes are variable and expression assignments.
	Assignment interface {
		AST
		assignmentNode()
	}

	//Name nodes are symbol names (both qualified and unqualified).
	Name interface {
		AST
		nameNode()
	}

	//ClockExpr nodes are expressions using clock and reset signals.
	ClockExpr interface {
		AST
		clockExprNode()
	}

	//Expr nodes are expressions which return values
	Expr interface {
		AST
		exprNode()
	}

	//TypeExpr nodes are expressions which return types
	TypeExpr interface {
		AST
		typeExprNode()
	}

	//Branch nodes are expressions or statements with if or match branches
	Branch interface {
		AST
		branchNode()
	}

	//Pattern nodes are pattern symbols in match branches
	Pattern interface {
		AST
		patternNode()
	}

	//ParameterDef nodes are param definitions in function, module, or type
	//declarations.
	ParameterDef interface {
		AST
		parameterDefNode()
	}

	//ReturnList nodes are single or multi returns from functions and modules.
	ReturnList interface {
		AST
		returnList()
	}

	//ArgumentList nodes are signal argument lists for functions and modules.
	ArgumentList interface {
		AST
		argumentListNode()
	}

	//StructDef nodes are declarations which define struct types.
	StructDef interface {
		AST
		structDefNode()
	}

	//TraitImpl nodes are declarations which define trait implementations.
	TraitImpl interface {
		AST
		traitImplNode()
	}

	//FuncDecl nodes are declarations which define functions.
	FuncDecl interface {
		AST
		funcDeclNode()
	}

	//ModuleDecl nodes are declarations which define modules.
	ModuleDecl interface {
		AST
		moduleDeclNode()
	}

	//Annotation nodes are declarations which define annotations.
	Annotation interface {
		AST
		annotationNode()
	}

	//Literal nodes are literal values in quark code.
	Literal interface {
		AST
		literalNode()
	}

	//InnerConcat are concatenation and replication expressions in concatenation
	//contexts. Examples are bit concatenation and replication.
	InnerConcat interface {
		AST
		innerConcatNode()
	}
)

func (p *Package) Start() *ObjectPosition {
	if len(p.Imports) > 0 {
		return p.Imports[0].Start()
	} else if len(p.Symbols) > 0 {
		return p.Symbols[0].Start()
	} else {
		return nil
	}
}

func (p *Package) End() *ObjectPosition {
	if len(p.Symbols) > 0 {
		return p.Symbols[len(p.Symbols) - 1].Start()
	} else if len(p.Imports) > 0 {
		return p.Imports[len(p.Imports) - 1].Start()
	} else {
		return nil
	}
}