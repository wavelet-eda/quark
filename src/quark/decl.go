package quark

type (
	//Declaration of a Struct symbol.
	StructDecl struct {
		Annotations []Annotation
		StructName *RealName

		Parameters []*ParameterDef

		Fields []*Field
		Functions []*FunctionDecl

		TraitImpls []Name

		KwStruct ObjectPosition
		CloseCurly ObjectPosition
	}

	TraitDecl struct {
		Annotations []Annotation
		TraitName *RealName

		Parameters []*ParameterDef
		Functions []*FunctionSignature

		KwTrait ObjectPosition
		CloseCurly ObjectPosition
	}

	//Declaration of a Function symbol.
	FunctionDecl struct {
		Signature *FunctionSignature

		Body []Stmt
		CloseCurly ObjectPosition
	}

	FunctionSignature struct {
		Annotations []Annotation

		SymbolName *RealName

		Parameters []*ParameterDef
		Arguments []*ArgumentDef
		Returns ReturnList

		KwFunction ObjectPosition
		CloseParen ObjectPosition
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

func (t *TraitDecl) Start() *ObjectPosition {
	if len(t.Annotations) > 0 {
		return t.Annotations[0].Start()
	} else {
		return &t.KwTrait
	}
}

func (s *FunctionSignature) Start() *ObjectPosition {
	if len(s.Annotations) > 0 {
		return s.Annotations[0].Start()
	} else {
		return &s.KwFunction
	}
}

func (s *FunctionDecl) Start() *ObjectPosition {
	return s.Signature.Start()
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

func (t *TraitDecl) End() *ObjectPosition {
	return &t.CloseCurly
}

func (s *FunctionSignature) End() *ObjectPosition {
	if s.Returns != nil {
		return s.Returns.End()
	} else {
		return &s.CloseParen
	}
}

func (s *FunctionDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *ModuleDecl) End() *ObjectPosition {
	return &s.CloseCurly
}

func (s *StructDecl) declNode() {}
func (t *TraitDecl) declNode() {}
func (s *FunctionSignature) declNode() {}
func (s *FunctionDecl) declNode() {}
func (s *ModuleDecl) declNode() {}

//Accept impls

func (s *StructDecl) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.StructName.Accept(v)
	visitParamList(s.Parameters, v)

	for _, trait := range s.TraitImpls {
		trait.Accept(v)
	}

	for _, field := range s.Fields {
		field.Accept(v)
	}
}

func (t *TraitDecl) Accept(v Visitor) {
	if v.Visit(t) == nil {
		return
	}
	t.TraitName.Accept(v)
	visitParamList(t.Parameters, v)

	for _, sig := range t.Functions {
		sig.Accept(v)
	}

}

func (s *FunctionSignature) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.SymbolName.Accept(v)
	visitParamList(s.Parameters, v)

	for _, arg := range s.Arguments {
		arg.Accept(v)
	}

	if s.Returns != nil {
		s.Returns.Accept(v)
	}

}

func (s *FunctionDecl) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}


	for _, stmt := range s.Body {
		stmt.Accept(v)
	}
}

func (s *ModuleDecl) Accept(v Visitor) {
	if v.Visit(s) == nil {
		return
	}

	s.SymbolName.Accept(v)
	visitParamList(s.Parameters, v)

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