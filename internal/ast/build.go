package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Building context
type ctx struct {
	i int // Current token index
}

// Build builds ast based on token stream.
func Build(s []token.Token) (File, error) {
	var (
		err  error
		f    File
		c    ctx
		cmnt Comment
		con  Const
		imp  Import
	)

	for {
		if _, ok := s[c.i].(token.Eof); ok {
			break
		}

		switch t := s[c.i].(type) {
		case token.Newline:
			c.i++
		case token.Comment:
			cmnt = buildComment(s, &c)
			f.Comments = append(f.Comments, cmnt)
		case token.Const:
			con, err = buildConst(s, &c)
			f.Consts = append(f.Consts, con)
		case token.Import:
			imp, err = buildImport(s, &c)
			f.Imports = append(f.Imports, imp)
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", token.Loc(t), t.Kind()))
		}

		if err != nil {
			return f, err
		}
	}

	return f, nil
}

func buildComment(s []token.Token, c *ctx) Comment {
	cmnt := Comment{}
	cmnt.Comments = append(cmnt.Comments, s[c.i].(token.Comment))

	prevNewline := false
	for {
		c.i++
		switch t := s[c.i].(type) {
		case token.Newline:
			if prevNewline {
				break
			} else {
				prevNewline = true
			}
		case token.Comment:
			cmnt.Comments = append(cmnt.Comments, t)
			prevNewline = false
		default:
			return cmnt
		}
	}
}

func buildConst(s []token.Token, c *ctx) (Const, error) {
	switch t := s[c.i+1].(type) {
	case token.Ident:
		return buildSingleConst(s, c)
	case token.Newline:
		return buildMultiConst(s, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(s []token.Token, c *ctx) (SingleConst, error) {
	sc := SingleConst{Name: s[c.i+1].(token.Ident)}

	c.i += 2
	if t, ok := s[c.i].(token.Ass); !ok {
		return sc, unexpected(t, "=")
	}

	c.i++
	expr, err := buildExpr(s, c, nil)
	if err != nil {
		return sc, err
	}
	sc.Expr = expr

	return sc, nil
}

func buildMultiConst(s []token.Token, c *ctx) (MultiConst, error) {
	mc := MultiConst{}

	const (
		Indent int = iota
		FirstId
		Ass
		Exp
		Id
	)
	state := Indent

	c.i += 1
tokenLoop:
	for {
		c.i++
		switch state {
		case Indent:
			switch t := s[c.i].(type) {
			case token.Newline:
				continue
			case token.Indent:
				state = FirstId
			default:
				return mc, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := s[c.i].(type) {
			case token.Ident:
				mc.Names = append(mc.Names, t)
				state = Ass
			default:
				return mc, unexpected(t, "identifier")
			}
		case Ass:
			switch t := s[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return mc, unexpected(t, "=")
			}
		case Exp:
			var (
				err  error
				expr Expr
			)
			expr, err = buildExpr(s, c, nil)
			if err != nil {
				return mc, err
			}
			mc.Exprs = append(mc.Exprs, expr)
			state = Id
		case Id:
			switch t := s[c.i].(type) {
			case token.Ident:
				mc.Names = append(mc.Names, t)
				state = Ass
			case token.Newline:
				continue
			case token.Dedent:
				break tokenLoop
			default:
				return mc, unexpected(t, "identifier or dedent")
			}
		}
	}

	return mc, nil
}

func buildImport(s []token.Token, c *ctx) (Import, error) {
	switch t := s[c.i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(s, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(s []token.Token, c *ctx) (SingleImport, error) {
	si := SingleImport{}

	c.i++
	switch t := s[c.i].(type) {
	case token.Ident:
		si.Name = t
		c.i++
		switch t := s[c.i].(type) {
		case token.String:
			si.Path = t
			c.i++
		default:
			return si, unexpected(t, "string")
		}
	case token.String:
		si.Path = t
		c.i++
	}

	return si, nil
}
