package quark

type (
	GenericImport struct {
		PackageName Name

		kwImport ObjectPosition
	}

	SingleImport struct {
		GenericImport
	}

	WildcardImport struct {
		GenericImport

		star ObjectPosition
	}
)

func (i *GenericImport) Start() *ObjectPosition {
	return &i.kwImport
}

func (i *GenericImport) End() *ObjectPosition {
	return i.PackageName.End()
}

func (w *WildcardImport) End() *ObjectPosition {
	return w.star.Next()
}

func (i *GenericImport) importDeclNode() {}