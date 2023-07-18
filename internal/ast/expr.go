package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Expr interface {
	exprNode()
}

// Expression nodes
type (
	Ident struct {
		Name token.Token
	}

	UnaryExpr struct {
		Op token.Token
		X  Expr
	}

	ParenExpr struct {
		Lparen token.Token
		X      Expr
		Rparen token.Token
	}
)

func (i Ident) exprNode()      {}
func (ue UnaryExpr) exprNode() {}
func (pe ParenExpr) exprNode() {}

func buildExpr(s token.Stream, i int) (int, Expr, error) {
	t := s[i]
	k := t.Kind
	if k == token.NEG || k == token.SUB || k == token.ADD {
		return buildUnaryExpr(s, i)
	} else if k == token.IDENT {
		return buildIdent(s, i)
	}

	return 0, Ident{}, fmt.Errorf(
		"%s: unexpected %s, expected expression", t.Loc(), k,
	)
}

func buildUnaryExpr(s token.Stream, i int) (int, UnaryExpr, error) {
	un := UnaryExpr{Op: s[i]}
	i, x, err := buildExpr(s, i+1)
	if err != nil {
		return 0, un, err
	}
	un.X = x
	return i, un, nil
}

func buildIdent(s token.Stream, i int) (int, Ident, error) {
	id := Ident{Name: s[i]}
	return i, id, nil
}
