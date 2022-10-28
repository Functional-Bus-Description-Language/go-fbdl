// Package hash implements hash calculation for public types.
// The types are public, however their hash functions should not be public.
package hash

import (
	"bytes"
	"encoding/binary"
	"hash/adler32"
	"io"
)

func Write(buf io.Writer, data any) {
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}
}

func Hash(buf bytes.Buffer) uint32 {
	return adler32.Checksum(buf.Bytes())
}
