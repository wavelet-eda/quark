package quark

type (
	//CompleteType is a quark type with no type parameters. CompleteType
	//represents either a named type or a type parameter.
	CompleteType struct {
		X Name
	}

	//ParameterizedType is a quark type with type or value parameters.
	ParameterizedType struct {
		BaseType TypeExpr
		Parameters []TypeParameter

		closeBrace ObjectPosition
	}

	//TypeParameter is a single parameter of a quark type. It is either
	//an expression or a type. It is invalid for both the expression and
	//type fields of this struct to have a value or be nil
	TypeParameter struct {
		XExpr Expr
		XType TypeExpr
	}
)

func (t *CompleteType) Start() *ObjectPosition {
	return t.X.Start()
}

func (t *ParameterizedType) Start() *ObjectPosition {
	return t.BaseType.Start()
}

func (t* TypeParameter) Start() *ObjectPosition {
	if t.XExpr != nil {
		return t.XExpr.Start()
	} else {
		return t.XType.Start()
	}
}


func (t *CompleteType) End() *ObjectPosition {
	return t.X.End()
}

func (t *ParameterizedType) End() *ObjectPosition {
	return &t.closeBrace
}

func (t *TypeParameter) End() *ObjectPosition {
	if t.XExpr != nil {
		return t.XExpr.End()
	} else {
		return t.XType.End()
	}
}


func (t *CompleteType) typeExprNode() {}
func (t *ParameterizedType) typeExprNode() {}
func (t *TypeParameter) typeExprNode() {}

