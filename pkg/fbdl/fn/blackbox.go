package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

type Blackbox struct {
	Func

	Size int64

	Sizes     access.Sizes
	AddrSpace value.SingleRange
}

func (b Blackbox) Type() string { return "blackbox" }
