package access

import (
	"bytes"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
)

type AddrSpace interface {
	Hash() uint32
}

type ass struct {
	Start int64
	End   int64
}

type AddrSpaceSingle struct {
	ass
}

func (ass AddrSpaceSingle) Start() int64 { return ass.ass.Start }
func (ass AddrSpaceSingle) End() int64   { return ass.ass.End }

func (ass AddrSpaceSingle) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, ass.Start())
	hash.Write(&buf, ass.End())
	return hash.Hash(buf)
}

func MakeAddrSpaceSingle(start, end int64) AddrSpaceSingle {
	return AddrSpaceSingle{
		ass: ass{
			Start: start,
			End:   end,
		},
	}
}

type asa struct {
	Start     int64
	End       int64
	BlockSize int64
	Count     int64
}

type AddrSpaceArray struct {
	asa
}

func (asa AddrSpaceArray) Start() int64     { return asa.asa.Start }
func (asa AddrSpaceArray) End() int64       { return asa.asa.End }
func (asa AddrSpaceArray) BlockSize() int64 { return asa.asa.BlockSize }
func (asa AddrSpaceArray) Count() int64     { return asa.asa.Count }

func (asa AddrSpaceArray) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, asa.Start())
	hash.Write(&buf, asa.End())
	hash.Write(&buf, asa.BlockSize())
	hash.Write(&buf, asa.Count())
	return hash.Hash(buf)
}

func MakeAddrSpaceArray(start, count, blockSize int64) AddrSpaceArray {
	return AddrSpaceArray{
		asa: asa{
			Start:     start,
			End:       start + count*blockSize - 1,
			BlockSize: blockSize,
			Count:     count,
		},
	}
}
