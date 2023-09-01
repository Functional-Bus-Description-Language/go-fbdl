package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Expr interface represents generic expression.
type Expr interface {
	exprNode()
}

// Expression nodes
type (
	BinaryExpr struct {
		X  Expr
		Op tok.Operator
		Y  Expr
	}

	BitString struct {
		X tok.BitString
	}

	Bool struct {
		X tok.Bool
	}

	// Function Call
	Call struct {
		Name tok.Ident
		Args []Expr
	}

	List struct {
		Xs []Expr
	}

	Ident struct {
		Name tok.Token
	}

	Int struct {
		X tok.Int
	}

	Real struct {
		X tok.Real
	}

	String struct {
		X tok.String
	}

	Time struct {
		X tok.Time
	}

	UnaryExpr struct {
		Op tok.Token
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
func (t Time) exprNode()        {}
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
func buildExpr(toks []tok.Token, c *ctx, leftOp tok.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := toks[c.i].(type) {
	case tok.Neg, tok.Sub, tok.Add:
		expr, err = buildUnaryExpr(toks, c)
	case tok.Ident:
		switch toks[c.i+1].(type) {
		case tok.LeftParen:
			expr, err = buildCallExpr(toks, c)
		default:
			expr, err = buildIdent(toks, c)
		}
	case tok.Bool:
		expr, err = buildBool(toks, c)
	case tok.Int:
		expr, err = buildInt(toks, c)
	case tok.Real:
		expr, err = buildReal(toks, c)
	case tok.String:
		expr, err = buildString(toks, c)
	case tok.Time:
		expr, err = buildTime(toks, c)
	case tok.BitString:
		expr, err = buildBitString(toks, c)
	case tok.LeftParen:
		expr, err = buildParenExpr(toks, c)
	case tok.LeftBracket:
		expr, err = buildList(toks, c)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp tok.Operator
		if op, ok := toks[c.i].(tok.Operator); ok {
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

func buildIdent(toks []tok.Token, c *ctx) (Ident, error) {
	id := Ident{Name: toks[c.i]}
	c.i++
	return id, nil
}

func buildBool(toks []tok.Token, c *ctx) (Bool, error) {
	b := Bool{toks[c.i].(tok.Bool)}
	c.i++
	return b, nil
}

func buildInt(toks []tok.Token, c *ctx) (Int, error) {
	int_ := Int{toks[c.i].(tok.Int)}
	c.i++
	return int_, nil
}

func buildReal(toks []tok.Token, c *ctx) (Real, error) {
	r := Real{toks[c.i].(tok.Real)}
	c.i++
	return r, nil
}

func buildString(toks []tok.Token, c *ctx) (String, error) {
	s := String{toks[c.i].(tok.String)}
	c.i++
	return s, nil
}

func buildTime(toks []tok.Token, c *ctx) (Time, error) {
	t := Time{toks[c.i].(tok.Time)}
	c.i++
	return t, nil
}

func buildBitString(toks []tok.Token, c *ctx) (BitString, error) {
	s := BitString{toks[c.i].(tok.BitString)}
	c.i++
	return s, nil
}

func buildParenExpr(toks []tok.Token, c *ctx) (ParenExpr, error) {
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

	if _, ok := toks[c.i].(tok.RightParen); ok {
		c.i++
	} else {
		return pe, unexpected(toks[c.i], "')'")
	}

	return pe, nil
}

func buildList(toks []tok.Token, c *ctx) (List, error) {
	l := List{}
	prevExpr := false
	lbi := c.i // Left bracket token index
	c.i++

tokenLoop:
	for {
		switch t := toks[c.i].(type) {
		case tok.RightBracket:
			c.i++
			break tokenLoop
		case tok.Comma:
			if c.i == lbi+1 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return l, unexpected(t, "',' or ']'")
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

func buildCallExpr(toks []tok.Token, c *ctx) (Call, error) {
	call := Call{Name: toks[c.i].(tok.Ident)}
	lpi := c.i // Left parenthesis token index
	c.i += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := toks[c.i].(type) {
		case tok.RightParen:
			c.i++
			break tokenLoop
		case tok.Comma:
			if c.i == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			c.i++
		default:
			if prevExpr {
				return call, unexpected(t, "',' or ')'")
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

func buildUnaryExpr(toks []tok.Token, c *ctx) (UnaryExpr, error) {
	un := UnaryExpr{Op: toks[c.i]}
	c.i++
	x, err := buildExpr(toks, c, nil)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
