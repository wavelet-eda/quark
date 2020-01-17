package quark

type (
	//An atomic pattern of a single name
	AtomicPattern struct {
		Atom RealName
	}

	//A pattern which matches a type
	TypePattern struct {
		TypeName Name
		Params []Pattern

		bBrace ObjectPosition
	}

	ArrayPattern struct {
		Values []Pattern

		openBrace ObjectPosition
		closeBrace ObjectPosition
	}

	LiteralPattern struct {
		Value Literal
	}

	StructPattern struct {

	}
)
