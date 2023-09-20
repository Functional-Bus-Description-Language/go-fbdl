package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The const struct represents constant.
type Const struct {
	Doc   Doc
	Name  tok.Ident
	Value Expr
}

func (c Const) eq(c2 Const) bool {
	return c.Doc.eq(c2.Doc) && c.Name == c2.Name && c.Value == c2.Value
}

func buildConst(toks []tok.Token, c *ctx) ([]Const, error) {
	switch t := toks[c.i+1].(type) {
	case tok.Ident:
		return buildSingleConst(toks, c)
	case tok.Newline:
		return buildMultiConst(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []tok.Token, c *ctx) ([]Const, error) {
	con := Const{Name: toks[c.i+1].(tok.Ident)}

	c.i += 2
	if _, ok := toks[c.i].(tok.Ass); !ok {
		return nil, unexpected(toks[c.i], "'='")
	}

	c.i++
	expr, err := buildExpr(toks, c, nil)
	if err != nil {
		return nil, err
	}
	con.Value = expr

	return []Const{con}, nil
}

func buildMultiConst(toks []tok.Token, c *ctx) ([]Const, error) {
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
			case tok.Newline:
				continue
			case tok.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[c.i].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[c.i].(type) {
			case tok.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(toks, c, nil)
			if err != nil {
				return nil, err
			}
			con.Value = expr
			consts = append(consts, con)
			con = Const{}
			c.i--
			state = Id
		case Id:
			switch t := toks[c.i].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			case tok.Comment:
				doc := buildDoc(toks, c)
				con.Doc = doc
				c.i--
			case tok.Newline:
				continue
			case tok.Dedent:
				c.i++
				break tokenLoop
			case tok.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or dedent")
			}
		}
	}

	return consts, nil
}
