package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Mask struct {
	Func

	Atomic     bool
	InitValue  val.BitStr
	ReadValue  val.BitStr
	ResetValue val.BitStr
	Width      int64

	Access access.Access
}

func (m Mask) Type() string { return "mask" }
