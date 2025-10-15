package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
)

func hashRange(r addrSpace.Range) uint32 {
	buf := bytes.Buffer{}

	write(&buf, r.Start)
	write(&buf, r.End)

	return adler32.Checksum(buf.Bytes())
}
