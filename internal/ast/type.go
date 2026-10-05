package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Type represents type definition.
type Type struct {
	Doc    Doc
	Name   token.Ident
	Params []Param
	Count  Expr        // If Count is not nil, then the type is a list
	Type   token.Token // Basic type, identifier or qualified identifier
	Args   ArgList
	Body   Body
}

func buildType(tokens *tokenStream) (Type, error) {
	typ := Type{}
	tokens.next()

	// Name
	if t, ok := tokens.peek(0).(token.Ident); ok {
		typ.Name = t
	} else {
		return typ, unexpected(tokens.peek(0), "identifier")
	}
	tokens.next()

	// Parameter List
	params, err := buildParamList(tokens)
	if err != nil {
		return typ, err
	}
	typ.Params = params

	// Count
	if _, ok := tokens.peek(0).(token.LBracket); ok {
		tokens.next()
		expr, err := buildExpr(tokens, nil)
		if err != nil {
			return typ, err
		}
		typ.Count = expr
		if _, ok := tokens.peek(0).(token.RBracket); !ok {
			return typ, unexpected(tokens.peek(0), "']'")
		}
		tokens.next()
	}

	// Type
	switch t := tokens.peek(0).(type) {
	case token.Functionality, token.Ident, token.QualIdent:
		typ.Type = t
		tokens.next()
	default:
		return typ, unexpected(t, "functionality type")
	}

	// Argument List
	args, err := buildArgList(tokens)
	if err != nil {
		return typ, err
	}
	typ.Args = args

	// Body
	switch t := tokens.peek(0).(type) {
	case token.Semicolon:
		tokens.next()
		props, err := buildPropAssignments(tokens)
		if err != nil {
			return typ, err
		}
		typ.Body.Props = props
	case token.Newline:
		if _, ok := tokens.peek(1).(token.Indent); ok {
			tokens.next()
			tokens.next()
			body, err := buildBody(tokens)
			if err != nil {
				return typ, err
			}
			typ.Body = body
		}
	case token.Eof:
		// Do nothing.
	default:
		return typ, unexpected(t, "';' or newline")
	}

	return typ, nil
}
