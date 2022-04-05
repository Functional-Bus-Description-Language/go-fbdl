package ins

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"
)

type Group struct {
	Name  string
	Elems []*Element
}

// IsStatus returns true if group contains only status elements.
func (g *Group) IsStatus() bool {
	for _, e := range g.Elems {
		if e.Type != "status" {
			return false
		}
	}
	return true
}

// IsConfig returns true if group contains only config elements.
func (g *Group) IsConfig() bool {
	for _, e := range g.Elems {
		if e.Type != "config" {
			return false
		}
	}
	return true
}

// IsArray returns true if group contains only array elements.
func (g *Group) IsArray() bool {
	for _, e := range g.Elems {
		if !e.IsArray {
			return false
		}
	}
	return true
}

func (g *Group) hash() uint32 {
	b := bytes.Buffer{}

	b.Write([]byte(g.Name))
	aux := make([]byte, 4)
	for _, e := range g.Elems {
		binary.LittleEndian.PutUint32(aux, e.hash())
		b.Write(aux)
	}

	return adler32.Checksum(b.Bytes())
}
