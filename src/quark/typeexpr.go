 package quark

type (
	//CompleteType is a quark type with no type parameters. CompleteType
	//represents either a named type or a type parameter.
	CompleteType struct {
		X Name
	}

	//ParameterizedType is a quark type with type or value parameters.
	ParameterizedType struct {
		BaseType   TypeExpr
		Parameters []*ParamArgument

		closeBrace ObjectPosition
	}

	//ParamArgument is a single parameter of a quark type. It is either
	//an expression or a type. It is invalid for both the expression and
	//type fields of this struct to have a value or be nil
	ParamArgument struct {
		XExpr Expr
		XType TypeExpr
		ParamName *RealName
		KwType ObjectPosition
	}
)

func NewParameterizedType(baseType TypeExpr, params []*ParamArgument, closeBrace ObjectPosition) *ParameterizedType {
	return &ParameterizedType{
		BaseType:   baseType,
		Parameters: params,
		closeBrace: closeBrace,
	}
}

func (t *CompleteType) Start() *ObjectPosition {
	return t.X.Start()
}

func (t *ParameterizedType) Start() *ObjectPosition {
	return t.BaseType.Start()
}

func (t*ParamArgument) Start() *ObjectPosition {
	if t.ParamName != nil {
		return t.ParamName.Start()
	} else if t.XExpr != nil {
		return t.XExpr.Start()
	} else {
		return &t.KwType
	}
}


func (t *CompleteType) End() *ObjectPosition {
	return t.X.End()
}

func (t *ParameterizedType) End() *ObjectPosition {
	return &t.closeBrace
}

func (t *ParamArgument) End() *ObjectPosition {
	if t.XExpr != nil {
		return t.XExpr.End()
	} else {
		return t.XType.End()
	}
}


func (t *CompleteType) typeExprNode() {}
func (t *ParameterizedType) typeExprNode() {}
func (t *ParamArgument) typeExprNode()     {}

func (t *CompleteType) Accept(v Visitor) {
	if v.Visit(t) == nil {
		return
	}

	t.X.Accept(v)
}

func (t *ParameterizedType) Accept(v Visitor) {
	if v.Visit(t) == nil {
		return
	}

	t.BaseType.Accept(v)
	for _, param := range t.Parameters {
		param.Accept(v)
	}
}

func (t *ParamArgument) Accept(v Visitor) {
	if v.Visit(t) == nil {
		return
	}

	if t.XType != nil {
		t.XType.Accept(v)
	}
	if t.XExpr != nil {
		t.XExpr.Accept(v)
	}
}
