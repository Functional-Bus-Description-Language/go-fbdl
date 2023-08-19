package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Config struct {
	Func

	Atomic     bool
	InitValue  val.BitStr
	Groups     []string
	Range      val.Range
	ReadValue  val.BitStr
	ResetValue val.BitStr
	Width      int64

	Access access.Access
}

func (c *Config) GroupNames() []string { return c.Groups }
