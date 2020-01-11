package quark

type (
	//An unqualified name.
	RealName struct {
		Text string

		firstChar ObjectPosition
		lastChar  ObjectPosition
	}

	//A qualified name.
	QualifiedName struct {
		Parts []RealName
	}
)

func NewRealName(text string, firstChar ObjectPosition, lastChar ObjectPosition) RealName {
	return RealName {
		Text: text,
		firstChar: firstChar,
		lastChar: lastChar,
	}
}


func (n *RealName) Start() *ObjectPosition {
	return &n.firstChar
}

func (n *QualifiedName) Start() *ObjectPosition {
	return n.Parts[0].Start()
}

func (n *RealName) End() *ObjectPosition {
	return n.lastChar.Next()
}

func (n *QualifiedName) End() *ObjectPosition {
	return n.Parts[len(n.Parts) - 1].End()
}

func (n *RealName) nameNode() {}
func (n *QualifiedName) nameNode() {}
