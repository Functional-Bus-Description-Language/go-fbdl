package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Expr interface represents generic expression.
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

	BitString struct {
		X token.BitString
	}

	Bool struct {
		X token.Bool
	}

	// Function Call
	Call struct {
		Name token.Ident
		Args []Expr
	}

	List struct {
		Xs []Expr
	}

	Ident struct {
		Name token.Token
	}

	Int struct {
		X token.Int
	}

	Real struct {
		X token.Real
	}

	String struct {
		X token.String
	}

	UnaryExpr struct {
		Op token.Token
		X  Expr
	}

	ParenExpr struct {
		X Expr
	}
)

func (be BinaryExpr) exprNode() {}
func (bs BitString) exprNode()  {}
func (b Bool) exprNode()        {}
func (c Call) exprNode()        {}
func (l List) exprNode()        {}
func (i Ident) exprNode()       {}
func (i Int) exprNode()         {}
func (r Real) exprNode()        {}
func (s String) exprNode()      {}
func (ue UnaryExpr) exprNode()  {}
func (pe ParenExpr) exprNode()  {}

func (c Call) eq(c2 Call) bool {
	if c.Name != c2.Name {
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

// leftOp is the operator on the left side of the expression.
func buildExpr(toks []token.Token, c *ctx, leftOp token.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := toks[c.i].(type) {
	case token.Neg, token.Sub, token.Add:
		expr, err = buildUnaryExpr(toks, c)
	case token.Ident:
		switch toks[c.i+1].(type) {
		case token.LeftParen:
			expr, err = buildCallExpr(toks, c)
		default:
			expr, err = buildIdent(toks, c)
		}
	case token.Bool:
		expr, err = buildBool(toks, c)
	case token.Int:
		expr, err = buildInt(toks, c)
	case token.Real:
		expr, err = buildReal(toks, c)
	case token.String:
		expr, err = buildString(toks, c)
	case token.BitString:
		expr, err = buildBitString(toks, c)
	case token.LeftParen:
		expr, err = buildParenExpr(toks, c)
	case token.LeftBracket:
		expr, err = buildList(toks, c)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp token.Operator
		if op, ok := toks[c.i].(token.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			c.i++
			expr, err = buildExpr(toks, c, rightOp)
			if err != nil {
				return expr, err
			}
			be.Y = expr
			expr = be
		} else if leftOp.Precedence() >= rightOp.Precedence() {
			return expr, nil
		}
	}
}

func buildIdent(toks []token.Token, c *ctx) (Ident, error) {
	id := Ident{Name: toks[c.i]}
	c.i++
	return id, nil
}

func buildBool(toks []token.Token, c *ctx) (Bool, error) {
	b := Bool{toks[c.i].(token.Bool)}
	c.i++
	return b, nil
}

func buildInt(toks []token.Token, c *ctx) (Int, error) {
	int_ := Int{toks[c.i].(token.Int)}
	c.i++
	return int_, nil
}

func buildReal(toks []token.Token, c *ctx) (Real, error) {
	r := Real{toks[c.i].(token.Real)}
	c.i++
	return r, nil
}

func buildString(toks []token.Token, c *ctx) (String, error) {
	s := String{toks[c.i].(token.String)}
	c.i++
	return s, nil
}

func buildBitString(toks []token.Token, c *ctx) (BitString, error) {
	s := BitString{toks[c.i].(token.BitString)}
	c.i++
	return s, nil
}

func buildParenExpr(toks []token.Token, c *ctx) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)
	c.i++
	expr, err = buildExpr(toks, c, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if _, ok := toks[c.i].(token.RightParen); ok {
		c.i++
	} else {
		return pe, unexpected(toks[c.i], ")")
	}

	return pe, nil
}

func buildList(toks []token.Token, c *ctx) (List, error) {
	l := List{}
	prevExpr := false
	lbi := c.i // Left bracket token index
	c.i++

tokenLoop:
	for {
		switch t := toks[c.i].(type) {
		case token.RightBracket:
			c.i++
			break tokenLoop
		case token.Comma:
			if c.i == lbi+1 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return l, unexpected(t, ", or ]")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return l, err
			}
			l.Xs = append(l.Xs, expr)
			prevExpr = true
		}
	}

	return l, nil
}

func buildCallExpr(toks []token.Token, c *ctx) (Call, error) {
	call := Call{Name: toks[c.i].(token.Ident)}
	lpi := c.i // Left parenthesis token index
	c.i += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := toks[c.i].(type) {
		case token.RightParen:
			c.i++
			break tokenLoop
		case token.Comma:
			if c.i == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return call, unexpected(t, ", or )")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(toks, c, nil)
			if err != nil {
				return call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return call, nil
}

func buildUnaryExpr(toks []token.Token, c *ctx) (UnaryExpr, error) {
	un := UnaryExpr{Op: toks[c.i]}
	c.i++
	x, err := buildExpr(toks, c, nil)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
