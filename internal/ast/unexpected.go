package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

func unexpected(t tok.Token, expected string) error {
	return tok.Error{
		Tok: t,
		Msg: fmt.Sprintf("unexpected %s, expected "+expected, t.Name()),
	}
}
