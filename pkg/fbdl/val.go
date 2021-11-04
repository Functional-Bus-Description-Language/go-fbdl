package fbdl

type Value interface {
	Type() string
}

// Bool represents FBDL bool type.
// Internal value representation is bool.
type Bool struct {
	V bool
}

func (b Bool) Type() string {
	return "bool"
}

// Int represents FBDL integer type.
// Internal value representation is int64.
type Int struct {
	V int64
}

func (i Int) Type() string {
	return "integer"
}

// List represents FBDL list type.
// Internal value representation is a list of type implementing Value interface.
type List struct {
	V []Value
}

func (l List) Type() string {
	return "list"
}

// Str represents FBDL string type.
// Internal value representation is string.
type Str struct {
	V string
}

func (s Str) Type() string {
	return "string"
}
