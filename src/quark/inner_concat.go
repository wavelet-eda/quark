package quark

//Atomic expression inside a concat context. This is just a normal
//expression.
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

//Replication expression in a concat context. First expression is the
//replication amount and the second expression is the replication value.
type ReplicateConcat struct {
	ReplicateAmount Expr
	ReplicateExpr   Expr

	rCurly ObjectPosition
}

func (r ReplicateConcat) innerConcatNode() {}

func (r ReplicateConcat) Start() *ObjectPosition {
	return r.ReplicateAmount.Start()
}

func (r ReplicateConcat) End() *ObjectPosition {
	return r.rCurly.Next()
}


//Accept impls
func (a *AtomicInnerConcat) Accept(v Visitor) {
	if v.Visit(a) == nil {
		return
	}

	a.X.Accept(v)
}

func (r *ReplicateConcat) Accept(v Visitor) {
	if v.Visit(r) == nil {
		return
	}

	r.ReplicateAmount.Accept(v)
	r.ReplicateExpr.Accept(v)
}