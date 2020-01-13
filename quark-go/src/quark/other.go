package quark

type (
		//A definition of a field inside a struct.
		Field struct {
			FieldType TypeExpr
			FieldName Name
		}

		//A definition of a field inside an interface.
		InterfaceField struct {
			Direction InterfaceDirection
			Field

			directionKw ObjectPosition
		}

		//An argument to a module, function, interace, or constructor invocation. If the argument is named
		//the FieldName and ValueExpr fields must both be non-nil. If the argument is by position, only
		//the ValueExpr will be non-nil.
		CallArgument struct {
			FieldName Name
			ValueExpr Expr
		}

		//The directionality of an interface field.
		InterfaceDirection = bool
)

const (
	Forward InterfaceDirection = true
	Reverse InterfaceDirection = false
)


func (f *Field) Start() *ObjectPosition {
	return f.FieldType.Start()
}

func (f *InterfaceField) Start() *ObjectPosition {
	return &f.directionKw
}

func (c *CallArgument) Start() *ObjectPosition {
	return c.FieldName.Start()
}


func (f *Field) End() *ObjectPosition {
	return f.FieldName.End()
}

func (c *CallArgument) End() *ObjectPosition {
	return c.ValueExpr.End()
}


