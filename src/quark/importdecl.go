package quark

type (
	GenericImport struct {
		PackageName Name

		KwImport ObjectPosition
	}

	SingleImport struct {
		GenericImport
	}

	WildcardImport struct {
		GenericImport

		star ObjectPosition
	}

	MultiImport struct {
		GenericImport
		Symbols []*RealName
		CloseParen ObjectPosition
	}
)

func (i *GenericImport) Start() *ObjectPosition {
	return &i.KwImport
}

func (i *GenericImport) End() *ObjectPosition {
	return i.PackageName.End()
}

func (w *WildcardImport) End() *ObjectPosition {
	return w.star.Next()
}

func (m *MultiImport) End() *ObjectPosition {
	return m.CloseParen.Next()
}

func (i *GenericImport) importDeclNode() {}

func (i *SingleImport) Accept(v Visitor) {
	if v.Visit(i) == nil {
		return
	}

	i.PackageName.Accept(v)
}

func (w *WildcardImport) Accept(v Visitor) {
	if v.Visit(w) == nil {
		return
	}

	w.PackageName.Accept(v)
}

func (m *MultiImport) Accept(v Visitor) {
	if v.Visit(m) == nil {
		return
	}

	m.PackageName.Accept(v)

	for _, symbol := range m.Symbols {
		symbol.Accept(v)
	}
}