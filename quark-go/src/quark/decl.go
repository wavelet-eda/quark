package quark

type (
	//Declaration of a Struct symbol.
	StructDecl struct {
		Annotations []Annotation
		StructName RealName

		Parameters []ParameterDef

		Fields []Field

		TraitImpls []Name

		kwStruct ObjectPosition
		closeCurly ObjectPosition
	}

	//Declaration of a Function symbol.
	FunctionDecl struct {
		Annotations []Annotation

		SymbolName RealName

		Parameters []ParameterDef
		Arguments []ArgumentDef
		Returns ReturnList

		Body []Stmt

		kwFunction ObjectPosition
		closeCurly ObjectPosition
	}

	//Declaration of a module symbol.
	ModuleDecl struct {
		Annotations []Annotation

		SymbolName RealName

		Parameters []ParameterDef
		Arguments []ArgumentDef
		Returns ReturnList

		InnerStructs []StructDecl
		Body []Stmt

		kwModule ObjectPosition
		closeCurly ObjectPosition
	}
)

func (s *StructDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.kwStruct
	}
}

func (s *FunctionDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.kwFunction
	}
}

func (s *ModuleDecl) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.kwModule
	}
}

func (s *StructDecl) End() *ObjectPosition {
	return &s.closeCurly
}

func (s *FunctionDecl) End() *ObjectPosition {
	return &s.closeCurly
}

func (s *ModuleDecl) End() *ObjectPosition {
	return &s.closeCurly
}

func (s *StructDecl) declNode() {}
func (s *FunctionDecl) declNode() {}
func (s *ModuleDecl) declNode() {}