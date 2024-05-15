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

func buildConst(toks []tok.Token, ctx *context) ([]Const, error) {
	switch t := toks[ctx.idx+1].(type) {
	case tok.Ident:
		return buildSingleConst(toks, ctx)
	case tok.Newline:
		return buildMultiConst(toks, ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []tok.Token, ctx *context) ([]Const, error) {
	con := Const{Name: toks[ctx.idx+1].(tok.Ident)}

	ctx.idx += 2
	if _, ok := toks[ctx.idx].(tok.Ass); !ok {
		return nil, unexpected(toks[ctx.idx], "'='")
	}

	ctx.idx++
	expr, err := buildExpr(toks, ctx, nil)
	if err != nil {
		return nil, err
	}
	con.Value = expr

	return []Const{con}, nil
}

func buildMultiConst(toks []tok.Token, ctx *context) ([]Const, error) {
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
			switch t := toks[ctx.idx].(type) {
			case tok.Newline:
				continue
			case tok.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[ctx.idx].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[ctx.idx].(type) {
			case tok.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(toks, ctx, nil)
			if err != nil {
				return nil, err
			}
			con.Value = expr
			consts = append(consts, con)
			con = Const{}
			ctx.idx--
			state = Id
		case Id:
			switch t := toks[ctx.idx].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			case tok.Comment:
				doc := buildDoc(toks, ctx)
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
