package quark

type (
	//An assignment of an expression to an assignable left hand side.
	AssignStmt struct {
		AssignTo Assignable
		AssignmentType AssignmentOp
		TheExpr Expr

		semi ObjectPosition
	}

	//An assignment to a clocked register.
	RegAssignStmt struct {
		Clock ClockExpr
		Reset ClockExpr

		*AssignStmt

		kwReg ObjectPosition
	}

	//A declaration of a variable without an assignment.
	DeclarationStmt struct {
		Declaration Assignable
		semi ObjectPosition
	}

	//A branch used in a statement context.
	BranchStmt struct {
		TheBranch Branch
	}

	//A future block.
	FutureStmt struct {
		FutureType TypeExpr
		FutureName *RealName

		KwFuture ObjectPosition
		Semi ObjectPosition
	}

	//A function or module return statement.
	ReturnStmt struct {
		ReturnExpr Expr

		kwReturn ObjectPosition
		semi ObjectPosition
	}
)

func NewAssignStmt(assignTo Assignable, op AssignmentOp, expr Expr, semi ObjectPosition) *AssignStmt {
	return &AssignStmt{
		AssignTo:       assignTo,
		AssignmentType: op,
		TheExpr:        expr,
		semi:           semi,
	}
}

func NewRegAssignStmt(clock ClockExpr, reset ClockExpr, assignment *AssignStmt, regPos ObjectPosition) *RegAssignStmt {
	return &RegAssignStmt{
		Clock:      clock,
		Reset:      reset,
		AssignStmt: assignment,
		kwReg:      regPos,
	}
}

func NewDeclStmt(assignable Assignable, semiPos ObjectPosition) *DeclarationStmt {
	return &DeclarationStmt{
		Declaration: assignable,
		semi:        semiPos,
	}
}

func NewReturnStmt(expr Expr, kwReturnPos ObjectPosition, semiPos ObjectPosition) *ReturnStmt {
	return &ReturnStmt{
		ReturnExpr: expr,
		kwReturn:   kwReturnPos,
		semi:       semiPos,
	}
}


func (s *AssignStmt) Start() *ObjectPosition {
	return s.AssignTo.Start()
}
func (s *RegAssignStmt) Start() *ObjectPosition {
	return &s.kwReg
}
func (s *DeclarationStmt) Start() *ObjectPosition {
	return s.Declaration.Start()
}
func (s *BranchStmt) Start() *ObjectPosition {
	return s.TheBranch.Start()
}
func (s *FutureStmt) Start() *ObjectPosition {
	return &s.KwFuture
}
func (s *ReturnStmt) Start() *ObjectPosition {
	return &s.kwReturn
}


func (s *AssignStmt) End() *ObjectPosition {
	return s.semi.Next()
}
func (s *RegAssignStmt) End() *ObjectPosition {
	return s.semi.Next()
}
func (s *DeclarationStmt) End() *ObjectPosition {
	return s.semi.Next()
}
func (s *BranchStmt) End() *ObjectPosition {
	return s.TheBranch.End()
}
func (s *FutureStmt) End() *ObjectPosition {
	return s.Semi.Next()
}
func (s *ReturnStmt) End() *ObjectPosition {
	return s.semi.Next()
}


func (s *AssignStmt) stmtNode() {}
func (s *RegAssignStmt) stmtNode() {}
func (s *DeclarationStmt) stmtNode() {}
func (s *BranchStmt) stmtNode() {}
func (s *FutureStmt) stmtNode() {}
func (s *ReturnStmt) stmtNode() {}

func (s *AssignStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.AssignTo.Accept(v)
	s.TheExpr.Accept(v)
}

func (s *RegAssignStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.Clock.Accept(v)
	if s.Reset != nil {
		s.Reset.Accept(v)
	}

	s.AssignTo.Accept(v)
	s.TheExpr.Accept(v)
}

func (s *DeclarationStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.Declaration.Accept(v)
}

func (s *BranchStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.TheBranch.Accept(v)
}

func (s *FutureStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.FutureType.Accept(v)
	s.FutureName.Accept(v)
}

func (s *ReturnStmt) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.ReturnExpr.Accept(v)
}