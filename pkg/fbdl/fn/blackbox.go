package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Blackbox struct {
	Func

	Size int64

	Sizes     types.Sizes
	AddrSpace types.SingleRange
}

func (b Blackbox) Type() string { return "blackbox" }
