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
func Build(toks []token.Token) (File, error) {
	var (
		err  error
		f    File
		c    ctx
		cmnt Comment
		con  Const
		imp  Import
	)

	for {
		if _, ok := toks[c.i].(token.Eof); ok {
			break
		}

		switch t := toks[c.i].(type) {
		case token.Newline:
			c.i++
		case token.Comment:
			cmnt = buildComment(toks, &c)
			f.Comments = append(f.Comments, cmnt)
		case token.Const:
			con, err = buildConst(toks, &c)
			f.Consts = append(f.Consts, con)
		case token.Import:
			imp, err = buildImport(toks, &c)
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

func buildComment(toks []token.Token, c *ctx) Comment {
	cmnt := Comment{}
	cmnt.Comments = append(cmnt.Comments, toks[c.i].(token.Comment))

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
			cmnt.Comments = append(cmnt.Comments, t)
			prevNewline = false
		default:
			return cmnt
		}
	}
}

func buildConst(toks []token.Token, c *ctx) (Const, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident:
		return buildSingleConst(toks, c)
	case token.Newline:
		return buildMultiConst(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []token.Token, c *ctx) (SingleConst, error) {
	sc := SingleConst{Name: toks[c.i+1].(token.Ident)}

	c.i += 2
	if t, ok := toks[c.i].(token.Ass); !ok {
		return sc, unexpected(t, "=")
	}

	c.i++
	expr, err := buildExpr(toks, c, nil)
	if err != nil {
		return sc, err
	}
	sc.Expr = expr

	return sc, nil
}

func buildMultiConst(toks []token.Token, c *ctx) (MultiConst, error) {
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
			switch t := toks[c.i].(type) {
			case token.Newline:
				continue
			case token.Indent:
				state = FirstId
			default:
				return mc, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[c.i].(type) {
			case token.Ident:
				mc.Names = append(mc.Names, t)
				state = Ass
			default:
				return mc, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[c.i].(type) {
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
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return mc, err
			}
			mc.Exprs = append(mc.Exprs, expr)
			state = Id
		case Id:
			switch t := toks[c.i].(type) {
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

func buildImport(toks []token.Token, c *ctx) (Import, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []token.Token, c *ctx) (SingleImport, error) {
	si := SingleImport{}

	c.i++
	switch t := toks[c.i].(type) {
	case token.Ident:
		si.Name = t
		c.i++
		switch t := toks[c.i].(type) {
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
