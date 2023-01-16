package elem

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Proc struct {
	Elem

	Delay *val.Time

	Params  []*Param
	Returns []*Return

	CallAddr int64
	ExitAddr int64
}

func (p *Proc) ParamsBufSize() int64 {
	params := p.Params
	l := len(params)

	if l == 0 {
		return 0
	}

	return params[l-1].Access.EndAddr() - params[0].Access.StartAddr() + 1
}

// ParamsStartAddr panics if proc has no params.
func (p *Proc) ParamsStartAddr() int64 {
	if len(p.Params) == 0 {
		panic(
			fmt.Sprintf("proc %s has no params", p.Name),
		)
	}

	return p.Params[0].Access.StartAddr()
}

func (p *Proc) ReturnsBufSize() int64 {
	rets := p.Returns
	l := len(rets)

	if l == 0 {
		return 0
	}

	return rets[l-1].Access.EndAddr() - rets[0].Access.StartAddr() + 1
}

// ReturnsStartAddr panics if proc has no returns.
func (p *Proc) ReturnsStartAddr() int64 {
	if len(p.Returns) == 0 {
		panic(
			fmt.Sprintf("proc %s has no returns", p.Name),
		)
	}

	return p.Returns[0].Access.StartAddr()
}
