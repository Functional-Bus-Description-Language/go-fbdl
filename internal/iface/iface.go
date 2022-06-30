package iface

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Element interface {
	Type() string
	Name() string
	Doc() string
	IsArray() bool
	Count() int64
	Hash() int64
}

type ConstContainer interface {
	BoolConsts() map[string]bool
	BoolListConsts() map[string][]bool
	IntConsts() map[string]int64
	IntListConsts() map[string][]int64
	StrConsts() map[string]string
}

type Block interface {
	Element

	Masters() int64
	Width() int64

	//ConstContainer

	Configs() []Config
	//Funcs() []Func
	Masks() []Mask
	Statuses() []Status
	//Streams() []Stream
	Subblocks() []Block

	Sizes() access.Sizes
	AddrSpace() access.AddrSpace
}

type Config interface {
	Element

	Atomic() bool
	Default() val.BitStr
	Groups() []string
	//Range()   Range
	Once() bool
	Width() int64

	Access() access.Access
}

type Func interface {
	Element

	Params() []Param
	Returns() []Return

	StbAddr() int64
	AckAddr() int64
}

type Mask interface {
	Element

	Atomic() bool
	Default() val.BitStr
	Groups() []string
	Once() bool
	Width() int64

	Access() access.Access
}

type Param interface {
	Element

	Groups() []string
	Width() int64

	Access() access.Access
}

type Return interface {
	Element

	Groups() []string
	Width() int64

	Access() access.Access
}

type Status interface {
	Element

	Access() access.Access

	Atomic() bool
	Default() val.BitStr
	Groups() []string
	Once() bool
	Width() int64
}

type Stream interface {
	Element

	Params() []Param
	Returns() []Return

	StbAddr() int64
}

type Groupable interface {
	Element
	Groups() []string
}
