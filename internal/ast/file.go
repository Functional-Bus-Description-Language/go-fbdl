package ast

type File struct {
	Imports []Import
	Consts  []Const
	Insts   []Inst
}
