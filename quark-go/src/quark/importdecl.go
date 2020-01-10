package quark

type (
	genericImport struct {
		PackageName Name

		kwImport ObjectPosition
	}

	SingleImport struct {
		genericImport
	}

	WildcardImport struct {
		genericImport

		star ObjectPosition
	}
)

func (i *genericImport) Start() *ObjectPosition {
	return &i.kwImport
}

func (i *genericImport) End() *ObjectPosition {
	return i.PackageName.End()
}

func (w *WildcardImport) End() *ObjectPosition {
	return w.star.Next()
}

func (i *genericImport) importDeclNode() {}