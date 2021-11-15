package fbdl

import (
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
		want     AccessSingle
	}{
		{0, 0, 1,
			AccessSingle{
				Strategy:  "Single",
				Address:   0,
				count:     1,
				FirstMask: Mask{Upper: 0, Lower: 0},
				LastMask:  Mask{Upper: 0, Lower: 0},
			},
		},
		{1, 31, 2,
			AccessSingle{
				Strategy:  "Linear",
				Address:   1,
				count:     2,
				FirstMask: Mask{Upper: 31, Lower: 31},
				LastMask:  Mask{Upper: 0, Lower: 0},
			},
		},
		{2, 30, 57,
			AccessSingle{
				Strategy:  "Linear",
				Address:   2,
				count:     3,
				FirstMask: Mask{Upper: 31, Lower: 30},
				LastMask:  Mask{Upper: 22, Lower: 0},
			},
		},
		{3, 0, 32,
			AccessSingle{
				Strategy:  "Single",
				Address:   3,
				count:     1,
				FirstMask: Mask{Upper: 31, Lower: 0},
				LastMask:  Mask{Upper: 31, Lower: 0},
			},
		},
	}

	for i, test := range tests {
		got := makeAccessSingle(test.baseAddr, test.baseBit, test.width)

		if got.Strategy != test.want.Strategy {
			t.Errorf("[%d] Strategy differs: %v %v", i, got, test)
		}
		if got.Address != test.want.Address {
			t.Errorf("[%d] Address differs: %v %v", i, got, test)
		}
		if got.count != test.want.count {
			t.Errorf("[%d] count differs: %v %v", i, got, test)
		}
		if got.FirstMask.Upper != test.want.FirstMask.Upper {
			t.Errorf("[%d] FirstMask.Upper differs: %v %v", i, got, test)
		}
		if got.FirstMask.Lower != test.want.FirstMask.Lower {
			t.Errorf("[%d] FirstMask.Lower differs: %v %v", i, got, test)
		}
		if got.LastMask.Upper != test.want.LastMask.Upper {
			t.Errorf("[%d] LastMask.Upper differs: %v %v", i, got, test)
		}
		if got.LastMask.Lower != test.want.LastMask.Lower {
			t.Errorf("[%d] LastMask.Lower differs: %v %v", i, got, test)
		}
	}
}
