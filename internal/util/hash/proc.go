package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashProc(p *elem.Proc) uint32 {
	buf := bytes.Buffer{}

	// Elem
	write(&buf, Hash(&p.Elem))

	// Params
	for _, p := range p.Params {
		write(&buf, Hash(p))
	}

	// Returns
	for _, r := range p.Returns {
		write(&buf, Hash(r))
	}

	// StbAddr
	write(&buf, p.StbAddr)

	// AckAddr
	write(&buf, p.AckAddr)

	return adler32.Checksum(buf.Bytes())
}
