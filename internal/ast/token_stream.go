package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type tokenStream struct {
	idx  int // Current token index
	toks []token.Token
}

func (ts *tokenStream) next() token.Token {
	tok := ts.peek(0)
	ts.idx++
	return tok
}

func (ts *tokenStream) peek(offset int) token.Token {
	if ts.idx+offset < len(ts.toks) {
		return ts.toks[ts.idx+offset]
	}
	return ts.toks[len(ts.toks)-1]
}
