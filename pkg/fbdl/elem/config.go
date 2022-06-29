package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// Config represents config element.
type Config struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	//Range   Range
	Once  bool
	Width int64
}

// HasDecreasingAccessOrder returns true if config must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic configs.
// If the end register is narrower, then starting writing from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (c *Config) HasDecreasingAccessOrder() bool {
	if !c.Atomic {
		return false
	}

	if asc, ok := c.Access.(access.SingleContinuous); ok {
		if !asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}
