package value

type Value interface {
	IsValue() bool
}

type Bool struct {
	V bool
}

func (b Bool) IsValue() bool {
	return true
}

type Integer struct {
	V int32
}

func (i Integer) IsValue() bool {
	return true
}

type String struct {
	v string
}

func (s String) IsValue() bool {
	return true
}

