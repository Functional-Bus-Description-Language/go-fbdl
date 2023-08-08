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
