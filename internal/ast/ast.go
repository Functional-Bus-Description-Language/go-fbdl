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
		case token.CONST:
			i, err = buildConst(s, i, &f)
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
	c.Comments = append(c.Comments, s[i])

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
			c.Comments = append(c.Comments, t)
			prevNewline = false
		} else {
			break
		}
	}

	f.Comments = append(f.Comments, c)

	return i
}

func buildConst(s token.Stream, i int, f *File) (int, error) {
	t := s[+1]
	switch t.Kind {
	case token.IDENT:
		return buildSingleConst(s, i, f)
	case token.NEWLINE:
		panic("buildMultiConst")
	default:
		return 0, fmt.Errorf(
			"%s: unexpected %s, expected identifier, string or newline",
			t.Loc(), t.Kind,
		)
	}
}

func buildSingleConst(s token.Stream, i int, f *File) (int, error) {
	c := SingleConst{Const: s[i], Name: s[i+1]}

	i += 2
	t := s[i]
	if t.Kind != token.ASS {
		return 0, fmt.Errorf("%s: unexpected %s, expected =", t.Loc(), t.Kind)
	}
	c.Ass = t

	i++
	i, expr, err := buildExpr(s, i)
	if err != nil {
		return 0, err
	}
	c.Expr = expr

	f.Consts = append(f.Consts, c)

	return i, nil
}

func buildImport(s token.Stream, i int, f *File) (int, error) {
	t := s[i+1]
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

	i++
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
