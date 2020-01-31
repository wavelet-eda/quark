package quark

//note array index and array slice can appear as assignments. See array.go
type (

	//The left hand side of simple variable assignments.
	ValueAssignment struct {
		Variable Name
		IsWildcard bool

		QuestionMark ObjectPosition
	}

	//The left hand side of variable definitions
	VariableDefinitionAssignable struct {
		IsMut bool
		IsVar bool
		VarType TypeExpr
		VarName *RealName

		KwVar ObjectPosition
		KwMut ObjectPosition
	}

	//The left hand side of tuple destructing assignments
	TupleDestructionAssignment struct {
		Assignables []Assignable
	}
)

func (a *ValueAssignment) Start() *ObjectPosition {
	if a.IsWildcard {
		return &a.QuestionMark
	} else {
		return a.Variable.Start()
	}
}

func (a *VariableDefinitionAssignable) Start() *ObjectPosition {
	if a.IsMut {
		return &a.KwMut
	} else if a.IsVar {
		return &a.KwVar
	} else {
		return a.VarType.Start()
	}
}

func (a *TupleDestructionAssignment) Start() *ObjectPosition {
	return a.Assignables[0].Start()
}


func (a *ValueAssignment) End() *ObjectPosition {
	if a.IsWildcard {
		return a.QuestionMark.Next()
	} else {
		return a.Variable.End()
	}
}

func (a *VariableDefinitionAssignable) End() *ObjectPosition {
	return a.VarName.End()
}

func (a *TupleDestructionAssignment) End() *ObjectPosition {
	return a.Assignables[len(a.Assignables) - 1].End()
}

func (a *ValueAssignment) assignableNode() {}
func (a *VariableDefinitionAssignable) assignableNode() {}
func (a *TupleDestructionAssignment) assignableNode() {}

//Accept impls
func (a *ValueAssignment) Accept(v Visitor) {
	if v.Visit(a) == nil {
		return
	}

	if a.Variable != nil {
		a.Variable.Accept(v)
	}
}

func (a *VariableDefinitionAssignable) Accept(v Visitor) {
	if v.Visit(a) == nil {
		return
	}

	if a.VarType != nil {
		a.VarType.Accept(v)
	}
	a.VarName.Accept(v)
}

func (a *TupleDestructionAssignment) Accept(v Visitor) {
	if v.Visit(a) == nil {
		return
	}

	for _, assignable := range a.Assignables {
		assignable.Accept(v)
	}
}
