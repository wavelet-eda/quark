package quark

type (
	LiteralPatten struct {
		X *Literal
	}

	NamedWildcardPattern struct {
		X *RealName
	}

	WildcardPattern struct {
		QuestionMark ObjectPosition
	}

	BitVectorPattern struct {
		Text string
		Token ObjectPosition
	}

	TuplePattern struct {
		Fields []Pattern

		OpenParen ObjectPosition
		CloseParen ObjectPosition
	}

	ArrayPattern struct {
		InnerPatterns []*InnerArrayPattern

		OpenBrace ObjectPosition
		CloseBrace ObjectPosition
	}

	EnumPattern struct {
		TypeName Name
		ParamPatterns []Pattern
		ArgPatterns []Pattern

		CloseParen ObjectPosition //either the arg pattern or param pattern end paren
	}

	InnerArrayPattern struct {
		X Pattern
		ConsumesMultiple bool

		DoubleDot ObjectPosition
	}
)

func (lp *LiteralPatten) Start() *ObjectPosition {
	return lp.X.Start()
}

func (nwp *NamedWildcardPattern) Start() *ObjectPosition {
	return nwp.X.Start()
}

func (wp *WildcardPattern) Start() *ObjectPosition {
	return &wp.QuestionMark
}

func (bvp *BitVectorPattern) Start() *ObjectPosition {
	return &bvp.Token
}

func (tp *TuplePattern) Start() *ObjectPosition {
	return &tp.OpenParen
}

func (ap *ArrayPattern) Start() *ObjectPosition {
	return &ap.OpenBrace
}

func (ep *EnumPattern) Start() *ObjectPosition {
	return ep.TypeName.Start()
}

func (iap *InnerArrayPattern) Start() *ObjectPosition {
	return iap.X.Start()
}

func (lp *LiteralPatten) End() *ObjectPosition {
	return lp.X.End()
}

func (nwp *NamedWildcardPattern) End() *ObjectPosition {
	return nwp.X.End()
}

func (wp *WildcardPattern) End() *ObjectPosition {
	return wp.QuestionMark.Next()
}

func (bvp *BitVectorPattern) End() *ObjectPosition {
	return bvp.Token.Next()
}

func (tp *TuplePattern) End() *ObjectPosition {
	return tp.CloseParen.Next()
}

func (ap *ArrayPattern) End() *ObjectPosition {
	return ap.CloseBrace.Next()
}

func (ep *EnumPattern) End() *ObjectPosition {
	if len(ep.ArgPatterns) > 0 || len(ep.ParamPatterns) > 0 {
		return ep.CloseParen.Next()
	} else {
		return ep.TypeName.End()
	}
}

func (iap *InnerArrayPattern) End() *ObjectPosition {
	if iap.ConsumesMultiple {
		return iap.DoubleDot.Next()
	} else {
		return iap.X.End()
	}
}

func (lp *LiteralPatten) Accept(v Visitor) {
	if v.Visit(lp) == nil {
		return
	}

	lp.X.Accept(v)
}

func (nwp *NamedWildcardPattern) Accept(v Visitor) {
	if v.Visit(nwp) == nil {
		return
	}

	nwp.X.Accept(v)
}

func (wp *WildcardPattern) Accept(v Visitor) {
	v.Visit(wp)
}

func (bvp *BitVectorPattern) Accept(v Visitor) {
	v.Visit(bvp)
}

func (tp *TuplePattern) Accept(v Visitor) {
	if v.Visit(tp) == nil {
		return
	}

	for _, field := range tp.Fields {
		field.Accept(v)
	}
}

func (ap *ArrayPattern) Accept(v Visitor) {
	if v.Visit(ap) == nil {
		return
	}

	for _, inner := range ap.InnerPatterns {
		inner.Accept(v)
	}
}

func (ep *EnumPattern) Accept(v Visitor) {
	if v.Visit(ep) == nil {
		return
	}

	ep.TypeName.Accept(v)
	for _, param := range ep.ParamPatterns {
		param.Accept(v)
	}

	for _, arg := range ep.ArgPatterns {
		arg.Accept(v)
	}
}

func (iap *InnerArrayPattern) Accept(v Visitor) {
	if v.Visit(iap) == nil {
		return
	}

	iap.X.Accept(v)
}

func (lp *LiteralPatten) patternNode() {}
func (nwp *NamedWildcardPattern) patternNode() {}
func (wp *WildcardPattern) patternNode() {}
func (bvp *BitVectorPattern) patternNode() {}
func (tp *TuplePattern) patternNode() {}
func (ap *ArrayPattern) patternNode()       {}
func (ep *EnumPattern) patternNode()        {}
func (iap *InnerArrayPattern) patternNode() {}
