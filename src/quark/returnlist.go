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

func (r *SingleReturn) End() *ObjectPosition {
	return r.ReturnType.End()
}

func (r *NamedReturn) End() *ObjectPosition {
	return &r.closeParen
}

func (r *SingleReturn) returnListNode() {}
func (r *NamedReturn) returnListNode() {}
