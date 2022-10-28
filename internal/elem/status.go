package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	fbdlAccess "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type status struct {
	Elem

	// Properties
	Atomic bool
	Groups []string
	Once   bool
	Width  int64

	Access fbdlAccess.Access
}

// Status represents status element.
type Status struct {
	status
}

func (s *Status) Type() string { return "status" }

func (c *Status) SetAtomic(a bool) { c.status.Atomic = a }
func (c *Status) Atomic() bool     { return c.status.Atomic }

func (c *Status) SetGroups(g []string) { c.status.Groups = g }
func (c *Status) Groups() []string     { return c.status.Groups }

func (c *Status) SetOnce(a bool) { c.status.Once = a }
func (c *Status) Once() bool     { return c.status.Once }

func (c *Status) SetWidth(w int64) { c.status.Width = w }
func (c *Status) Width() int64     { return c.status.Width }

func (c *Status) SetAccess(a fbdlAccess.Access) { c.status.Access = a }
func (c *Status) Access() fbdlAccess.Access     { return c.status.Access }

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
