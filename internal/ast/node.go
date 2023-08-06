package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Doc struct {
	Lines []token.Comment
}

// endLine returns line number of the last line in the documentation comment.
// If doc has no lines 0 is returned.
func (d Doc) endLine() int {
	if len(d.Lines) == 0 {
		return 0
	}
	return d.Lines[len(d.Lines)-1].Line()
}

func (d Doc) eq(d2 Doc) bool {
	if len(d.Lines) != len(d2.Lines) {
		return false
	}
	for i := range d.Lines {
		if d.Lines[i] != d2.Lines[i] {
			return false
		}
	}
	return true
}

type Import interface {
	importNode()
}

// Import types
type (
	SingleImport struct {
		Name token.Ident
		Path token.String
	}
)

func (si SingleImport) importNode() {}

type File struct {
	Imports []Import
	Consts  []Const
}
