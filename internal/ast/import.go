package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The import struct represents package import.
type Import struct {
	Name tok.Ident
	Path tok.String
}

func buildImport(toks []tok.Token, c *ctx) ([]Import, error) {
	switch t := toks[c.i+1].(type) {
	case tok.Ident, tok.String:
		return buildSingleImport(toks, c)
	case tok.Newline:
		return buildMultiImport(toks, c)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []tok.Token, c *ctx) ([]Import, error) {
	i := Import{}

	c.i++
	switch t := toks[c.i].(type) {
	case tok.Ident:
		i.Name = t
		c.i++
		switch t := toks[c.i].(type) {
		case tok.String:
			i.Path = t
			c.i++
		default:
			return nil, unexpected(t, "string")
		}
	case tok.String:
		i.Path = t
		c.i++
	}

	return []Import{i}, nil
}

func buildMultiImport(toks []tok.Token, c *ctx) ([]Import, error) {
	imps := []Import{}
	i := Import{}

	c.i += 2
	if _, ok := toks[c.i].(tok.Indent); !ok {
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
			case tok.Ident:
				i.Name = t
				state = Path
			case tok.String:
				i.Path = t
				imps = append(imps, i)
				i = Import{}
			case tok.Newline:
				// Do nothing
			case tok.Dedent:
				c.i++
				break tokenLoop
			case tok.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or string")
			}
		case Path:
			switch t := toks[c.i].(type) {
			case tok.String:
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
