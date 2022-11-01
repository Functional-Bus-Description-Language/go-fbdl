package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashParam(p *elem.Param) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&p.Elem))

	// Groups
	for _, g := range p.Groups {
		buf.Write([]byte(g))
	}

	// Width
	write(&buf, p.Width)

	// Access
	write(&buf, Hash(p.Access))

	return adler32.Checksum(buf.Bytes())
}
