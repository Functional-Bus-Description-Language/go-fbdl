package fn

import (
	"fmt"
)

type Functionality interface {
	isFunctionality()
}

type Func struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
}

func (f Func) isFunctionality() {}

func Name(f Functionality) string {
	switch f := f.(type) {
	case *Block:
		return f.Name
	case *Config:
		return f.Name
	case *Irq:
		return f.Name
	case *Mask:
		return f.Name
	case *Memory:
		return f.Name
	case *Param:
		return f.Name
	case *Proc:
		return f.Name
	case *Return:
		return f.Name
	case *Static:
		return f.Name
	case *Status:
		return f.Name
	case *Stream:
		return f.Name
	default:
		panic(
			fmt.Sprintf("%T is not an element", f),
		)
	}
}

func Type(f Functionality) string {
	switch f.(type) {
	case *Block:
		return "block"
	case *Config:
		return "config"
	case *Irq:
		return "irq"
	case *Mask:
		return "mask"
	case *Memory:
		return "memory"
	case *Param:
		return "param"
	case *Proc:
		return "proc"
	case *Return:
		return "return"
	case *Static:
		return "static"
	case *Status:
		return "status"
	case *Stream:
		return "stream"
	default:
		panic(
			fmt.Sprintf("%T is not functionality", f),
		)
	}
}
