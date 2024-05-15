package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Const represents a constant.
type Const struct {
	Doc   Doc
	Name  tok.Ident
	Value Expr
}

func buildConst(ctx *context) ([]Const, error) {
	switch t := ctx.nextTok().(type) {
	case tok.Ident:
		return buildSingleConst(ctx)
	case tok.Newline:
		return buildMultiConst(ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(ctx *context) ([]Const, error) {
	con := Const{Name: ctx.nextTok().(tok.Ident)}

	ctx.idx += 2
	if _, ok := ctx.tok().(tok.Ass); !ok {
		return nil, unexpected(ctx.tok(), "'='")
	}

	ctx.idx++
	expr, err := buildExpr(ctx, nil)
	if err != nil {
		return nil, err
	}
	con.Value = expr

	return []Const{con}, nil
}

func buildMultiConst(ctx *context) ([]Const, error) {
	consts := []Const{}
	con := Const{}

	type State int
	const (
		Indent State = iota
		FirstId
		Ass
		Exp
		Id
	)
	state := Indent

	ctx.idx += 1
tokenLoop:
	for {
		ctx.idx++
		switch state {
		case Indent:
			switch t := ctx.tok().(type) {
			case tok.Newline:
				continue
			case tok.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := ctx.tok().(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := ctx.tok().(type) {
			case tok.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(ctx, nil)
			if err != nil {
				return nil, err
			}
			con.Value = expr
			consts = append(consts, con)
			con = Const{}
			ctx.idx--
			state = Id
		case Id:
			switch t := ctx.tok().(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			case tok.Comment:
				doc := buildDoc(ctx)
				con.Doc = doc
				ctx.idx--
			case tok.Newline:
				continue
			case tok.Dedent:
				ctx.idx++
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
