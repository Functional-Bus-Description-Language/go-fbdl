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

	Int struct {
		Val token.Int
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
func (i Int) exprNode()        {}
func (ue UnaryExpr) exprNode() {}
func (pe ParenExpr) exprNode() {}

func buildExpr(s []token.Token, i int) (int, Expr, error) {
	switch t := s[i].(type) {
	case token.Neg, token.Sub, token.Add:
		return buildUnaryExpr(s, i)
	case token.Ident:
		return buildIdent(s, i)
	case token.Int:
		return buildInt(s, i)
	default:
		return 0, Ident{}, fmt.Errorf(
			"%s: unexpected %s, expected expression",
			token.Loc(t), t.Kind(),
		)
	}
}

func buildUnaryExpr(s []token.Token, i int) (int, UnaryExpr, error) {
	un := UnaryExpr{Op: s[i]}
	i, x, err := buildExpr(s, i+1)
	if err != nil {
		return 0, un, err
	}
	un.X = x
	return i, un, nil
}

func buildIdent(s []token.Token, i int) (int, Ident, error) {
	id := Ident{Name: s[i]}
	return i + 1, id, nil
}

func buildInt(s []token.Token, i int) (int, Int, error) {
	int_ := Int{Val: s[i].(token.Int)}
	return i + 1, int_, nil
}
