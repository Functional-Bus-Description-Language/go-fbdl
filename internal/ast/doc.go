package ast

import (
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Doc represents a documentation comment.
type Doc struct {
	Lines []token.Comment
}

func (d Doc) Text(src []byte) string {
	var text strings.Builder
	for i, l := range d.Lines {
		t := l.Text(src)
		start := 1
		if len(t) > 1 {
			if t[1] == ' ' {
				start = 2
			}
		}

		if len(t) > start {
			text.WriteString(t[start:])
		}

		if i < len(d.Lines)-1 {
			text.WriteByte('\n')
		}
	}
	return text.String()
}

func (d Doc) isEmpty() bool {
	return len(d.Lines) == 0
}

func emptyDoc() Doc {
	return Doc{}
}

func buildDoc(tokens *tokenStream) Doc {
	doc := Doc{}
	doc.Lines = append(doc.Lines, tokens.next().(token.Comment))

	for {
		switch tokens.peek(0).(type) {
		case token.Newline:
			tokens.next()
			if _, ok := tokens.peek(0).(token.Newline); ok {
				return emptyDoc()
			}
		case token.Comment:
			doc.Lines = append(doc.Lines, tokens.next().(token.Comment))
		default:
			return doc
		}
	}
}
