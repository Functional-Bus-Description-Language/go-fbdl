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
