package ast

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

func unexpected(t token.Token, expected string) error {
	return fmt.Errorf(
		"%s: unexpected %s, expected "+expected,
		token.Loc(t), t.Kind(),
	)
}
