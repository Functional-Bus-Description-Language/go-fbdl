package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Build builds ast based on token stream.
func Build(s []token.Token) (File, error) {
	var (
		err error
		f   File
		i   int // Current index in the stream
	)

	for {
		if _, ok := s[i].(token.Eof); ok {
			break
		}

		t := s[i]
		switch t.(type) {
		case token.Newline:
			i++
		case token.Comment:
			i = buildComment(s, i, &f)
		case token.Const:
			i, err = buildConst(s, i, &f)
		case token.Import:
			i, err = buildImport(s, i, &f)
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", token.Loc(t), t.Kind()))
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}

func buildComment(s []token.Token, i int, f *File) int {
	c := Comment{}
	c.Comments = append(c.Comments, s[i].(token.Comment))

	prevNewline := false
	for {
		i++
		switch t := s[i].(type) {
		case token.Newline:
			if prevNewline {
				break
			} else {
				prevNewline = true
			}
		case token.Comment:
			c.Comments = append(c.Comments, t)
			prevNewline = false
		default:
			f.Comments = append(f.Comments, c)
			return i
		}
	}
}

func buildConst(s []token.Token, i int, f *File) (int, error) {
	t := s[i+1]
	switch t.(type) {
	case token.Ident:
		return buildSingleConst(s, i, f)
	case token.Newline:
		panic("buildMultiConst")
	default:
		return 0, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(s []token.Token, i int, f *File) (int, error) {
	c := Const{Name: s[i+1].(token.Ident)}

	i += 2
	if t, ok := s[i].(token.Ass); !ok {
		return 0, unexpected(t, "=")
	}

	i++
	i, expr, err := buildExpr(s, i, nil)
	if err != nil {
		return 0, err
	}
	c.Expr = expr

	f.Consts = append(f.Consts, c)

	return i, nil
}

func buildImport(s []token.Token, i int, f *File) (int, error) {
	switch t := s[i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(s, i, f)
	default:
		return 0, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(s []token.Token, i int, f *File) (int, error) {
	si := SingleImport{}

	i++
	switch t := s[i].(type) {
	case token.Ident:
		si.Name = t
		i++
		switch t := s[i].(type) {
		case token.String:
			si.Path = t
			i++
		default:
			return 0, unexpected(t, "string")
		}
	case token.String:
		si.Path = t
		i++
	}

	f.Imports = append(f.Imports, si)

	return i, nil
}
