package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Import represents a package import.
type Import struct {
	Name token.Token // token.Ident or nil
	Path token.String
}

func buildImport(tokens *tokenStream) ([]Import, error) {
	switch t := tokens.peek(1).(type) {
	case token.Ident, token.String:
		return buildSingleImport(tokens)
	case token.Newline:
		return buildMultiImport(tokens)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(tokens *tokenStream) ([]Import, error) {
	i := Import{}

	tokens.next()
	switch t := tokens.peek(0).(type) {
	case token.Ident:
		i.Name = t
		tokens.next()
		switch t := tokens.peek(0).(type) {
		case token.String:
			i.Path = t
			tokens.next()
		default:
			return nil, unexpected(t, "string")
		}
	case token.String:
		i.Path = t
		tokens.next()
	}

	return []Import{i}, nil
}

func buildMultiImport(tokens *tokenStream) ([]Import, error) {
	tokens.next()
	tokens.next()
	if _, ok := tokens.peek(0).(token.Indent); !ok {
		return nil, unexpected(tokens.peek(0), "indent increase")
	}

	var imps []Import

	for {
		tokens.next()

		switch t := tokens.peek(0).(type) {
		case token.Newline:
			// Go to next line
		case token.Dedent:
			tokens.next()
			return imps, nil
		case token.Eof:
			return imps, nil
		case token.String:
			imps = append(imps, Import{Path: t})
		case token.Ident:
			tokens.next()
			path, ok := tokens.peek(0).(token.String)
			if !ok {
				return nil, unexpected(tokens.peek(0), "string")
			}
			imps = append(imps, Import{Name: t, Path: path})
		default:
			return nil, unexpected(t, "identifier or string")
		}
	}
}
