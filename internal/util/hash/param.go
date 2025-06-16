package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashParam(p *fn.Param) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&p.Func))

	// Width
	write(&buf, p.Width)

	// Access
	write(&buf, Hash(p.Access))

	return adler32.Checksum(buf.Bytes())
}
