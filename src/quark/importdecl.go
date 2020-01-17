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