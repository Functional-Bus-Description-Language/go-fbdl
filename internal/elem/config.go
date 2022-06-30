package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type cfg struct {
	Elem

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	//Range   Range
	Once  bool
	Width int64

	Access access.Access
}

// Config represents config element.
type Config struct {
	cfg
}

func (c *Config) Type() string { return "config" }

func (c *Config) SetAtomic(a bool) { c.cfg.Atomic = a }
func (c *Config) Atomic() bool     { return c.cfg.Atomic }

func (c *Config) SetDefault(d val.BitStr) { c.cfg.Default = d }
func (c *Config) Default() val.BitStr     { return c.cfg.Default }

func (c *Config) SetGroups(g []string) { c.cfg.Groups = g }
func (c *Config) Groups() []string     { return c.cfg.Groups }

func (c *Config) SetOnce(a bool) { c.cfg.Once = a }
func (c *Config) Once() bool     { return c.cfg.Once }

func (c *Config) SetWidth(w int64) { c.cfg.Width = w }
func (c *Config) Width() int64     { return c.cfg.Width }

func (c *Config) SetAccess(a access.Access) { c.cfg.Access = a }
func (c *Config) Access() access.Access     { return c.cfg.Access }

// HasDecreasingAccessOrder returns true if config must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic configs.
// If the end register is narrower, then starting writing from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (c *Config) HasDecreasingAccessOrder() bool {
	if !c.cfg.Atomic {
		return false
	}

	if asc, ok := c.cfg.Access.(access.SingleContinuous); ok {
		if !asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}

func (c *Config) Hash() int64 {
	return 0
}
