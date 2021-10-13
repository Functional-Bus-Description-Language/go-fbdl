package parse

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/expr"
	"strings"
)

type Constant struct {
	base
	value expr.Expression
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
