package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Body represents a functionality body.
type Body struct {
	Consts []Const
	Insts  []Inst
	Props  []Prop
	Types  []Type
}

func buildBody(ctx *context) (Body, error) {
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
		if _, ok := ctx.token().(token.Eof); ok {
			break
		}

		switch tok := ctx.token().(type) {
		case token.Newline:
			ctx.idx++
		case token.Comment:
			doc = buildDoc(ctx)
		case token.Const:
			consts, err = buildConst(ctx)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				body.Consts = append(body.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(ctx)
			body.Insts = append(body.Insts, ins)
		case token.Property:
			props, err = buildPropAssignments(ctx)
			if err != nil {
				return body, err
			}
			if props != nil {
				body.Props = append(body.Props, props...)
			}
		case token.Type:
			typ, err = buildType(ctx)
			body.Types = append(body.Types, typ)
		case token.Dedent:
			ctx.idx++
			break tokenLoop
		default:
			return body, unexpected(tok, "const, type, identifier, or comment")
		}

		if err != nil {
			return body, err
		}
	}

	return body, nil
}
