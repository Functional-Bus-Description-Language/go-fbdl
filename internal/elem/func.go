package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	fbdl "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

type fun struct {
	Elem

	// Properties
	// Currently Func has no properties.

	Params  []fbdl.Param
	Returns []fbdl.Return

	StbAddr int64 // Strobe address
	AckAddr int64 // Acknowledgment address
}

// Func represents func element.
type Func struct {
	fun
}

func (f *Func) Type() string { return "func" }

func (f *Func) SetStbAddr(a int64) { f.fun.StbAddr = a }
func (f *Func) StbAddr() int64     { return f.fun.StbAddr }

func (f *Func) SetAckAddr(a int64) { f.fun.AckAddr = a }
func (f *Func) AckAddr() int64     { return f.fun.AckAddr }

func (f *Func) AddParam(p *Param)    { f.fun.Params = append(f.fun.Params, p) }
func (f *Func) Params() []fbdl.Param { return f.fun.Params }

func (f *Func) AddReturn(r *Return)    { f.fun.Returns = append(f.fun.Returns, r) }
func (f *Func) Returns() []fbdl.Return { return f.fun.Returns }

func (f *Func) HasElement(name string) bool {
	for i, _ := range f.fun.Params {
		if f.fun.Params[i].Name() == name {
			return true
		}
	}
	for i, _ := range f.fun.Returns {
		if f.fun.Returns[i].Name() == name {
			return true
		}
	}
	return false
}

func (f *Func) ParamsStartAddr() int64 {
	if len(f.fun.Params) == 0 {
		return f.fun.StbAddr
	}

	return f.fun.Params[0].Access().StartAddr()
}

// AreAllParamsSingleSingle returns true if accesses to all parameters are of type AccessSingleSingle.
func (f *Func) AreAllParamsSingleSingle() bool {
	for _, p := range f.fun.Params {
		switch p.Access().(type) {
		case access.SingleSingle:
			continue
		default:
			return false
		}
	}
	return true
}

func (f *Func) Hash() int64 {
	return 0
}
