package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func unexpected(t token.Token, expected string) error {
	return token.Error{
		Msg:  fmt.Sprintf("unexpected %s, expected "+expected, t.Name()),
		Toks: []token.Token{t},
	}
}
