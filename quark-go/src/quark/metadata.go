package quark

import "os"

//Metadata about a file participating in this execution of quarkc
type QuarkFile struct {
	File os.File
}

//A line and character position in a specific file.
type FilePosition struct {
	Line uint32
	Character uint32
}


//A position of a compiler object (AST Node or Lexer Token) globally
//for this run of the compiler.
type ObjectPosition struct {
	File *QuarkFile
	Position FilePosition
}

func (o ObjectPosition) Next() *ObjectPosition {
	return &ObjectPosition{
		File:     o.File,
		Position: FilePosition{
			Line: o.Position.Line,
			Character: o.Position.Character,
		},
	}
}