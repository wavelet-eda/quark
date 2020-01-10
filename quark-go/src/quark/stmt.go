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

		AssignStmt

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
		semi ObjectPosition
	}

	//A future block.
	FutureStmt struct {
		Futures []ArgumentDef
		Body []Stmt
		FutureAssignments []ConstructorField

		kwFuture ObjectPosition
		semi ObjectPosition
	}

	//A function or module return statement.
	ReturnStmt struct {
		ReturnExpr Expr

		kwReturn ObjectPosition
		semi ObjectPosition
	}
)

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
	return &s.kwFuture
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
	return s.semi.Next()
}
func (s *FutureStmt) End() *ObjectPosition {
	return s.semi.Next()
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