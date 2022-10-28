package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	fbdlAccess "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type static struct {
	Elem

	// Properties
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64

	Access fbdlAccess.Access
}

// Static represents static element.
type Static struct {
	static
}

func (s *Static) Type() string { return "static" }

func (c *Static) SetDefault(d val.BitStr) { c.static.Default = d }
func (c *Static) Default() val.BitStr     { return c.static.Default }

func (c *Static) SetGroups(g []string) { c.static.Groups = g }
func (c *Static) Groups() []string     { return c.static.Groups }

func (c *Static) SetOnce(a bool) { c.static.Once = a }
func (c *Static) Once() bool     { return c.static.Once }

func (c *Static) SetWidth(w int64) { c.static.Width = w }
func (c *Static) Width() int64     { return c.static.Width }

func (c *Static) SetAccess(a fbdlAccess.Access) { c.static.Access = a }
func (c *Static) Access() fbdlAccess.Access     { return c.static.Access }

func (s *Static) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(s.Elem.Hash())

	// Default
	buf.Write([]byte(s.Default()))

	// Groups
	for _, g := range s.Groups() {
		buf.Write([]byte(g))
	}

	// Once
	write(s.Once())

	// Width
	write(s.Width())

	// Access
	write(s.Access().(access.Access).Hash())

	return adler32.Checksum(buf.Bytes())
}
