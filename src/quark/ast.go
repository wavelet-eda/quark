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

		//Accepts a visitor and calls visit on its children iff
		//visitor.Visit(this) does not return nil.
		Accept(visitor Visitor)
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
		stmtNode()
	}

	//Assignable nodes appear on the left hand side of assignments.
	Assignable interface {
		AST
		assignableNode()
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
	//declarations. If IsType is true, ExprVal must be null.
	ParameterDef struct {
		IsType bool
		TypeVal TypeExpr
		ParamName *RealName

		KwType ObjectPosition
	}

	//ArgumentDef nodes are functions or module argument declarations.
	ArgumentDef struct {
		ArgType TypeExpr
		ArgName *RealName
		IsFuture bool

		KwFuture ObjectPosition
	}

	//ReturnList nodes are single or multi returns from functions and modules.
	ReturnList interface {
		AST
		returnListNode()
	}

	//Annotation nodes are declarations which define annotations.
	Annotation struct {
		AnnotationName *RealName

		Pos ObjectPosition
	}

	//Literal nodes are literal values in quark code.
	Literal struct {
		Text string

		token ObjectPosition
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

func (p *ParameterDef) Start() *ObjectPosition {
	if p.IsType {
		return &p.KwType
	} else {
		return p.TypeVal.Start()
	}
}

func (a *ArgumentDef) Start() *ObjectPosition {
	if a.IsFuture {
		return &a.KwFuture
	}
	return a.ArgType.Start()
}

func (a *Annotation) Start() *ObjectPosition {
	return a.AnnotationName.Start()
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

func (p *ParameterDef) End() *ObjectPosition {
	if p.IsType {
		return p.TypeVal.End()
	} else {
		return p.ParamName.End()
	}
}

func (a *ArgumentDef) End() *ObjectPosition {
	return a.ArgName.End()
}

func (a *Annotation) End() *ObjectPosition {
	return a.AnnotationName.End()
}

func (l *Literal) Start() *ObjectPosition {
	return &l.token
}

func (l *Literal) End() *ObjectPosition {
	return l.token.Next()
}

func NewLiteral(text string, pos ObjectPosition) *Literal {
	return &Literal{text, pos}
}

//Accept impls
func (p *Package) Accept(v Visitor) {
	if v.Visit(p) == nil {
		return
	}

	for _, importDecl := range p.Imports {
		importDecl.Accept(v)
	}

	for _, symbolDecl := range p.Symbols {
		symbolDecl.Accept(v)
	}
}

func (p *ParameterDef) Accept(v Visitor) {
	if v.Visit(p) == nil {
		return
	}

	if p.TypeVal != nil {
		p.TypeVal.Accept(v)
	}

	p.ParamName.Accept(v)
}

func (a *ArgumentDef) Accept(v Visitor) {
	if v.Visit(a) != nil {
		return
	}

	a.ArgType.Accept(v)
	a.ArgName.Accept(v)
}

func (a *Annotation) Accept(v Visitor) {
	if v.Visit(a) != nil {
		return
	}

	a.AnnotationName.Accept(v)
}

func (l *Literal) Accept(v Visitor) {
	v.Visit(l)
}