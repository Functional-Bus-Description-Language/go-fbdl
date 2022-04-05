// Package val provides types for Functional Bus Description Language type system.
package val

import (
	"encoding/binary"
)

type Value interface {
	Type() string
	Bytes() []byte
}

// Bool represents FBDL bool type.
type Bool bool

func (b Bool) Type() string {
	return "bool"
}

func (b Bool) Bytes() []byte {
	if b {
		return []byte{1}
	}
	return []byte{0}
}

// Int represents FBDL integer type.
type Int int64

func (i Int) Type() string {
	return "integer"
}

func (i Int) Bytes() []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return b
}

// List represents FBDL list type.
type List []Value

func (l List) Type() string {
	return "list"
}

func (l List) Bytes() []byte {
	b := []byte{}
	for _, v := range l {
		b = append(b, v.Bytes()...)
	}
	return b
}

// Str represents FBDL string type.
type Str string

func (s Str) Type() string {
	return "string"
}

func (s Str) Bytes() []byte {
	return []byte(s)
}
