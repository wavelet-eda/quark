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


type FieldExpr struct {
	Selectable Expr
	FieldName RealName
}

func (e FieldExpr) exprNode() {}

func (e FieldExpr) Start() ObjectPosition {
	return e.Selectable.Start()
}

func (e FieldExpr) End() ObjectPosition {
	return e.FieldName.End()
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
	return e.closeParen.Next()
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
	return e.closeParen.Next()
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
	return e.closeCurly.Next()
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
	return e.closeParen.Next()
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
	return e.closeCurly.Next()
}


type UnOp struct {
	Expr Expr
	Op UnaryOp

	opPosition ObjectPosition
}

func (e UnOp) exprNode() {}

func (e UnOp) Start() ObjectPosition {
	return e.opPosition
}

func (e UnOp) End() ObjectPosition {
	return e.Expr.End()
}


type ConcatExpr struct {
	ConcatPieces []InnerConcat

	lCurly ObjectPosition
	rCurly ObjectPosition
}

func (e ConcatExpr) exprNode() {}

func (e ConcatExpr) Start() ObjectPosition {
	return e.lCurly
}

func (e ConcatExpr) End() ObjectPosition {
	return e.rCurly.Next()
}

type BinOp struct {
	Left Expr
	Right Expr
	Op BinaryOp

	opPosition ObjectPosition
}

func (e BinOp) exprNode() {}

func (e BinOp) Start() ObjectPosition {
	return e.Left.Start()
}

func (e BinOp) End() ObjectPosition {
	return e.Right.End()
}


type TernaryExpr struct {
	IfExpr Expr
	Cond Expr
	ElseExpr Expr
}

func (e TernaryExpr) exprNode() {}

func (e TernaryExpr) Start() ObjectPosition {
	return e.IfExpr.Start()
}

func (e TernaryExpr) End() ObjectPosition {
	return e.ElseExpr.End()
}


type BranchExpr struct {
	X Branch
}

func (e BranchExpr) exprNode() {}

func (e BranchExpr) Start() ObjectPosition {
	return e.X.Start()
}

func (e BranchExpr) End() ObjectPosition {
	return e.X.End()
}


type SliceExpr struct {
	X Expr

	Msb Expr
	Lsb Expr
	Step Expr

	closeBrace ObjectPosition
}

func (e SliceExpr) exprNode() {}

func (e SliceExpr) Start() ObjectPosition {
	return e.X.Start()
}

func (e SliceExpr) End() ObjectPosition {
	return e.closeBrace.Next()
}


type ArrayIndexExpr struct {
	X Expr
	Indices []Expr

	closeBrace ObjectPosition
}

func (e ArrayIndexExpr) exprNode() {}

func (e ArrayIndexExpr) Start() ObjectPosition {
	return e.X.Start()
}

func (e ArrayIndexExpr) End() ObjectPosition {
	return e.closeBrace.Next()
}


type ArrayLiteralExpr struct {
	Values []Expr

	openBrace ObjectPosition
	closeBrace ObjectPosition
}

func (e ArrayLiteralExpr) exprNode() {}

func (e ArrayLiteralExpr) Start() ObjectPosition {
	return e.openBrace
}

func (e ArrayLiteralExpr) End() ObjectPosition {
	return e.closeBrace.Next()
}


type ClockToExpr struct {
	Clock ClockExpr

	kwSignal ObjectPosition
	closeParen ObjectPosition
}

func (e ClockToExpr) exprNode() {}

func (e ClockToExpr) Start() ObjectPosition {
	return e.kwSignal
}

func (e ClockToExpr) End() ObjectPosition {
	return e.closeParen.Next()
}