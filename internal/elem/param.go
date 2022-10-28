package elem

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	fbdlAccess "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type param struct {
	Elem

	// Properties
	// TODO: Should Default be supported for param?
	// Default BitStr
	//Range Range
	Groups []string
	Width  int64

	Access fbdlAccess.Access
}

func (p *Param) Type() string { return "param" }

func (p *Param) SetGroups(g []string) { p.param.Groups = g }
func (p *Param) Groups() []string     { return p.param.Groups }

func (p *Param) SetWidth(w int64) { p.param.Width = w }
func (p *Param) Width() int64     { return p.param.Width }

func (p *Param) SetAccess(a fbdlAccess.Access) { p.param.Access = a }
func (p *Param) Access() fbdlAccess.Access     { return p.param.Access }

// Param represents param element.
type Param struct {
	param
}

func (p *Param) Hash() uint32 {
	buf := bytes.Buffer{}

	write := func(data any) {
		err := binary.Write(&buf, binary.LittleEndian, data)
		if err != nil {
			panic(err)
		}
	}

	// Elem
	write(p.Elem.Hash())

	// Groups
	for _, g := range p.Groups() {
		buf.Write([]byte(g))
	}

	// Width
	write(p.Width())

	// Access
	write(p.Access().(access.Access).Hash())

	return adler32.Checksum(buf.Bytes())
}
