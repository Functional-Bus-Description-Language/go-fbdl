package hash

import (
	"bytes"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

func hashAccessSizes(sizes access.Sizes) uint32 {
	buf := bytes.Buffer{}

	write(&buf, sizes.BlockAligned)
	write(&buf, sizes.Compact)
	write(&buf, sizes.Own)

	return adler32.Checksum(buf.Bytes())
}

func hashAccessAccess(a access.Access) uint32 {
	buf := bytes.Buffer{}

	write(&buf, a.GetRegCount())
	write(&buf, a.GetStartAddr())
	write(&buf, a.GetEndAddr())
	write(&buf, a.GetStartBit())
	write(&buf, a.GetEndBit())
	write(&buf, a.GetStartRegWidth())
	write(&buf, a.GetEndRegWidth())

	return adler32.Checksum(buf.Bytes())
}
