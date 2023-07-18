package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Build builds ast based on token stream.
func Build(s token.Stream) (File, error) {
	var (
		err error
		f   File
		i   int // Current index in the stream
	)

	for {
		if s[i].Kind == token.EOF {
			break
		}

		t := s[i]
		switch t.Kind {
		case token.NEWLINE:
			i++
		case token.COMMENT:
			i = buildComment(s, i, &f)
		case token.IMPORT:
			i, err = buildImport(s, i, &f)
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", t.Loc(), t.Kind))
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}

func buildComment(s token.Stream, i int, f *File) int {
	c := Comment{}
	c.add(s[i])

	prevNewline := false
	for {
		i++
		t := s[i]
		k := t.Kind
		if k == token.NEWLINE && prevNewline {
			break
		} else if k == token.NEWLINE {
			prevNewline = true
		} else if k == token.COMMENT {
			c.add(t)
			prevNewline = false
		} else {
			break
		}
	}

	f.Comments = append(f.Comments, c)

	return i
}

func buildImport(s token.Stream, i int, f *File) (int, error) {
	i++
	t := s[i]
	switch t.Kind {
	case token.IDENT, token.STRING:
		return buildSingleImport(s, i, f)
	default:
		return 0, fmt.Errorf(
			"%s: unexpected %s, expected identifier, string or newline",
			t.Loc(), t.Kind,
		)
	}
}

func buildSingleImport(s token.Stream, i int, f *File) (int, error) {
	si := SingleImport{Import: s[i]}

	t := s[i]
	switch t.Kind {
	case token.IDENT:
		si.Name = s[i]
		i++
		t = s[i]
		switch t.Kind {
		case token.STRING:
			si.Path = s[i]
			i++
		default:
			return 0, fmt.Errorf(
				"%s: unexpected %s, expected string", t.Loc(), t.Kind,
			)
		}
	case token.STRING:
		si.Path = s[i]
		i++
	}

	f.Imports = append(f.Imports, si)

	return i, nil
}
