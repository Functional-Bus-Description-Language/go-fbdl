package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

type Status struct {
	Func

	Atomic    bool
	ReadValue value.BitStr
	Width     int64

	Access access.Access
}

func (s Status) Type() string { return "status" }
