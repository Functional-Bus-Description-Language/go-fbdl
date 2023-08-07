package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Const interface {
	constNode()
}

type SingleConst struct {
	Doc  Doc
	Name token.Ident
	Expr Expr
}

func (sc SingleConst) constNode() {}

func (sc SingleConst) eq(sc2 SingleConst) bool {
	return sc.Doc.eq(sc2.Doc) && sc.Name == sc2.Name && sc.Expr == sc2.Expr
}

type MultiConst struct {
	Consts []SingleConst
}

func (mc MultiConst) constNode() {}

func (mc MultiConst) eq(mc2 MultiConst) bool {
	if len(mc.Consts) != len(mc2.Consts) {
		return false
	}

	for i, c := range mc.Consts {
		if !c.eq(mc2.Consts[i]) {
			return false
		}
	}

	return true
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
	sc := SingleConst{}

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
				sc.Name = t
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
			sc.Expr = expr
			mc.Consts = append(mc.Consts, sc)
			sc = SingleConst{}
			c.i--
			state = Id
		case Id:
			switch t := toks[c.i].(type) {
			case token.Ident:
				sc.Name = t
				state = Ass
			case token.Comment:
				doc := buildDoc(toks, c)
				sc.Doc = doc
				c.i--
			case token.Newline:
				continue
			case token.Dedent, token.Eof:
				break tokenLoop
			default:
				return mc, unexpected(t, "identifier or dedent")
			}
		}
	}

	return mc, nil
}
