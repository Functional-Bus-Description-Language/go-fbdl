package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The import struct represents package import.
type Import struct {
	Name token.Ident
	Path token.String
}

func buildImport(toks []token.Token, c *ctx) ([]Import, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(toks, c)
	case token.Newline:
		return buildMultiImport(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []token.Token, c *ctx) ([]Import, error) {
	i := Import{}

	c.i++
	switch t := toks[c.i].(type) {
	case token.Ident:
		i.Name = t
		c.i++
		switch t := toks[c.i].(type) {
		case token.String:
			i.Path = t
			c.i++
		default:
			return nil, unexpected(t, "string")
		}
	case token.String:
		i.Path = t
		c.i++
	}

	return []Import{i}, nil
}

func buildMultiImport(toks []token.Token, c *ctx) ([]Import, error) {
	imps := []Import{}
	i := Import{}

	c.i += 2
	if _, ok := toks[c.i].(token.Indent); !ok {
		return nil, unexpected(toks[c.i], "indent increase")
	}

	const (
		Name = iota
		Path
	)
	state := Name

tokenLoop:
	for {
		c.i++
		switch state {
		case Name:
			switch t := toks[c.i].(type) {
			case token.Ident:
				i.Name = t
				state = Path
			case token.String:
				i.Path = t
				imps = append(imps, i)
				i = Import{}
			case token.Newline:
				// Do nothing
			case token.Dedent:
				c.i++
				break tokenLoop
			case token.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or string")
			}
		case Path:
			switch t := toks[c.i].(type) {
			case token.String:
				i.Path = t
				imps = append(imps, i)
				i = Import{}
				state = Name
			default:
				return nil, unexpected(t, "string")
			}
		}
	}

	return imps, nil
}
