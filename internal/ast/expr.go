package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Expr interface represents generic expression.
//
// The Tok method returns token which position spans the whole expression.
// It is useful for error messages.
// No assumptions shall be made on the returned token type.
type Expr interface {
	expr()
	Tok() token.Token
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
		LBracket token.LBracket
		Xs       []Expr
		RBracket token.RBracket
	}

	Ident struct {
		Name token.Token
	}

	QualIdent struct {
		Name token.Token
	}

	Int struct {
		X token.Int
	}

	Float struct {
		X token.Float
	}

	String struct {
		X token.String
	}

	Time struct {
		X token.Time
	}

	UnaryExpr struct {
		Op token.Token
		X  Expr
	}

	ParenExpr struct {
		LParen token.LParen
		X      Expr
		RParen token.RParen
	}
)

func (be BinaryExpr) expr()            {}
func (be BinaryExpr) Tok() token.Token { return token.Join(be.X.Tok(), be.Y.Tok()) }

func (bs BitString) expr()            {}
func (bs BitString) Tok() token.Token { return bs.X }

func (b Bool) expr()            {}
func (b Bool) Tok() token.Token { return b.X }

func (c Call) expr()            {}
func (c Call) Tok() token.Token { return c.Name }

func (l List) expr()            {}
func (l List) Tok() token.Token { return token.Join(l.LBracket, l.RBracket) }

func (i Ident) expr()            {}
func (i Ident) Tok() token.Token { return i.Name }

func (qi QualIdent) expr()            {}
func (qi QualIdent) Tok() token.Token { return qi.Name }

func (i Int) expr()            {}
func (i Int) Tok() token.Token { return i.X }

func (f Float) expr()            {}
func (f Float) Tok() token.Token { return f.X }

func (s String) expr()            {}
func (s String) Tok() token.Token { return s.X }

func (t Time) expr()            {}
func (t Time) Tok() token.Token { return t.X }

func (ue UnaryExpr) expr()            {}
func (ue UnaryExpr) Tok() token.Token { return token.Join(ue.Op, ue.X.Tok()) }

func (pe ParenExpr) expr()            {}
func (pe ParenExpr) Tok() token.Token { return token.Join(pe.LParen, pe.RParen) }

// leftOp is the operator on the left side of the expression.
func buildExpr(ctx *context, leftOp token.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := ctx.token().(type) {
	case token.Neg, token.Sub, token.Add:
		expr, err = buildUnaryExpr(ctx)
	case token.Ident:
		switch ctx.nextTok().(type) {
		case token.LParen:
			expr, err = buildCallExpr(ctx)
		default:
			expr, err = buildIdent(ctx)
		}
	case token.QualIdent:
		switch ctx.nextTok().(type) {
		case token.LParen:
			expr, err = buildCallExpr(ctx)
		default:
			expr, err = buildQualIdent(ctx)
		}
	case token.Bool:
		expr, err = buildBool(ctx)
	case token.Int:
		expr, err = buildInt(ctx)
	case token.Float:
		expr, err = buildFloat(ctx)
	case token.String:
		expr, err = buildString(ctx)
	case token.Time:
		expr, err = buildTime(ctx)
	case token.BitString:
		expr, err = buildBitString(ctx)
	case token.LParen:
		expr, err = buildParenExpr(ctx)
	case token.LBracket:
		expr, err = buildList(ctx)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp token.Operator
		if op, ok := ctx.token().(token.Operator); ok {
			rightOp = op
		} else {
			return expr, nil
		}

		if (leftOp == nil) ||
			(leftOp != nil && (leftOp.Precedence() < rightOp.Precedence())) {
			be := BinaryExpr{X: expr, Op: rightOp}
			ctx.idx++
			expr, err = buildExpr(ctx, rightOp)
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

func buildIdent(ctx *context) (Ident, error) {
	id := Ident{Name: ctx.token()}
	ctx.idx++
	return id, nil
}

func buildQualIdent(ctx *context) (QualIdent, error) {
	id := QualIdent{Name: ctx.token()}
	ctx.idx++
	return id, nil
}

func buildBool(ctx *context) (Bool, error) {
	b := Bool{ctx.token().(token.Bool)}
	ctx.idx++
	return b, nil
}

func buildInt(ctx *context) (Int, error) {
	int_ := Int{ctx.token().(token.Int)}
	ctx.idx++
	return int_, nil
}

func buildFloat(ctx *context) (Float, error) {
	r := Float{ctx.token().(token.Float)}
	ctx.idx++
	return r, nil
}

func buildString(ctx *context) (String, error) {
	s := String{ctx.token().(token.String)}
	ctx.idx++
	return s, nil
}

func buildTime(ctx *context) (Time, error) {
	t := Time{ctx.token().(token.Time)}
	ctx.idx++
	return t, nil
}

func buildBitString(ctx *context) (BitString, error) {
	s := BitString{ctx.token().(token.BitString)}
	ctx.idx++
	return s, nil
}

func buildParenExpr(ctx *context) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)

	pe.LParen = ctx.token().(token.LParen)

	ctx.idx++
	expr, err = buildExpr(ctx, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if rp, ok := ctx.token().(token.RParen); ok {
		pe.RParen = rp
		ctx.idx++
	} else {
		return pe, unexpected(ctx.token(), "')'")
	}

	return pe, nil
}

func buildList(ctx *context) (List, error) {
	l := List{}
	prevExpr := false
	l.LBracket = ctx.token().(token.LBracket)
	lbi := ctx.idx // Left bracket token index
	ctx.idx++

tokenLoop:
	for {
		switch t := ctx.token().(type) {
		case token.RBracket:
			l.RBracket = t
			ctx.idx++
			break tokenLoop
		case token.Comma:
			if ctx.idx == lbi+1 {
				return l, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.idx++
		default:
			if prevExpr {
				return l, unexpected(t, "',' or ']'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(ctx, nil)
			if err != nil {
				return l, err
			}
			l.Xs = append(l.Xs, expr)
			prevExpr = true
		}
	}

	return l, nil
}

func buildCallExpr(ctx *context) (Call, error) {
	call := Call{Name: ctx.token().(token.Ident)}
	lpi := ctx.idx // Left parenthesis token index
	ctx.idx += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := ctx.token().(type) {
		case token.RParen:
			ctx.idx++
			break tokenLoop
		case token.Comma:
			if ctx.idx == lpi+2 {
				return call, unexpected(t, "expression")
			}
			prevExpr = false
			ctx.idx++
		default:
			if prevExpr {
				return call, unexpected(t, "',' or ')'")
			}

			var (
				expr Expr
				err  error
			)
			expr, err = buildExpr(ctx, nil)
			if err != nil {
				return call, err
			}
			call.Args = append(call.Args, expr)
			prevExpr = true
		}
	}

	return call, nil
}

func buildUnaryExpr(ctx *context) (UnaryExpr, error) {
	op := ctx.token().(token.Operator)
	un := UnaryExpr{Op: op}
	ctx.idx++
	x, err := buildExpr(ctx, op)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
