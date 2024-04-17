package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
	"reflect"
)

// The const struct represents constant.
type Const struct {
	Doc   Doc
	Name  tok.Ident
	Value Expr
}

func (c Const) eq(c2 Const) bool {
	return reflect.DeepEqual(c, c2)
}

func buildConst(toks []tok.Token, ctx *context) ([]Const, error) {
	switch t := toks[ctx.i+1].(type) {
	case tok.Ident:
		return buildSingleConst(toks, ctx)
	case tok.Newline:
		return buildMultiConst(toks, ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(toks []tok.Token, ctx *context) ([]Const, error) {
	con := Const{Name: toks[ctx.i+1].(tok.Ident)}

	ctx.i += 2
	if _, ok := toks[ctx.i].(tok.Ass); !ok {
		return nil, unexpected(toks[ctx.i], "'='")
	}

	ctx.i++
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

	const (
		Indent int = iota
		FirstId
		Ass
		Exp
		Id
	)
	state := Indent

	ctx.i += 1
tokenLoop:
	for {
		ctx.i++
		switch state {
		case Indent:
			switch t := toks[ctx.i].(type) {
			case tok.Newline:
				continue
			case tok.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := toks[ctx.i].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := toks[ctx.i].(type) {
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
			ctx.i--
			state = Id
		case Id:
			switch t := toks[ctx.i].(type) {
			case tok.Ident:
				con.Name = t
				state = Ass
			case tok.Comment:
				doc := buildDoc(toks, ctx)
				con.Doc = doc
				ctx.i--
			case tok.Newline:
				continue
			case tok.Dedent:
				ctx.i++
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
