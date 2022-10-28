package access

import (
	"reflect"
	"testing"
)

func init() {
	busWidth = 32
}

func TestMakeSingle(t *testing.T) {
	var tests = []struct {
		baseAddr int64
		baseBit  int64
		width    int64
		want     Access
	}{
		{0, 0, 1, SingleSingle{ss: ss{Addr: 0, Mask: makeMask(0, 0)}}},
		{1, 31, 2, SingleContinuous{
			sc: sc{
				StartAddr: 1,
				EndAddr:   2,
				StartMask: makeMask(31, 31),
				EndMask:   makeMask(0, 0),
			},
		}},
		{2, 30, 57, SingleContinuous{
			sc: sc{
				StartAddr: 2,
				EndAddr:   4,
				StartMask: makeMask(30, 31),
				EndMask:   makeMask(0, 22),
			},
		}},
		{3, 0, 32, SingleSingle{ss: ss{Addr: 3, Mask: makeMask(0, 31)}}},
	}

	for i, test := range tests {
		got := MakeSingle(test.baseAddr, test.baseBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
