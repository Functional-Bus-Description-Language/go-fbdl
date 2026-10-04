package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Const represents a constant.
type Const struct {
	Doc   Doc
	Name  token.Ident
	Value Expr
}

func buildConst(ctx *context) ([]Const, error) {
	switch t := ctx.nextTok().(type) {
	case token.Ident:
		return buildSingleConst(ctx)
	case token.Newline:
		return buildMultiConst(ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(ctx *context) ([]Const, error) {
	con := Const{Name: ctx.nextTok().(token.Ident)}

	ctx.idx += 2
	if _, ok := ctx.token().(token.Ass); !ok {
		return nil, unexpected(ctx.token(), "'='")
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
			switch t := ctx.token().(type) {
			case token.Newline:
				continue
			case token.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := ctx.token().(type) {
			case token.Ident:
				con.Name = t
				state = Ass
			case token.Comment:
				con.Doc = buildDoc(ctx)
				ctx.idx--
			case token.Newline:
				con.Doc = emptyDoc()
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := ctx.token().(type) {
			case token.Ass:
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
			switch t := ctx.token().(type) {
			case token.Ident:
				con.Name = t
				state = Ass
			case token.Comment:
				doc := buildDoc(ctx)
				con.Doc = doc
				ctx.idx--
			case token.Newline:
				con.Doc = emptyDoc()
				continue
			case token.Dedent:
				ctx.idx++
				break tokenLoop
			case token.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or dedent")
			}
		}
	}

	return consts, nil
}
