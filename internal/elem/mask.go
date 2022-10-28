package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	fbdlAccess "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// Mask represents mask element.
type mask struct {
	Elem

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64

	Access fbdlAccess.Access
}

type Mask struct {
	mask
}

func (m *Mask) Type() string { return "mask" }

func (m *Mask) SetAtomic(a bool) { m.mask.Atomic = a }
func (m *Mask) Atomic() bool     { return m.mask.Atomic }

func (m *Mask) SetDefault(d val.BitStr) { m.mask.Default = d }
func (m *Mask) Default() val.BitStr     { return m.mask.Default }

func (m *Mask) SetGroups(g []string) { m.mask.Groups = g }
func (m *Mask) Groups() []string     { return m.mask.Groups }

func (m *Mask) SetOnce(a bool) { m.mask.Once = a }
func (m *Mask) Once() bool     { return m.mask.Once }

func (m *Mask) SetWidth(w int64) { m.mask.Width = w }
func (m *Mask) Width() int64     { return m.mask.Width }

func (m *Mask) SetAccess(a fbdlAccess.Access) { m.mask.Access = a }
func (m *Mask) Access() fbdlAccess.Access     { return m.mask.Access }

func (m *Mask) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(m.Elem.Hash())

	// Atomic
	write(m.Atomic())

	// Default
	buf.Write([]byte(m.Default()))

	// Groups
	for _, g := range m.Groups() {
		buf.Write([]byte(g))
	}

	// Once
	write(m.Once())

	// Width
	write(m.Width())

	// Access
	write(m.Access().(access.Access).Hash())

	return adler32.Checksum(buf.Bytes())
}
