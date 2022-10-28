package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	fbdlAccess "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Return represents return element.
type ret struct {
	Elem

	// Properties
	Groups []string
	Width  int64

	Access fbdlAccess.Access
}

// Return represents return element.
type Return struct {
	ret
}

func (r *Return) Type() string { return "return" }

func (r *Return) SetGroups(g []string) { r.ret.Groups = g }
func (r *Return) Groups() []string     { return r.ret.Groups }

func (r *Return) SetWidth(w int64) { r.ret.Width = w }
func (r *Return) Width() int64     { return r.ret.Width }

func (r *Return) SetAccess(a fbdlAccess.Access) { r.ret.Access = a }
func (r *Return) Access() fbdlAccess.Access     { return r.ret.Access }

func (r *Return) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(r.Elem.Hash())

	// Groups
	for _, g := range r.Groups() {
		buf.Write([]byte(g))
	}

	// Width
	write(r.Width())

	// Access
	write(r.Access().(access.Access).Hash())

	return adler32.Checksum(buf.Bytes())
}
