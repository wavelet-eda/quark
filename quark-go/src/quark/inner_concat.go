package quark

type AtomicInnerConcat struct {
	X Expr
}

func (a AtomicInnerConcat) innerConcatNode() {}

func (a AtomicInnerConcat) Start() *ObjectPosition {
	return a.X.Start()
}

func (a AtomicInnerConcat) End() *ObjectPosition {
	return a.X.End()
}


type ReplicateConcat struct {
	ReplicateAmount Expr
	ReplicateExpr Expr

	rCurly ObjectPosition
}

func (r ReplicateConcat) innerConcatNode() {}

func (r ReplicateConcat) Start() *ObjectPosition {
	return r.ReplicateAmount.Start()
}

func (r ReplicateConcat) End() *ObjectPosition {
	return r.rCurly.Next()
}