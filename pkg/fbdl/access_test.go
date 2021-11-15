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
		{0, 0, 1, AccessSingle{Strategy: "Single", Address: 0, count: 1, Mask: Mask{Upper: 0, Lower: 0}}},
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
		if got.Mask.Upper != test.want.Mask.Upper {
			t.Errorf("[%d] Mask.Upper differs: %v %v", i, got, test)
		}
		if got.Mask.Lower != test.want.Mask.Lower {
			t.Errorf("[%d] Mask.Lower differs: %v %v", i, got, test)
		}
	}
}
