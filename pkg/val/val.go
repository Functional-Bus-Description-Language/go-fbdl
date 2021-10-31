package val

type Value interface {
	Type() string
}

type Bool struct {
	V bool
}

func (b Bool) Type() string {
	return "bool"
}

type Int struct {
	V int32
}

func (i Int) Type() string {
	return "integer"
}

type List struct {
	V []Value
}

func (l List) Type() string {
	return "list"
}

type Str struct {
	V string
}

func (s Str) Type() string {
	return "string"
}
