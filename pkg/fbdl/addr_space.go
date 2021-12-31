package fbdl

import (
	"encoding/json"
)

type AddrSpace interface {
	Start() int64
	End() int64
	IsArray() bool
	Count() int64
}

type AddrSpaceSingle struct {
	start int64
	end   int64
}

func (s AddrSpaceSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Start, End int64
	}{
		Start: s.start,
		End:   s.end,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (s AddrSpaceSingle) Start() int64  { return s.start }
func (s AddrSpaceSingle) End() int64    { return s.end }
func (s AddrSpaceSingle) IsArray() bool { return false }
func (s AddrSpaceSingle) Count() int64  { return 1 }

type AddrSpaceArray struct {
	start     int64
	count     int64
	BlockSize int64
}

func (a AddrSpaceArray) IsArray() bool { return true }
func (a AddrSpaceArray) Count() int64  { return a.count }
func (a AddrSpaceArray) Start() int64  { return a.start }

func (a AddrSpaceArray) GetAddress(i int64) (start int64, end int64) {
	start = a.start + i*a.BlockSize
	end = start + a.BlockSize - 1

	return
}

func (a AddrSpaceArray) End() int64 {
	return a.start + a.count*a.BlockSize - 1
}
