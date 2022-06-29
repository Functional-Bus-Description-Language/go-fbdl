package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Func represents func element.
type Func struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	StbAddr int64 // Strobe address
	AckAddr int64 // Acknowledgment address

	// Properties

	Params  []*Param
	Returns []*Return
}

func (f *Func) ParamsStartAddr() int64 {
	if len(f.Params) == 0 {
		return f.StbAddr
	}

	return f.Params[0].Access.StartAddr()
}

// AreAllParamsSingleSingle returns true if accesses to all parameters are of type AccessSingleSingle.
func (f *Func) AreAllParamsSingleSingle() bool {
	for _, p := range f.Params {
		switch p.Access.(type) {
		case access.SingleSingle:
			continue
		default:
			return false
		}
	}
	return true
}
