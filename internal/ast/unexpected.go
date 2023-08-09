package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
)

func unexpected(t tok.Token, expected string) error {
	return fmt.Errorf(
		"%s: unexpected %s, expected "+expected,
		tok.Loc(t), t.Kind(),
	)
}
