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

func buildDoc(toks []token.Token, c *ctx) Doc {
	doc := Doc{}
	doc.Lines = append(doc.Lines, toks[c.i].(token.Comment))

	prevNewline := false
	for {
		c.i++
		switch t := toks[c.i].(type) {
		case token.Newline:
			if prevNewline {
				break
			} else {
				prevNewline = true
			}
		case token.Comment:
			doc.Lines = append(doc.Lines, t)
			prevNewline = false
		default:
			return doc
		}
	}
}