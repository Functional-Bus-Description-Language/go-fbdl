package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type Import interface {
	importNode()
}

// Import types
type (
	SingleImport struct {
		Name token.Ident
		Path token.String
	}
)

func (si SingleImport) importNode() {}

func buildImport(toks []token.Token, c *ctx) (Import, error) {
	switch t := toks[c.i+1].(type) {
	case token.Ident, token.String:
		return buildSingleImport(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []token.Token, c *ctx) (SingleImport, error) {
	si := SingleImport{}

	c.i++
	switch t := toks[c.i].(type) {
	case token.Ident:
		si.Name = t
		c.i++
		switch t := toks[c.i].(type) {
		case token.String:
			si.Path = t
			c.i++
		default:
			return si, unexpected(t, "string")
		}
	case token.String:
		si.Path = t
		c.i++
	}

	return si, nil
}
