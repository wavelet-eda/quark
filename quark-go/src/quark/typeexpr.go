package quark

type (
	CompleteType struct {
		X Name
	}

	ParameterizedType struct {
		BaseType TypeExpr
		Parameters []TypeParameter

		closeBrace ObjectPosition
	}

	TypeParameter struct {
		XExpr Expr
		XType TypeExpr
	}
)

func (t *CompleteType) Start() ObjectPosition {
	return t.X.Start()
}

func (t *ParameterizedType) Start() ObjectPosition {
	return t.BaseType.Start()
}

func (t* TypeParameter) Start() ObjectPosition {
	if t.XExpr != nil {
		return t.XExpr.Start()
	} else {
		return t.XType.Start()
	}
}


func (t *CompleteType) End() ObjectPosition {
	return t.X.End()
}

func (t *ParameterizedType) End() ObjectPosition {
	return t.closeBrace
}

func (t *TypeParameter) End() ObjectPosition {
	if t.XExpr != nil {
		return t.XExpr.End()
	} else {
		return t.XType.End()
	}
}


func (t *CompleteType) typeExprNode() {}
func (t *ParameterizedType) typeExprNode() {}
func (t *TypeParameter) typeExprNode() {}

