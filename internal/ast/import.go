package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Import represents a package import.
type Import struct {
	Name tok.Token // tok.Ident or nil
	Path tok.String
}

func buildImport(ctx *context) ([]Import, error) {
	switch t := ctx.nextTok().(type) {
	case tok.Ident, tok.String:
		return buildSingleImport(ctx)
	case tok.Newline:
		return buildMultiImport(ctx)
	default:
		return nil, unexpected(t, "identifier, string or newline")
	}
}

func buildSingleImport(ctx *context) ([]Import, error) {
	i := Import{}

	ctx.idx++
	switch t := ctx.tok().(type) {
	case tok.Ident:
		i.Name = t
		ctx.idx++
		switch t := ctx.tok().(type) {
		case tok.String:
			i.Path = t
			ctx.idx++
		default:
			return nil, unexpected(t, "string")
		}
	case tok.String:
		i.Path = t
		ctx.idx++
	}

	return []Import{i}, nil
}

func buildMultiImport(ctx *context) ([]Import, error) {
	imps := []Import{}
	i := Import{}

	ctx.idx += 2
	if _, ok := ctx.tok().(tok.Indent); !ok {
		return nil, unexpected(ctx.tok(), "indent increase")
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
			switch t := ctx.tok().(type) {
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
				ctx.idx++
				break tokenLoop
			case tok.Eof:
				break tokenLoop
			default:
				return nil, unexpected(t, "identifier or string")
			}
		case Path:
			switch t := ctx.tok().(type) {
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
