package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// The import struct represents package import.
type Import struct {
	Name tok.Token // tok.Ident or nil
	Path tok.String
}

func buildImport(toks []tok.Token, ctx *context) ([]Import, error) {
	switch t := toks[ctx.i+1].(type) {
	case tok.Ident, tok.String:
		return buildSingleImport(toks, ctx)
	case tok.Newline:
		return buildMultiImport(toks, ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(toks []tok.Token, ctx *context) ([]Import, error) {
	i := Import{}

	ctx.i++
	switch t := toks[ctx.i].(type) {
	case tok.Ident:
		i.Name = t
		ctx.i++
		switch t := toks[ctx.i].(type) {
		case tok.String:
			i.Path = t
			ctx.i++
		default:
			return nil, unexpected(t, "string")
		}
	case tok.String:
		i.Path = t
		ctx.i++
	}

	return []Import{i}, nil
}

func buildMultiImport(toks []tok.Token, ctx *context) ([]Import, error) {
	imps := []Import{}
	i := Import{}

	ctx.i += 2
	if _, ok := toks[ctx.i].(tok.Indent); !ok {
		return nil, unexpected(toks[ctx.i], "indent increase")
	}

	const (
		Name = iota
		Path
	)
	state := Name

tokenLoop:
	for {
		ctx.i++
		switch state {
		case Name:
			switch t := toks[ctx.i].(type) {
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
				ctx.i++
				break tokenLoop
			case tok.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or string")
			}
		case Path:
			switch t := toks[ctx.i].(type) {
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
