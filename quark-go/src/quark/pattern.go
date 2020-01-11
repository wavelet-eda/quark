package quark

type (
	AtomicPattern struct {
		Atom RealName
	}

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
