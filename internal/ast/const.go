package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Const struct {
	Doc  Doc
	Name token.Ident
	Expr Expr
}

func (c Const) eq(c2 Const) bool {
	return c.Doc.eq(c2.Doc) && c.Name == c2.Name && c.Expr == c2.Expr
}

func buildConst(toks []token.Token, c *ctx) ([]Const, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident:
		return buildSingleConst(toks, c)
	case token.Newline:
		return buildMultiConst(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []token.Token, c *ctx) ([]Const, error) {
	con := Const{Name: toks[c.i+1].(token.Ident)}

	c.i += 2
	if t, ok := toks[c.i].(token.Ass); !ok {
		return nil, unexpected(t, "=")
	}

	c.i++
	expr, err := buildExpr(toks, c, nil)
	if err != nil {
		return nil, err
	}
	con.Expr = expr

	return []Const{con}, nil
}

func buildMultiConst(toks []token.Token, c *ctx) ([]Const, error) {
	consts := []Const{}
	con := Const{}

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
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[c.i].(type) {
			case token.Ident:
				con.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "=")
			}
		case Exp:
			var (
				err  error
				expr Expr
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			con.Expr = expr
			consts = append(consts, con)
			con = Const{}
			c.i--
			state = Id
		case Id:
			switch t := toks[c.i].(type) {
			case token.Ident:
				con.Name = t
				state = Ass
			case token.Comment:
				doc := buildDoc(toks, c)
				con.Doc = doc
				c.i--
			case token.Newline:
				continue
			case token.Dedent, token.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or dedent")
			}
		}
	}

	return consts, nil
}
