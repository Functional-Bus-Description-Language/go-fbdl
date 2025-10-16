package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Mask struct {
	Func

	Atomic     bool
	InitValue  types.BitStr
	ReadValue  types.BitStr
	ResetValue types.BitStr
	Width      int64

	Access access.Access
}

func (m Mask) Type() string { return "mask" }
