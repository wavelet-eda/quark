package quark

type (
		Field struct {
			FieldType TypeExpr
			FieldName Name
		}

		InterfaceField struct {
			Direction InterfaceDirection
			Field

			directionKw ObjectPosition
		}

		RealName struct {
			text string

			firstChar ObjectPosition
			lastChar ObjectPosition
		}

		ConstructorField struct {
			name Name
		}

		InterfaceDirection = bool
)

const (
	Forward InterfaceDirection = true
	Reverse InterfaceDirection = false
)


func (n RealName) Start() *ObjectPosition {
	return &n.firstChar
}

func (f *Field) Start() *ObjectPosition {
	return f.FieldType.Start()
}

func (f *InterfaceField) Start() *ObjectPosition {
	return &f.directionKw
}



func (n RealName) End() *ObjectPosition {
	return n.lastChar.Next()
}

func (f *Field) End() *ObjectPosition {
	return f.FieldName.End()
}


