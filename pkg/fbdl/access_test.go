package fbdl

import (
	"reflect"
	"testing"
)

func init() {
	busWidth = 32
}

func TestMakeAccessSingle(t *testing.T) {
	var tests = []struct {
		baseAddr int64
		baseBit  int64
		width    int64
		want     Access
	}{
		{0, 0, 1,
			AccessSingleSingle{
				Addr: 0,
				Mask: Mask{Upper: 0, Lower: 0},
			},
		},
		{1, 31, 2,
			AccessSingleContinuous{
				count:     2,
				StartAddr: 1,
				StartMask: Mask{Upper: 31, Lower: 31},
				EndMask:   Mask{Upper: 0, Lower: 0},
			},
		},
		{2, 30, 57,
			AccessSingleContinuous{
				count:     3,
				StartAddr: 2,
				StartMask: Mask{Upper: 31, Lower: 30},
				EndMask:   Mask{Upper: 22, Lower: 0},
			},
		},
		{3, 0, 32,
			AccessSingleSingle{
				Addr: 3,
				Mask: Mask{Upper: 31, Lower: 0},
			},
		},
	}

	for i, test := range tests {
		got := makeAccessSingle(test.baseAddr, test.baseBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
