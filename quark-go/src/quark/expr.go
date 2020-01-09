package quark

type LiteralExpr struct {
	Value Literal
}

func (e LiteralExpr) exprNode() {} //sentinel impl that ensures the types all work out. Wow it'd be cool if Go had proper traits

func (e LiteralExpr) Start() ObjectPosition {
	return e.Value.Start()
}

func (e LiteralExpr) End() ObjectPosition {
	return e.Value.End()
}
