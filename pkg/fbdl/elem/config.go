package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Config struct {
	Elem

	Atomic    bool
	InitValue val.BitStr
	Groups    []string
	Range     val.Range
	Width     int64

	Access access.Access
}

func (c *Config) GroupNames() []string { return c.Groups }
