package quark

type (
	//The left hand side of array index assignments.
	ArrayIndexAssignment struct {
		ArrayExpr Expr
		Indicies []Expr

		rBrace ObjectPosition
	}

	//The left hand side of array slice assignments.
	ArraySliceAssignment struct {
		ArrayExpr Expr
		Msb Expr
		Lsb Expr
		Step Expr

		rBrace ObjectPosition
	}

	//The left hand side of simple variable assignments.
	ValueAssignment struct {
		Variable Name
	}

	//The left hand side of variable definitions
	VariableDefinitionAssignable struct {
		IsMut bool
		VarType TypeExpr
		VarName RealName

		kwMut ObjectPosition
	}

	//The left hand side of tuple destructing assignments
	TupleDestructionAssignment struct {
		Assignables []Assignable
	}
)

func (a *ArrayIndexAssignment) Start() *ObjectPosition {
	return a.ArrayExpr.Start()
}

func (a *ArraySliceAssignment) Start() *ObjectPosition {
	return a.ArrayExpr.Start()
}

func (a *ValueAssignment) Start() *ObjectPosition {
	return a.Variable.Start()
}

func (a *VariableDefinitionAssignable) Start() *ObjectPosition {
	if a.IsMut {
		return &a.kwMut
	} else {
		return a.VarType.Start()
	}
}

func (a *TupleDestructionAssignment) Start() *ObjectPosition {
	return a.Assignables[0].Start()
}


func (a *ArrayIndexAssignment) End() *ObjectPosition {
	return a.rBrace.Next()
}

func (a *ArraySliceAssignment) End() *ObjectPosition {
	return a.rBrace.Next()
}

func (a *ValueAssignment) End() *ObjectPosition {
	return a.Variable.End()
}

func (a *VariableDefinitionAssignable) End() *ObjectPosition {
	return a.VarName.End()
}

func (a *TupleDestructionAssignment) End() *ObjectPosition {
	return a.Assignables[len(a.Assignables) - 1].End()
}

func (a *ArrayIndexAssignment) assignableNode() {}
func (a *ArraySliceAssignment) assignableNode() {}
func (a *ValueAssignment) assignableNode() {}
func (a *VariableDefinitionAssignable) assignableNode() {}
func (a *TupleDestructionAssignment) assignableNode() {}