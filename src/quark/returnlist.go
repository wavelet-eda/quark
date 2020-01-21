package quark

type (
	SingleReturn struct {
		ReturnType TypeExpr
	}

	NamedReturn struct {
		ReturnTypes []OneNamedReturn

		openParen ObjectPosition
		closeParen ObjectPosition
	}

	OneNamedReturn struct {
		ReturnName *RealName
		ReturnType TypeExpr
	}
)

func NewNamedReturn(returns []OneNamedReturn, start ObjectPosition, end ObjectPosition) *NamedReturn {
	return &NamedReturn{
		ReturnTypes: returns,
		openParen:   start,
		closeParen:  end,
	}
}

func (r *SingleReturn) Start() *ObjectPosition {
	return r.ReturnType.Start()
}

func (r *NamedReturn) Start() *ObjectPosition {
	return &r.openParen
}

func (o OneNamedReturn) Start() *ObjectPosition {
	return o.ReturnType.Start()
}

func (r *SingleReturn) End() *ObjectPosition {
	return r.ReturnType.End()
}

func (r *NamedReturn) End() *ObjectPosition {
	return &r.closeParen
}

func (o OneNamedReturn) End() *ObjectPosition {
	return o.ReturnName.End()
}



func (o OneNamedReturn) returnListNode() {}
func (r *SingleReturn) returnListNode() {}
func (r *NamedReturn) returnListNode() {}


func (r *SingleReturn) Accept(v Visitor) {
	if v.Visit(r) == nil {
		return
	}

	r.ReturnType.Accept(v)
}

func (r *NamedReturn) Accept(v Visitor) {
	if v.Visit(r) == nil {
		return
	}

	for _, item := range r.ReturnTypes {
		item.Accept(v)
	}
}

func (o OneNamedReturn) Accept(v Visitor) {
	if v.Visit(o) == nil {
		return
	}

	o.ReturnType.Accept(v)
	o.ReturnName.Accept(v)
}