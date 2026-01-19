package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Expr interface represents generic expression.
//
// The Tok method returns token which position spans the whole expression.
// It is useful for error messages.
// No assumptions shall be made on the returned token type.
type Expr interface {
	expr()
	Tok() tok.Token
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
		LBracket tok.LBracket
		Xs       []Expr
		RBracket tok.RBracket
	}

	Ident struct {
		Name tok.Token
	}

	QualIdent struct {
		Name tok.Token
	}

	Int struct {
		X tok.Int
	}

	Float struct {
		X tok.Float
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
		LParen tok.LParen
		X      Expr
		RParen tok.RParen
	}
)

func (be BinaryExpr) expr()          {}
func (be BinaryExpr) Tok() tok.Token { return tok.Join(be.X.Tok(), be.Y.Tok()) }

func (bs BitString) expr()          {}
func (bs BitString) Tok() tok.Token { return bs.X }

func (b Bool) expr()          {}
func (b Bool) Tok() tok.Token { return b.X }

func (c Call) expr()          {}
func (c Call) Tok() tok.Token { return c.Name }

func (l List) expr()          {}
func (l List) Tok() tok.Token { return tok.Join(l.LBracket, l.RBracket) }

func (i Ident) expr()          {}
func (i Ident) Tok() tok.Token { return i.Name }

func (qi QualIdent) expr()          {}
func (qi QualIdent) Tok() tok.Token { return qi.Name }

func (i Int) expr()          {}
func (i Int) Tok() tok.Token { return i.X }

func (f Float) expr()          {}
func (f Float) Tok() tok.Token { return f.X }

func (s String) expr()          {}
func (s String) Tok() tok.Token { return s.X }

func (t Time) expr()          {}
func (t Time) Tok() tok.Token { return t.X }

func (ue UnaryExpr) expr()          {}
func (ue UnaryExpr) Tok() tok.Token { return tok.Join(ue.Op, ue.X.Tok()) }

func (pe ParenExpr) expr()          {}
func (pe ParenExpr) Tok() tok.Token { return tok.Join(pe.LParen, pe.RParen) }

// leftOp is the operator on the left side of the expression.
func buildExpr(ctx *context, leftOp tok.Operator) (Expr, error) {
	var (
		err  error
		expr Expr
	)

	switch t := ctx.tok().(type) {
	case tok.Neg, tok.Sub, tok.Add:
		expr, err = buildUnaryExpr(ctx)
	case tok.Ident:
		switch ctx.nextTok().(type) {
		case tok.LParen:
			expr, err = buildCallExpr(ctx)
		default:
			expr, err = buildIdent(ctx)
		}
	case tok.QualIdent:
		switch ctx.nextTok().(type) {
		case tok.LParen:
			expr, err = buildCallExpr(ctx)
		default:
			expr, err = buildQualIdent(ctx)
		}
	case tok.Bool:
		expr, err = buildBool(ctx)
	case tok.Int:
		expr, err = buildInt(ctx)
	case tok.Float:
		expr, err = buildFloat(ctx)
	case tok.String:
		expr, err = buildString(ctx)
	case tok.Time:
		expr, err = buildTime(ctx)
	case tok.BitString:
		expr, err = buildBitString(ctx)
	case tok.LParen:
		expr, err = buildParenExpr(ctx)
	case tok.LBracket:
		expr, err = buildList(ctx)
	default:
		return Ident{}, unexpected(t, "expression")
	}

	if err != nil {
		return expr, err
	}

	for {
		var rightOp tok.Operator
		if op, ok := ctx.tok().(tok.Operator); ok {
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
	id := Ident{Name: ctx.tok()}
	ctx.idx++
	return id, nil
}

func buildQualIdent(ctx *context) (QualIdent, error) {
	id := QualIdent{Name: ctx.tok()}
	ctx.idx++
	return id, nil
}

func buildBool(ctx *context) (Bool, error) {
	b := Bool{ctx.tok().(tok.Bool)}
	ctx.idx++
	return b, nil
}

func buildInt(ctx *context) (Int, error) {
	int_ := Int{ctx.tok().(tok.Int)}
	ctx.idx++
	return int_, nil
}

func buildFloat(ctx *context) (Float, error) {
	r := Float{ctx.tok().(tok.Float)}
	ctx.idx++
	return r, nil
}

func buildString(ctx *context) (String, error) {
	s := String{ctx.tok().(tok.String)}
	ctx.idx++
	return s, nil
}

func buildTime(ctx *context) (Time, error) {
	t := Time{ctx.tok().(tok.Time)}
	ctx.idx++
	return t, nil
}

func buildBitString(ctx *context) (BitString, error) {
	s := BitString{ctx.tok().(tok.BitString)}
	ctx.idx++
	return s, nil
}

func buildParenExpr(ctx *context) (ParenExpr, error) {
	pe := ParenExpr{}
	var (
		err  error
		expr Expr
	)

	pe.LParen = ctx.tok().(tok.LParen)

	ctx.idx++
	expr, err = buildExpr(ctx, nil)
	if err != nil {
		return pe, err
	}
	pe.X = expr

	if rp, ok := ctx.tok().(tok.RParen); ok {
		pe.RParen = rp
		ctx.idx++
	} else {
		return pe, unexpected(ctx.tok(), "')'")
	}

	return pe, nil
}

func buildList(ctx *context) (List, error) {
	l := List{}
	prevExpr := false
	l.LBracket = ctx.tok().(tok.LBracket)
	lbi := ctx.idx // Left bracket token index
	ctx.idx++

tokenLoop:
	for {
		switch t := ctx.tok().(type) {
		case tok.RBracket:
			l.RBracket = t
			ctx.idx++
			break tokenLoop
		case tok.Comma:
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
	call := Call{Name: ctx.tok().(tok.Ident)}
	lpi := ctx.idx // Left parenthesis token index
	ctx.idx += 2

	prevExpr := false

tokenLoop:
	for {
		switch t := ctx.tok().(type) {
		case tok.RParen:
			ctx.idx++
			break tokenLoop
		case tok.Comma:
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
	op := ctx.tok().(tok.Operator)
	un := UnaryExpr{Op: op}
	ctx.idx++
	x, err := buildExpr(ctx, op)
	if err != nil {
		return un, err
	}
	un.X = x
	return un, nil
}
