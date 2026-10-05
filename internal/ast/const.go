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
	case token.Ident:
		return buildSingleConst(tokens)
	case token.Newline:
		return buildMultiConst(tokens)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleConst(tokens *tokenStream) ([]Const, error) {
	con := Const{Name: tokens.next().(token.Ident)}

	if _, ok := tokens.peek(0).(token.Ass); !ok {
		return nil, unexpected(tokens.peek(0), "'='")
	}
	tokens.next()

	expr, err := buildExpr(tokens, nil)
	if err != nil {
		return nil, err
	}
	con.Value = expr

	return []Const{con}, nil
}

func buildMultiConst(tokens *tokenStream) ([]Const, error) {
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

tokenLoop:
	for {
		switch state {
		case Indent:
			switch t := tokens.next().(type) {
			case token.Newline:
				continue
			case token.Indent:
				state = FirstId
			default:
				return nil, unexpected(t, "indent or newline")
			}
		case FirstId:
			switch t := tokens.peek(0).(type) {
			case token.Ident:
				con.Name = t
				tokens.next()
				state = Ass
			case token.Comment:
				con.Doc = buildDoc(tokens)
			case token.Newline:
				tokens.next()
				con.Doc = emptyDoc()
			default:
				return nil, unexpected(t, "identifier")
			}
		case Ass:
			switch t := tokens.next().(type) {
			case token.Ass:
				state = Exp
			default:
				return nil, unexpected(t, "'='")
			}
		case Exp:
			expr, err := buildExpr(tokens, nil)
			if err != nil {
				return nil, err
			}
			con.Value = expr
			consts = append(consts, con)
			con = Const{}
			state = Id
		case Id:
			switch t := tokens.peek(0).(type) {
			case token.Ident:
				con.Name = t
				tokens.next()
				state = Ass
			case token.Comment:
				doc := buildDoc(tokens)
				con.Doc = doc
			case token.Newline:
				tokens.next()
				con.Doc = emptyDoc()
				continue
			case token.Dedent:
				tokens.next()
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
