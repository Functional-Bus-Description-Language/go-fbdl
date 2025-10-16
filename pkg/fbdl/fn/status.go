package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Status struct {
	Func

	Atomic    bool
	ReadValue types.BitStr
	Width     int64

	Access types.Access
}

func (s Status) Type() string { return "status" }
