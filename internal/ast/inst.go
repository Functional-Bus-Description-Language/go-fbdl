package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Inst represents functionality instantiation.
type Inst struct {
	Doc     Doc
	Name    token.Ident
	Count   Expr        // If not nil, then it is a list
	Type    token.Token // Basic type, identifier or qualified identifier
	ArgList ArgList
	Body    Body
}

func buildInst(tokens *tokenStream) (Inst, error) {
	inst := Inst{Name: tokens.peek(0).(token.Ident)}
	tokens.next()

	// Count
	if _, ok := tokens.peek(0).(token.LBracket); ok {
		tokens.next()
		expr, err := buildExpr(tokens, nil)
		if err != nil {
			return inst, err
		}
		inst.Count = expr
		if _, ok := tokens.peek(0).(token.RBracket); !ok {
			return inst, unexpected(tokens.peek(0), "']'")
		}
		tokens.next()
	}

	// Type
	switch t := tokens.peek(0).(type) {
	case token.Functionality, token.Ident, token.QualIdent:
		inst.Type = t
		tokens.next()
	default:
		return inst, unexpected(t, "functionality type")
	}

	// Argument List
	argList, err := buildArgList(tokens)
	if err != nil {
		return inst, err
	}
	inst.ArgList = argList

	// Body
	switch t := tokens.peek(0).(type) {
	case token.Semicolon:
		tokens.next()
		props, err := buildPropAssignments(tokens)
		if err != nil {
			return inst, err
		}
		inst.Body.Props = props
	case token.Newline:
		if _, ok := tokens.peek(1).(token.Indent); ok {
			tokens.next()
			tokens.next()
			body, err := buildBody(tokens)
			if err != nil {
				return inst, err
			}
			inst.Body = body
		}
	case token.Eof:
		break
	default:
		return inst, unexpected(t, "';' or newline")
	}

	return inst, nil
}
