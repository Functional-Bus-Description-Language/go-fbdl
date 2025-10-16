package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Config struct {
	Func

	Atomic     bool
	InitValue  types.BitStr
	Range      types.Range
	ReadValue  types.BitStr
	ResetValue types.BitStr
	Width      int64

	Access access.Access
}

func (c Config) Type() string { return "config" }
