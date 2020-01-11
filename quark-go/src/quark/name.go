package quark

type (
	//An unqualified name.
	RealName struct {
		Text string

		pos ObjectPosition
	}

	//A qualified name.
	QualifiedName struct {
		Parts []*RealName
	}
)

func NewRealName(text string, pos ObjectPosition) *RealName {
	return &RealName {
		Text: text,
		pos: pos,
	}
}


func (n *RealName) Start() *ObjectPosition {
	return &n.pos
}

func (n *QualifiedName) Start() *ObjectPosition {
	return n.Parts[0].Start()
}

func (n *RealName) End() *ObjectPosition {
	return n.pos.Next()
}

func (n *QualifiedName) End() *ObjectPosition {
	return n.Parts[len(n.Parts) - 1].End()
}

func (n *RealName) nameNode() {}
func (n *QualifiedName) nameNode() {}
