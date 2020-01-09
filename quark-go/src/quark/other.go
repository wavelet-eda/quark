package quark


type Field struct {

}

type RealName struct {
	text string

	firstChar ObjectPosition
	lastChar ObjectPosition
}

func (n RealName) Start() ObjectPosition {
	return n.firstChar
}

func (n RealName) End() ObjectPosition {
	return n.lastChar.Next()
}

type ConstructorField struct {
	name Name
}
