package prs

import (
	"strings"
)

// Const represents constant definition.
type Const struct {
	base
	Value Expr
}

func (c Const) GetSymbol(s string) (Symbol, error) {
	if strings.Contains(s, ".") {
		panic("To be implemented")
	}

	if c.parent != nil {
		return c.parent.GetSymbol(s)
	}

	return c.file.GetSymbol(s)
}
