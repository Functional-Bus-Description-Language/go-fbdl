package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func hashRange(ran types.Range) uint32 {
	switch r := ran.(type) {
	case types.SingleRange:
		return hashSingleRange(r)
	case types.MultiRange:
		panic("unimplemented")
	default:
		panic("unimplemented")
	}
}

func hashSingleRange(sr types.SingleRange) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sr.Start)
	write(&buf, sr.End)

	return adler32.Checksum(buf.Bytes())
}
