package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
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

func (c *Config) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(c.Elem.Hash())

	// Atomic
	write(c.Atomic())

	// Default
	buf.Write([]byte(c.Default()))

	// Groups
	for _, g := range c.Groups() {
		buf.Write([]byte(g))
	}

	// Once
	write(c.Once())

	// Width
	write(c.Width())

	// Access
	write(hash.AccessAccess(c.Access()))

	return adler32.Checksum(buf.Bytes())
}
