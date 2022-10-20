package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type st struct {
	Elem

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64

	Access access.Access
}

// Status represents status element.
type Status struct {
	st
}

func (s *Status) Type() string { return "status" }

func (c *Status) SetAtomic(a bool) { c.st.Atomic = a }
func (c *Status) Atomic() bool     { return c.st.Atomic }

func (c *Status) SetDefault(d val.BitStr) { c.st.Default = d }
func (c *Status) Default() val.BitStr     { return c.st.Default }

func (c *Status) SetGroups(g []string) { c.st.Groups = g }
func (c *Status) Groups() []string     { return c.st.Groups }

func (c *Status) SetOnce(a bool) { c.st.Once = a }
func (c *Status) Once() bool     { return c.st.Once }

func (c *Status) SetWidth(w int64) { c.st.Width = w }
func (c *Status) Width() int64     { return c.st.Width }

func (c *Status) SetAccess(a access.Access) { c.st.Access = a }
func (c *Status) Access() access.Access     { return c.st.Access }

func (s *Status) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(s.Elem.Hash())

	// Atomic
	write(s.Atomic())

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
	write(hash.AccessAccess(s.Access()))

	return adler32.Checksum(buf.Bytes())
}
