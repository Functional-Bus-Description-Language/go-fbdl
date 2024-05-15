package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

// Building context
type context struct {
	idx  int // Current token index
	toks []tok.Token
}

func (ctx context) tok() tok.Token     { return ctx.toks[ctx.idx] }
func (ctx context) nextTok() tok.Token { return ctx.toks[ctx.idx+1] }
