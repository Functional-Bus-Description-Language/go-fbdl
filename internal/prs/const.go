package prs

import (
	"strings"
)

type Constant struct {
	base
	Value Expr
}

func (c Constant) GetSymbol(s string) (Symbol, error) {
	if strings.Contains(s, ".") {
		panic("To be implemented")
	}

	if c.parent != nil {
		return c.parent.GetSymbol(s)
	}

	return c.file.GetSymbol(s)
}
