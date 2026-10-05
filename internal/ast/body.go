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

func buildBody(tokens *tokenStream) (Body, error) {
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
		if _, ok := tokens.peek(0).(token.Eof); ok {
			break
		}

		switch tok := tokens.peek(0).(type) {
		case token.Newline:
			doc = emptyDoc()
			tokens.next()
		case token.Comment:
			doc = buildDoc(tokens)
		case token.Const:
			consts, err = buildConst(tokens)
			if len(consts) > 0 {
				if !doc.isEmpty() {
					consts[0].Doc = doc
				}
				body.Consts = append(body.Consts, consts...)
			}
		case token.Ident:
			ins, err = buildInst(tokens)
			body.Insts = append(body.Insts, ins)
		case token.Property:
			props, err = buildPropAssignments(tokens)
			if err != nil {
				return body, err
			}
			if props != nil {
				body.Props = append(body.Props, props...)
			}
		case token.Type:
			typ, err = buildType(tokens)
			body.Types = append(body.Types, typ)
		case token.Dedent:
			tokens.next()
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
