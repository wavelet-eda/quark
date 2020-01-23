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

func (s *FunctionDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *ModuleDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *StructDecl) declNode() {}
func (s *FunctionDecl) declNode() {}
func (s *ModuleDecl) declNode() {}

//Accept impls

func (s *StructDecl) Accept(v Visitor) {
	if v.Visit(s) != nil {
		return
	}

	s.StructName.Accept(v)
	for _, params := range s.Parameters {
		params.Accept(v)
	}

	for _, trait := range s.TraitImpls {
		trait.Accept(v)
	}

	for _, field := range s.Fields {
		field.Accept(v)
	}
}

func (s *FunctionDecl) Accept(v Visitor) {
	if v.Visit(s) != nil {
		return
	}

	s.SymbolName.Accept(v)
	for _, param := range s.Parameters {
		param.Accept(v)
	}

	for _, arg := range s.Arguments {
		arg.Accept(v)
	}

	if s.Returns != nil {
		s.Returns.Accept(v)
	}

	for _, stmt := range s.Body {
		stmt.Accept(v)
	}
}

func (s *ModuleDecl) Accept(v Visitor) {
	if v.Visit(s) != nil {
		return
	}

	s.SymbolName.Accept(v)
	for _, param := range s.Parameters {
		param.Accept(v)
	}

	for _, arg := range s.Arguments {
		arg.Accept(v)
	}

	if s.Returns != nil {
		s.Returns.Accept(v)
	}

	for _, stmt := range s.Body {
		stmt.Accept(v)
	}
}