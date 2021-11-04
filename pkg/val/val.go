// Package val provides types for Functional Bus Description Language type system.
package val

type Value interface {
	Type() string
}

// Bool represents FBDL bool type.
type Bool bool

func (b Bool) Type() string {
	return "bool"
}

// Int represents FBDL integer type.
type Int int64

func (i Int) Type() string {
	return "integer"
}

// List represents FBDL list type.
// Internal value representation is a list of type implementing Value interface.
type List struct {
	Items []Value
}

func (l List) Type() string {
	return "list"
}

// Str represents FBDL string type.
type Str string

func (s Str) Type() string {
	return "string"
}
