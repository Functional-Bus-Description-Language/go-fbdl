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

func buildConst(tokens *tokenStream) ([]Const, error) {
	tokens.next()

	switch t := tokens.peek(0).(type) {
	case token.Ident, token.Functionality, token.WordProperty:
		return buildSingleConst(tokens)
	case token.Newline:
		return buildMultiConst(tokens)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(tokens *tokenStream) ([]Const, error) {
	var name token.Ident
	switch tok := tokens.next().(type) {
	case token.Ident:
		name = tok
	case token.Functionality:
		name = tok.Ident()
	case token.WordProperty:
		name = tok.Ident()
	default:
		return nil, unexpected(tok, "identifier")
	}

	if _, ok := tokens.peek(0).(token.Ass); !ok {
		return nil, unexpected(tokens.peek(0), "'='")
	}
	tokens.next()

	expr, err := buildExpr(tokens, nil)
	if err != nil {
		return nil, err
	}

	return []Const{{Name: name, Value: expr}}, nil
}

func buildMultiConst(tokens *tokenStream) ([]Const, error) {
	for {
		if _, ok := tokens.peek(0).(token.Newline); !ok {
			break
		}
		tokens.next()
	}
	if _, ok := tokens.peek(0).(token.Indent); !ok {
		return nil, unexpected(tokens.peek(0), "indent or newline")
	}
	tokens.next()

	consts := []Const{}
	doc := emptyDoc()

	for {
		switch t := tokens.peek(0).(type) {
		case token.Newline:
			tokens.next()
			doc = emptyDoc()
		case token.Comment:
			doc = buildDoc(tokens)
		case token.Ident, token.Functionality, token.WordProperty:
			con, err := buildSingleConst(tokens)
			if err != nil {
				return nil, err
			}
			con[0].Doc = doc
			consts = append(consts, con...)
			doc = emptyDoc()
		case token.Dedent:
			if len(consts) == 0 {
				return nil, unexpected(t, "identifier")
			}
			tokens.next()
			return consts, nil
		case token.Eof:
			if len(consts) == 0 {
				return nil, unexpected(t, "identifier")
			}
			return consts, nil
		default:
			if len(consts) == 0 {
				return nil, unexpected(t, "identifier")
			}
			return nil, unexpected(t, "identifier or dedent")
		}
	}
}
