package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The Body struct represents functionality body.
type Body struct {
	Consts []Const
	Insts  []Inst
	Props  []Prop
	Types  []Type
}

func (b Body) eq(b2 Body) bool {
	if len(b.Consts) != len(b2.Consts) ||
		len(b.Insts) != len(b2.Insts) ||
		len(b.Props) != len(b2.Props) {
		return false
	}

	for i := range b.Consts {
		if !b.Consts[i].eq(b2.Consts[i]) {
			return false
		}
	}
	for i := range b.Insts {
		if !b.Insts[i].eq(b2.Insts[i]) {
			return false
		}
	}
	for i := range b.Props {
		if b.Props[i] != b2.Props[i] {
			return false
		}
	}
	for i := range b.Types {
		if !b.Types[i].eq(b2.Types[i]) {
			return false
		}
	}

	return true
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
		if _, ok := toks[ctx.i].(tok.Eof); ok {
			break
		}

		switch t := toks[ctx.i].(type) {
		case tok.Newline:
			ctx.i++
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
			ctx.i++
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
