package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Expr interface {
	exprNode()
}

// Expression nodes
type (
	BinaryExpr struct {
		X  Expr
		Op token.Operator
		Y  Expr
	}

	Bool struct {
		Val token.Bool
	}

	CallExpr struct {
		Name   token.Ident
		Lparen token.LeftParen
		Args   []Expr
		Rparen token.RightParen
	}

	Ident struct {
		Name token.Token
	}

	Int struct {
		Val token.Int
	}

	Real struct {
		Val token.Real
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

func (be BinaryExpr) exprNode() {}
func (b Bool) exprNode()        {}
func (c CallExpr) exprNode()    {}
func (i Ident) exprNode()       {}
func (i Int) exprNode()         {}
func (r Real) exprNode()        {}
func (ue UnaryExpr) exprNode()  {}
func (pe ParenExpr) exprNode()  {}

func (c CallExpr) eq(c2 CallExpr) bool {
	if c.Name != c2.Name ||
		c.Lparen != c2.Lparen ||
		c.Rparen != c2.Rparen {
		return false
	}

	if len(c.Args) != len(c2.Args) {
		return false
	}

	for i := range c.Args {
		if c.Args[i] != c2.Args[i] {
			return false
		}
	}

	return true
}
