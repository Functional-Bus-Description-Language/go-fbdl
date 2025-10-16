package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Static struct {
	Func

	InitValue  types.BitStr
	ReadValue  types.BitStr
	ResetValue types.BitStr
	Width      int64

	Access types.Access
}

func (s Static) Type() string { return "static" }
