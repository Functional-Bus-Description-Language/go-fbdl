package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Body represents a functionality body.
type Body struct {
	Consts []Const
	Insts  []Inst
	Props  []Prop
	Types  []Type
}

func buildBody(toks []tok.Token, ctx *context) (Body, error) {
	var (
		err    error
		body   Body
		consts []Const
		doc    Doc
		ins    Inst
		props  []Prop
		typ    Type
	)

tokenLoop:
	for {
		if _, ok := toks[ctx.idx].(tok.Eof); ok {
			break
		}

		switch t := toks[ctx.idx].(type) {
		case tok.Newline:
			ctx.idx++
		case tok.Comment:
			doc = buildDoc(toks, ctx)
		case tok.Const:
			consts, err = buildConst(toks, ctx)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				body.Consts = append(body.Consts, consts...)
			}
		case tok.Ident:
			ins, err = buildInst(toks, ctx)
			body.Insts = append(body.Insts, ins)
		case tok.Property:
			props, err = buildPropAssignments(toks, ctx)
			if err != nil {
				return body, err
			}
			if props != nil {
				body.Props = append(body.Props, props...)
			}
		case tok.Type:
			typ, err = buildType(toks, ctx)
			body.Types = append(body.Types, typ)
		case tok.Dedent:
			ctx.idx++
			break tokenLoop
		default:
			return body, unexpected(t, "const, type, identifier, or comment")
		}

		if err != nil {
			return body, err
		}
	}

	return body, nil
}
