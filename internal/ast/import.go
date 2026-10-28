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
	case token.Ident, token.Functionality, token.WordProperty, token.String:
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
	tok := tokens.peek(0)
	switch t := tok.(type) {
	case token.Functionality:
		tok = t.Ident()
	case token.WordProperty:
		tok = t.Ident()
	}
	switch t := tok.(type) {
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

	switch t := tokens.peek(0).(type) {
	case token.Newline, token.Eof:
	default:
		return nil, unexpected(t, "newline or end of file")
	}

	return []Import{i}, nil
}

func buildMultiImport(tokens *tokenStream) ([]Import, error) {
	importTok := tokens.next().(token.Import)
	emptyErr := token.Error{
		Msg:  "empty import list",
		Toks: []token.Token{importTok},
	}
	tokens.skipNewlines()
	if _, ok := tokens.peek(0).(token.Eof); ok {
		return nil, emptyErr
	}
	if _, ok := tokens.peek(0).(token.Indent); !ok {
		return nil, unexpected(tokens.peek(0), "indent increase")
	}

	var imps []Import

	for {
		tokens.next()

		tok := tokens.peek(0)
		switch t := tok.(type) {
		case token.Functionality:
			tok = t.Ident()
		case token.WordProperty:
			tok = t.Ident()
		}
		switch t := tok.(type) {
		case token.Newline, token.Comment:
			continue
		case token.Dedent, token.Eof:
			if len(imps) == 0 {
				return nil, emptyErr
			}
			if _, ok := t.(token.Dedent); ok {
				tokens.next()
			}
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

		switch t := tokens.peek(1).(type) {
		case token.Newline, token.Eof:
		default:
			return nil, unexpected(t, "newline or end of file")
		}
	}
}
