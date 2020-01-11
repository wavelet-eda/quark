package quark

//Metadata about a file participating in this execution of quarkc
type QuarkFile struct {
	FileName *string
}

//A position of a compiler object (AST Node or Lexer Token) globally
//for this run of the compiler.
type ObjectPosition struct {
	OriginFile *QuarkFile
	Index int
}

func NewObjectPosition(file *QuarkFile, index int) ObjectPosition {
	return ObjectPosition{
		OriginFile: file,
		Index:      index,
	}
}

func (o ObjectPosition) Next() *ObjectPosition {
	return &ObjectPosition{o.OriginFile, o.Index + 1}
}