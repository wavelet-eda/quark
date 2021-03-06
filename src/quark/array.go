package quark

type (
	//An array or bit slice. Can appear as expression or assignable.
	Slice struct {
		ArrayExpr Expr

		Msb  Expr
		Lsb  Expr
		Step Expr

		closeBrace ObjectPosition
	}

//An array index. Can appear as expression or assignable.
	ArrayIndex struct {
		ArrayExpr Expr
		Indices   []Expr

		closeBrace ObjectPosition
	}
)

func NewArrayIndex(array Expr, indices []Expr, rBrace ObjectPosition) *ArrayIndex {
	return &ArrayIndex{
		ArrayExpr: array,
		Indices:  indices,
		closeBrace:    rBrace,
	}
}

func NewArraySlice(array Expr, msb Expr, lsb Expr, step Expr, closeBrace ObjectPosition) *Slice {
	return &Slice{
		ArrayExpr:  array,
		Msb:        msb,
		Lsb:        lsb,
		Step:       step,
		closeBrace: closeBrace,
	}
}

func (e *Slice) exprNode() {}
func (e Slice) assignableNode() {}

func (e *Slice) Start() *ObjectPosition {
	return e.ArrayExpr.Start()
}

func (e *Slice) End() *ObjectPosition {
	return e.closeBrace.Next()
}



func (e *ArrayIndex) exprNode() {}
func (e *ArrayIndex) assignableNode() {}

func (e *ArrayIndex) Start() *ObjectPosition {
	return e.ArrayExpr.Start()
}

func (e *ArrayIndex) End() *ObjectPosition {
	return e.closeBrace.Next()
}

//Accept impls

func (e *Slice) Accept(v Visitor) {
	if v.Visit(e) == nil {
		return
	}

	e.ArrayExpr.Accept(v)

	if e.Msb != nil {
		e.Msb.Accept(v)
	}
	if e.Lsb != nil {
		e.Lsb.Accept(v)
	}
	if e.Step != nil {
		e.Step.Accept(v)
	}
}

func (e *ArrayIndex) Accept(v Visitor) {
	if v.Visit(e) == nil {
		return
	}

	e.ArrayExpr.Accept(v)

	for _, index := range e.Indices {
		index.Accept(v)
	}
}