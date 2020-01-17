package quark

type (
	//Declaration of a Struct symbol.
	StructDecl struct {
		Annotations []Annotation
		StructName *RealName

		Parameters []*ParameterDef

		Fields []*Field

		TraitImpls []Name

		KwStruct ObjectPosition
		CloseCurly ObjectPosition
	}

	//Declaration of an Interface symbol.
	InterfaceDecl struct {
		Annotations []Annotation
		StructName *RealName

		Parameters []*ParameterDef

		Fields []*InterfaceField

		TraitImpls []Name

		KwInterface ObjectPosition
		CloseCurly ObjectPosition
	}

	//Declaration of a Function symbol.
	FunctionDecl struct {
		Annotations []Annotation

		SymbolName *RealName

		Parameters []*ParameterDef
		Arguments []*ArgumentDef
		Returns ReturnList

		Body []Stmt

		KwFunction ObjectPosition
		CloseCurly ObjectPosition
	}

	//Declaration of a module symbol.
	ModuleDecl struct {
		Annotations []Annotation

		SymbolName *RealName

		Parameters []*ParameterDef
		Arguments []*ArgumentDef
		Returns ReturnList

		Body []Stmt

		KwModule ObjectPosition
		CloseCurly ObjectPosition
	}
)

func (s *StructDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.KwStruct
	}
}

func (i *InterfaceDecl) Start() *ObjectPosition {
	if len(i.Annotations) > 0 {
		return i.Annotations[0].Start()
	} else {
		return &i.KwInterface
	}
}

func (s *FunctionDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.KwFunction
	}
}

func (s *ModuleDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.KwModule
	}
}

func (s *StructDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (i *InterfaceDecl) End() *ObjectPosition {
	return &i.CloseCurly
}

func (s *FunctionDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *ModuleDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *StructDecl) declNode() {}
func (i *InterfaceDecl) declNode() {}
func (s *FunctionDecl) declNode() {}
func (s *ModuleDecl) declNode() {}