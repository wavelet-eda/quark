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

		//A field as used in a struct constructor.
		ConstructorField struct {
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

func (c *ConstructorField) Start() *ObjectPosition {
	return c.FieldName.Start()
}


func (f *Field) End() *ObjectPosition {
	return f.FieldName.End()
}

func (c *ConstructorField) End() *ObjectPosition {
	return c.ValueExpr.End()
}


