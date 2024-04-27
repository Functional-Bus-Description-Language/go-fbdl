package ast

// File represents .fbd file.
type File struct {
	Imports []Import
	Consts  []Const
	Insts   []Inst
	Types   []Type
}
