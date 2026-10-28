package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

type tokenStream struct {
	idx     int // Next token index in toks
	toks    []token.Token
	pending []token.Token // Expanded tokens consumed before toks[idx:]
}

func (ts *tokenStream) addPending(toks ...token.Token) {
	ts.pending = append(ts.pending, toks...)
}

func (ts *tokenStream) next() token.Token {
	tok := ts.peek(0)
	if len(ts.pending) > 0 {
		ts.pending = ts.pending[1:]
	} else {
		ts.idx++
	}
	return tok
}

func (ts *tokenStream) skipNewlines() {
	for {
		if _, ok := ts.peek(0).(token.Newline); !ok {
			return
		}
		ts.next()
	}
}

func (ts *tokenStream) peek(offset int) token.Token {
	if offset < len(ts.pending) {
		return ts.pending[offset]
	}
	offset -= len(ts.pending)
	if ts.idx+offset < len(ts.toks) {
		return ts.toks[ts.idx+offset]
	}
	return ts.toks[len(ts.toks)-1]
}
