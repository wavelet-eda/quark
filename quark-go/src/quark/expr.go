package quark

//Expression of a literal value.
type LiteralExpr struct {
	Value *Literal
}

func (e LiteralExpr) exprNode() {} //sentinel impl that ensures the types all work out. Wow it'd be cool if Go had proper traits

func (e LiteralExpr) Start() *ObjectPosition {
	return e.Value.Start()
}

func (e LiteralExpr) End() *ObjectPosition {
	return e.Value.End()
}

//Expression of a variable name.
type VarExpr struct {
	VarName Name
}

func (e VarExpr) exprNode() {}

func (e VarExpr) Start() *ObjectPosition {
	return e.VarName.Start()
}

func (e VarExpr) End() *ObjectPosition {
	return e.VarName.End()
}

//Struct, enum, or interface field access expression.
type FieldExpr struct {
	Selectable Expr
	FieldName  *RealName
}

func (e FieldExpr) exprNode() {}

func (e FieldExpr) Start() *ObjectPosition {
	return e.Selectable.Start()
}

func (e FieldExpr) End() *ObjectPosition {
	return e.FieldName.End()
}

//Expression enclosed in parens.
type ParensExpr struct {
	SubExpr Expr

	openParen  ObjectPosition
	closeParen ObjectPosition
}

func NewParensExpr(expr Expr, openParen ObjectPosition, closeParen ObjectPosition) *ParensExpr {
	return &ParensExpr{
		SubExpr:    expr,
		openParen:  openParen,
		closeParen: closeParen,
	}
}

func (e ParensExpr) exprNode() {}

func (e ParensExpr) Start() *ObjectPosition {
	return &e.openParen
}

func (e ParensExpr) End() *ObjectPosition {
	return e.closeParen.Next()
}


//A tuple constructor.
type TupleExpr struct {
	Exprs      []Expr

	openParen  ObjectPosition
	closeParen ObjectPosition
}

func NewTupleExpr(exprs []Expr, openParen ObjectPosition, closeParen ObjectPosition) *TupleExpr {
	return &TupleExpr{
		Exprs:    exprs,
		openParen:  openParen,
		closeParen: closeParen,
	}
}

func (e TupleExpr) exprNode() {}

func (e TupleExpr) Start() *ObjectPosition {
	return &e.openParen
}

func (e TupleExpr) End() *ObjectPosition {
	return e.closeParen.Next()
}


//A struct constructor.
type ConstructorExpr struct {
	FieldAssignments []*CallArgument

	openCurly  ObjectPosition
	closeCurly ObjectPosition
}

func NewConstructorExpr(assignments []*CallArgument, openCurly ObjectPosition, closeCurly ObjectPosition) *ConstructorExpr {
	return &ConstructorExpr{
		FieldAssignments: assignments,
		openCurly:        openCurly,
		closeCurly:       closeCurly,
	}
}

func (e ConstructorExpr) exprNode() {}

func (e ConstructorExpr) Start() *ObjectPosition {
	return &e.openCurly
}

func (e ConstructorExpr) End() *ObjectPosition {
	return e.closeCurly.Next()
}



//A module call.
type NewModuleExpr struct {
	ModuleType TypeExpr
	Arguments  []*CallArgument

	newKw      ObjectPosition
	closeParen ObjectPosition
}

func NewNewModuleExpr(moduleType TypeExpr, args []*CallArgument, newKw ObjectPosition, closeParen ObjectPosition) *NewModuleExpr {
	return &NewModuleExpr{
		ModuleType: moduleType,
		Arguments:  args,
		newKw:      newKw,
		closeParen: closeParen,
	}
}

func (e NewModuleExpr) exprNode() {}

func (e NewModuleExpr) Start() *ObjectPosition {
	return &e.newKw
}

func (e NewModuleExpr) End() *ObjectPosition {
	return e.closeParen.Next()
}

//An expression where an interface is opened
type OpenExpr struct {
	InterfaceType TypeExpr
	Arguments []*CallArgument

	closeParen ObjectPosition
}

func NewOpenExpr(interfaceType TypeExpr, args []*CallArgument, closeParen ObjectPosition) *OpenExpr {
	return &OpenExpr{
		InterfaceType: interfaceType,
		Arguments:     args,
		closeParen:    closeParen,
	}
}

func (e *OpenExpr) exprNode() {}
func (e *OpenExpr) Start() *ObjectPosition {
	return e.InterfaceType.Start()
}
func (e *OpenExpr) End() *ObjectPosition {
	return &e.closeParen
}


//An expression where an interface is closed
type CloseExpr struct {
	InterfaceExpr Expr
	Arguments []*CallArgument

	closeParen ObjectPosition
}

func NewCloseExpr(expr Expr, args []*CallArgument, closeParens ObjectPosition) *CloseExpr {
	return &CloseExpr{
		InterfaceExpr: expr,
		Arguments: args,
		closeParen: closeParens,
	}
}
//A function invocation expression.
func (e *CloseExpr) exprNode() {}
func (e *CloseExpr) Start() *ObjectPosition {
	return e.InterfaceExpr.Start()
}
func (e *CloseExpr) End() *ObjectPosition {
	return &e.closeParen
}


type FunctionCall struct {
	FunctionExpr Expr
	Arguments []*CallArgument

	closeParen ObjectPosition
}

func NewFunctionCall(expr Expr, args []*CallArgument, closeParen ObjectPosition) *FunctionCall {
	return &FunctionCall{
		FunctionExpr: expr,
		Arguments:    args,
		closeParen:   closeParen,
	}
}

func (e *FunctionCall) exprNode() {}
func (e *FunctionCall) Start() *ObjectPosition {
	return e.FunctionExpr.Start()
}
func (e *FunctionCall) End() *ObjectPosition {
	return &e.closeParen
}


//An in line lambda.
type LambdaExpr struct {
	Arguments []*ArgumentDef

	Body      []Stmt
	FinalExpr Expr

	kwLambda   ObjectPosition
	closeCurly ObjectPosition
}

func NewLambdaExpr(arguments []*ArgumentDef, body []Stmt, finalExpr Expr, lambdaPos ObjectPosition, closeCurly ObjectPosition) *LambdaExpr {
	return &LambdaExpr{
		Arguments:  arguments,
		Body:       body,
		FinalExpr:  finalExpr,
		kwLambda:   lambdaPos,
		closeCurly: closeCurly,
	}
}

func (e LambdaExpr) exprNode() {}

func (e LambdaExpr) Start() *ObjectPosition {
	return &e.kwLambda
}

func (e LambdaExpr) End() *ObjectPosition {
	return e.closeCurly.Next()
}


//A unary operator usage.
type UnOp struct {
	Expr Expr
	Op   UnaryOp

	opPosition ObjectPosition
}

func NewUnOp(expr Expr, op UnaryOp, opPos ObjectPosition) *UnOp {
	return &UnOp{
		Expr:       expr,
		Op:         op,
		opPosition: opPos,
	}
}

func (e UnOp) exprNode() {}

func (e UnOp) Start() *ObjectPosition {
	return &e.opPosition
}

func (e UnOp) End() *ObjectPosition {
	return e.Expr.End()
}


//A concatenation or replication expression.
type ConcatExpr struct {
	ConcatPieces []InnerConcat

	lCurly ObjectPosition
	rCurly ObjectPosition
}

func NewConcatExpr(inner []InnerConcat, lCurly ObjectPosition, rCurly ObjectPosition) *ConcatExpr {
	return &ConcatExpr{
		ConcatPieces: inner,
		lCurly:       lCurly,
		rCurly:       rCurly,
	}
}

func (e ConcatExpr) exprNode() {}

func (e ConcatExpr) Start() *ObjectPosition {
	return &e.lCurly
}

func (e ConcatExpr) End() *ObjectPosition {
	return e.rCurly.Next()
}


//A binary operator usage.
type BinOp struct {
	Left  Expr
	Right Expr
	Op    BinaryOp

	opPosition ObjectPosition
}

func NewBinOp(left Expr, right Expr, op BinaryOp, opPosition ObjectPosition) *BinOp {
	return &BinOp{
		Left:       left,
		Right:      right,
		Op:         op,
		opPosition: opPosition,
	}
}

func (e BinOp) exprNode() {}

func (e BinOp) Start() *ObjectPosition {
	return e.Left.Start()
}

func (e BinOp) End() *ObjectPosition {
	return e.Right.End()
}


//A ternary if expression.
type TernaryExpr struct {
	IfExpr   Expr
	Cond     Expr
	ElseExpr Expr
}

func (e TernaryExpr) exprNode() {}

func (e TernaryExpr) Start() *ObjectPosition {
	return e.IfExpr.Start()
}

func (e TernaryExpr) End() *ObjectPosition {
	return e.ElseExpr.End()
}


//An expression usage of an if or match branch.
type BranchExpr struct {
	X Branch
}

func (e BranchExpr) exprNode() {}

func (e BranchExpr) Start() *ObjectPosition {
	return e.X.Start()
}

func (e BranchExpr) End() *ObjectPosition {
	return e.X.End()
}


//A literal array constructor usage.
type ArrayLiteralExpr struct {
	Values []Expr

	openBrace  ObjectPosition
	closeBrace ObjectPosition
}

func NewArrayLiteralExpr(values []Expr, openBrace ObjectPosition, closeBrace ObjectPosition) *ArrayLiteralExpr {
	return &ArrayLiteralExpr{
		Values:     values,
		openBrace:  openBrace,
		closeBrace: closeBrace,
	}
}

func (e ArrayLiteralExpr) exprNode() {}

func (e ArrayLiteralExpr) Start() *ObjectPosition {
	return &e.openBrace
}

func (e ArrayLiteralExpr) End() *ObjectPosition {
	return e.closeBrace.Next()
}


//A use of the clock to signal function which converts clock expressions
//to signal expressions.
type ClockToExpr struct {
	Clock ClockExpr

	kwSignal   ObjectPosition
	closeParen ObjectPosition
}

func NewClockToExpr(clk ClockExpr, kwSignal ObjectPosition, closeParen ObjectPosition) *ClockToExpr {
	return &ClockToExpr{
		Clock:      clk,
		kwSignal:   kwSignal,
		closeParen: closeParen,
	}
}

func (e ClockToExpr) exprNode() {}

func (e ClockToExpr) Start() *ObjectPosition {
	return &e.kwSignal
}

func (e ClockToExpr) End() *ObjectPosition {
	return e.closeParen.Next()
}