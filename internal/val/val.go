// Package val provides types for Functional Bus Description Language type system.
package val

type Value interface {
	Type() string
}

// Bool represents FBDL bool type.
type Bool bool

func (b Bool) Type() string { return "bool" }

// Float represents FBDL float type.
type Float float64

func (f Float) Type() string { return "float" }

// Int represents FBDL integer type.
type Int int64

func (i Int) Type() string { return "integer" }

// List represents FBDL list type.
type List []Value

func (l List) Type() string { return "list" }

// Str represents FBDL string type.
type Str string

func (s Str) Type() string { return "string" }

// Time represents FBDL time type.
type Time struct {
	S  int64
	Ns int64
}

func (t Time) Type() string { return "time" }

func (t *Time) Normalize() {
	if t.Ns < 1000000000 {
		return
	}

	t.S += t.Ns / 1000000000
	t.Ns = t.Ns % 1000000000
}
