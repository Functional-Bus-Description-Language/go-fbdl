package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// Building context
type context struct {
	idx  int // Current token index
	toks []token.Token
}

func (ctx context) token() token.Token   { return ctx.toks[ctx.idx] }
func (ctx context) nextTok() token.Token { return ctx.toks[ctx.idx+1] }
