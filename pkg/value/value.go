package value

type Value interface {
	Type() string
}

type Bool struct {
	V bool
}

func (b Bool) Type() string {
	return "bool"
}

type Integer struct {
	V int32
}

func (i Integer) Type() string {
	return "integer"
}

type String struct {
	v string
}

func (s String) Type() string {
	return "string"
}
