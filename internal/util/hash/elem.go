package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func hashElem(e *elem.Elem) uint32 {
	buf := bytes.Buffer{}

	// Name
	buf.Write([]byte(e.Name))

	// Doc
	buf.Write([]byte(e.Doc))

	// IsArray
	if e.IsArray {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}

	// Count
	err := binary.Write(&buf, binary.LittleEndian, e.Count)
	if err != nil {
		panic(err)
	}

	return adler32.Checksum(buf.Bytes())
}
