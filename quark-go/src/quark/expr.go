package quark

type LiteralExpr struct {
	Value Literal
}

func (e LiteralExpr) exprNode() {} //sentinel impl that ensures the types all work out. Wow it'd be cool if Go had proper traits

func (e LiteralExpr) Start() ObjectPosition {
	return e.Value.Start()
}

func (e LiteralExpr) End() ObjectPosition {
	return e.Value.End()
}


type VarExpr struct {
	VarName Name
}

func (e VarExpr) exprNode() {}

func (e VarExpr) Start() ObjectPosition {
	return e.VarName.Start()
}

func (e VarExpr) End() ObjectPosition {
	return e.VarName.End()
}


type ParensExpr struct {
	SubExpr    Expr

	openParen  ObjectPosition
	closeParen ObjectPosition
}

func (e ParensExpr) exprNode() {}

func (e ParensExpr) Start() ObjectPosition {
	return e.openParen
}

func (e ParensExpr) End() ObjectPosition {
	return e.closeParen
}


type TupleExpr struct {
	Exprs      []Expr

	openParen  ObjectPosition
	closeParen ObjectPosition
}

func (e TupleExpr) exprNode() {}

func (e TupleExpr) Start() ObjectPosition {
	return e.openParen
}

func (e TupleExpr) End() ObjectPosition {
	return e.closeParen
}


type ConstructorExpr struct {
	FieldAssignments []ConstructorField

	openCurly ObjectPosition
	closeCurly ObjectPosition
}

func (e ConstructorExpr) exprNode() {}

func (e ConstructorExpr) Start() ObjectPosition {
	return e.openCurly
}

func (e ConstructorExpr) End() ObjectPosition {
	return e.closeCurly
}


type NewModuleExpr struct {
	ModuleType TypeExpr
	Arguments []Expr

	newKw ObjectPosition
	closeParen ObjectPosition
}

func (e NewModuleExpr) exprNode() {}

func (e NewModuleExpr) Start() ObjectPosition {
	return e.newKw
}

func (e NewModuleExpr) End() ObjectPosition {
	return e.closeParen
}


type LambdaExpr struct {
	Arguments ArgumentList

	Body []Stmt
	FinalExpr Expr

	kwLambda ObjectPosition
	closeCurly ObjectPosition
}

func (e LambdaExpr) exprNode() {}

func (e LambdaExpr) Start() ObjectPosition {
	return e.kwLambda
}

func (e LambdaExpr) End() ObjectPosition {
	return e.closeCurly
}


type UnopExpr struct {
	Expr Expr
	Op UnaryOp

	opPosition ObjectPosition
}

func (e UnopExpr) exprNode() {}

func (e UnopExpr) Start() ObjectPosition {
	return e.opPosition
}

func (e UnopExpr) End() ObjectPosition {
	return e.Expr.End()
}