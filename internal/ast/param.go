package ast

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/token"
)

// The Param struct represents type parameter.
type Param struct {
	Name  token.Ident
	Value Expr // Default value of the parameter
}
