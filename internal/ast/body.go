package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Body struct represents functionality body.
type Body struct {
	Consts []Const
	Insts  []Inst
	Props  []Prop
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

	return true
}

func buildBody(toks []token.Token, c *ctx) (Body, error) {
	var (
		err    error
		body   Body
		consts []Const
		doc    Doc
		ins    Inst
		props  []Prop
	)

	for {
		if _, ok := toks[c.i].(token.Eof); ok {
			break
		}

		switch t := toks[c.i].(type) {
		case token.Newline:
			c.i++
		case token.Comment:
			doc = buildDoc(toks, c)
		case token.Const:
			consts, err = buildConst(toks, c)
			if len(consts) > 0 {
				if doc.endLine() == consts[0].Name.Line()+1 {
					consts[0].Doc = doc
				}
				body.Consts = append(body.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(toks, c)
			body.Insts = append(body.Insts, ins)
		case token.Property:
			props, err = buildPropAssignments(toks, c)
			if err != nil {
				return body, err
			}
			if props != nil {
				body.Props = append(body.Props, props...)
			}
		default:
			panic(fmt.Sprintf("%s: unhandled token %s", token.Loc(t), t.Kind()))
		}

		if err != nil {
			return body, err
		}
	}

	return body, nil
}
