package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/value"
)

type Config struct {
	Func

	Atomic     bool
	InitValue  value.BitStr
	Range      value.Range
	ReadValue  value.BitStr
	ResetValue value.BitStr
	Width      int64

	Access access.Access
}

func (c Config) Type() string { return "config" }
