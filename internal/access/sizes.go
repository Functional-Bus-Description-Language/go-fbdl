package access

import (
	"bytes"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
)

type sizes struct {
	Own          int64
	Compact      int64
	BlockAligned int64
}

type Sizes struct {
	sizes
}

func (s Sizes) Own() int64          { return s.sizes.Own }
func (s Sizes) Compact() int64      { return s.sizes.Compact }
func (s Sizes) BlockAligned() int64 { return s.sizes.BlockAligned }

func MakeSizes(own, compact, blockAligned int64) Sizes {
	return Sizes{
		sizes: sizes{
			Own:          own,
			Compact:      compact,
			BlockAligned: blockAligned,
		},
	}
}

func (s Sizes) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, s.Own())
	hash.Write(&buf, s.Compact())
	hash.Write(&buf, s.BlockAligned())
	return hash.Hash(buf)
}
