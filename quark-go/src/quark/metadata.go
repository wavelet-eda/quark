package quark

import "os"

type QuarkFile struct {
	File os.File
}

type FilePosition struct {
	Line uint32
	Character uint32
}

type ObjectPosition struct {
	File *QuarkFile
	Position FilePosition
}