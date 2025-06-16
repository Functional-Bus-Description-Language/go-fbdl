package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func hashReturn(r *fn.Return) uint32 {
	buf := bytes.Buffer{}

	// Func
	write(&buf, Hash(&r.Func))

	// Width
	write(&buf, r.Width)

	// Access
	write(&buf, Hash(r.Access))

	return adler32.Checksum(buf.Bytes())
}
