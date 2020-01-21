package quark

type (
		//A definition of a field inside a struct.
		Field struct {
			FieldType TypeExpr
			FieldName Name

			IsFuture bool

			KwFuture ObjectPosition
		}

		//An argument to a module, function, interace, or constructor invocation. If the argument is named
		//the FieldName and ValueExpr fields must both be non-nil. If the argument is by position, only
		//the ValueExpr will be non-nil.
		CallArgument struct {
			FieldName Name
			ValueExpr Expr
		}

)

func (f *Field) Start() *ObjectPosition {
	if f.IsFuture {
		return &f.KwFuture
	} else {
		return f.FieldType.Start()
	}
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


//Accept impl
func (f *Field) Accept(v Visitor) {
	if v.Visit(f) == nil {
		return
	}

	f.FieldType.Accept(v)
	f.FieldName.Accept(v)
}

func (c *CallArgument) Accept(v Visitor) {
	if v.Visit(c) == nil {
		return
	}

	if c.FieldName != nil {
		c.FieldName.Accept(v)
	}

	c.ValueExpr.Accept(v)
}

