package elem

import (
	"fmt"
)

type Element interface {
	isElement() bool
}

type Elem struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
}

func (e Elem) isElement() bool { return true }

func Name(e Element) string {
	switch e := e.(type) {
	case *Block:
		return e.Name
	case *Config:
		return e.Name
	case *Mask:
		return e.Name
	case *Memory:
		return e.Name
	case *Param:
		return e.Name
	case *Proc:
		return e.Name
	case *Return:
		return e.Name
	case *Static:
		return e.Name
	case *Status:
		return e.Name
	case *Stream:
		return e.Name
	default:
		panic(
			fmt.Sprintf("%T is not an element", e),
		)
	}
}

func Type(e Element) string {
	switch e.(type) {
	case *Block:
		return "block"
	case *Config:
		return "config"
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
			fmt.Sprintf("%T is not an element", e),
		)
	}
}
