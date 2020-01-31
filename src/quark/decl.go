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

	EnumDecl struct {
		Annotations []Annotation
		EnumName *RealName
		Parameters []*ParameterDef

		TraitImpls []Name

		Constructors []*EnumConstructor
		Functions []*FunctionDecl

		KwEnum ObjectPosition
		CloseCurly ObjectPosition
	}

	EnumConstructor struct {
		ConstructorName *RealName
		Arguments []*EnumConstructorArgument
		CloseParen ObjectPosition
	}

	EnumConstructorArgument struct {
		ArgType TypeExpr
		ArgName Name
		IsFuture bool

		KwFuture ObjectPosition
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

func (e *EnumDecl) Start() *ObjectPosition {
	if len(e.Annotations) > 0 {
		return e.Annotations[0].Start()
	} else {
		return &e.KwEnum
	}
}

func (ec *EnumConstructor) Start() *ObjectPosition {
	return ec.ConstructorName.Start()
}

func (arg *EnumConstructorArgument) Start() *ObjectPosition {
	if arg.IsFuture {
		return &arg.KwFuture
	} else {
		return arg.ArgType.Start()
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

func (e *EnumDecl) End() *ObjectPosition {
	return &e.CloseCurly
}

func (ec *EnumConstructor) End() *ObjectPosition {
	if len(ec.Arguments) > 0 {
		return &ec.CloseParen
	} else {
		return ec.ConstructorName.End()
	}
}

func (arg *EnumConstructorArgument) End() *ObjectPosition {
	if arg.ArgName != nil {
		return arg.ArgName.End()
	} else {
		return arg.ArgType.End()
	}
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
func (e *EnumDecl) declNode() {}
func (ec *EnumConstructor) declNode() {}
func (arg *EnumConstructorArgument) declNode() {}
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

func (e *EnumDecl) Accept(v Visitor) {
	if v.Visit(e) == nil {
		return
	}

	e.EnumName.Accept(v)
	visitParamList(e.Parameters, v)

	for _, name := range e.TraitImpls {
		name.Accept(v)
	}

	for _, constructor := range e.Constructors {
		constructor.Accept(v)
	}

	for _, function := range e.Functions {
		function.Accept(v)
	}
}

func (ec *EnumConstructor) Accept(v Visitor) {
	if v.Visit(ec) == nil {
		return
	}

	ec.ConstructorName.Accept(v)
	for _, arg := range ec.Arguments {
		arg.Accept(v)
	}
}

func (arg *EnumConstructorArgument) Accept(v Visitor) {
	if v.Visit(arg) == nil {
		return
	}

	arg.ArgType.Accept(v)
	if arg.ArgName != nil {
		arg.ArgName.Accept(v)
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