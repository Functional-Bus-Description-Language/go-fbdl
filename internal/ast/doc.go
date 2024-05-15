package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Doc represents a documentation comment.
type Doc struct {
	Lines []tok.Comment
}

// endLine returns line number of the last line in the documentation comment.
// If doc has no lines 0 is returned.
func (d Doc) endLine() int {
	if len(d.Lines) == 0 {
		return 0
	}
	return d.Lines[len(d.Lines)-1].Line()
}

func (d Doc) Text(src []byte) string {
	text := ""
	for i, l := range d.Lines {
		t := tok.Text(l, src)
		start := 1
		if len(t) > 1 {
			if t[1] == ' ' {
				start = 2
			}
		}

		if len(t) > 2 {
			text += t[start:]
		}

		if i < len(d.Lines)-1 {
			text += "\n"
		}
	}
	return text
}

func buildDoc(ctx *context) Doc {
	doc := Doc{}
	doc.Lines = append(doc.Lines, ctx.tok().(tok.Comment))

	prevNewline := false
	for {
		ctx.idx++
		switch t := ctx.tok().(type) {
		case tok.Newline:
			if prevNewline {
				break
			} else {
				prevNewline = true
			}
		case tok.Comment:
			doc.Lines = append(doc.Lines, t)
			prevNewline = false
		default:
			return doc
		}
	}
}
