package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

type Static struct {
	Func

	InitValue  value.BitStr
	ReadValue  value.BitStr
	ResetValue value.BitStr
	Width      int64

	Access access.Access
}

func (s Static) Type() string { return "static" }
