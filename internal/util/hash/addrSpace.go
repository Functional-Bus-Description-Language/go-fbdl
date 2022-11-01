package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
)

func hashAddrSpace(as addrSpace.AddrSpace) uint32 {
	switch as := as.(type) {
	case addrSpace.Single:
		return hashAddrSpaceSingle(as)
	case addrSpace.Array:
		return hashAddrSpaceArray(as)
	default:
		panic("should never happen")
	}
}

func hashAddrSpaceSingle(s addrSpace.Single) uint32 {
	buf := bytes.Buffer{}

	write(&buf, s.Start)
	write(&buf, s.End)

	return adler32.Checksum(buf.Bytes())
}

func hashAddrSpaceArray(a addrSpace.Array) uint32 {
	buf := bytes.Buffer{}

	write(&buf, a.Start)
	write(&buf, a.Count)
	write(&buf, a.BlockSize)

	return adler32.Checksum(buf.Bytes())
}
