package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Import represents a package import.
type Import struct {
	Name token.Token // token.Ident or nil
	Path token.String
}

func buildImport(ctx *context) ([]Import, error) {
	switch t := ctx.nextTok().(type) {
	case token.Ident, token.String:
		return buildSingleImport(ctx)
	case token.Newline:
		return buildMultiImport(ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(ctx *context) ([]Import, error) {
	i := Import{}

	ctx.idx++
	switch t := ctx.token().(type) {
	case token.Ident:
		i.Name = t
		ctx.idx++
		switch t := ctx.token().(type) {
		case token.String:
			i.Path = t
			ctx.idx++
		default:
			return nil, unexpected(t, "string")
		}
	case token.String:
		i.Path = t
		ctx.idx++
	}

	return []Import{i}, nil
}

func buildMultiImport(ctx *context) ([]Import, error) {
	imps := []Import{}
	i := Import{}

	ctx.idx += 2
	if _, ok := ctx.token().(token.Indent); !ok {
		return nil, unexpected(ctx.token(), "indent increase")
	}

	type State int
	const (
		Name State = iota
		Path
	)
	state := Name

tokenLoop:
	for {
		ctx.idx++
		switch state {
		case Name:
			switch t := ctx.token().(type) {
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
				ctx.idx++
				break tokenLoop
			case token.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or string")
			}
		case Path:
			switch t := ctx.token().(type) {
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
