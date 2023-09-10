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
		{0, 0, 1,
			SingleOneReg{
				Strategy: "SingleOneReg",
				Addr:     0,
				StartBit: 0,
				EndBit:   0,
			},
		},
		{1, 31, 2,
			SingleNRegs{
				Strategy:  "SingleNRegs",
				RegCount:  2,
				StartAddr: 1,
				StartBit:  31,
				EndBit:    0,
			},
		},
		{2, 30, 57,
			SingleNRegs{
				Strategy:  "SingleNRegs",
				RegCount:  3,
				StartAddr: 2,
				StartBit:  30,
				EndBit:    22,
			},
		},
		{3, 0, 32,
			SingleOneReg{
				Strategy: "SingleOneReg",
				Addr:     3,
				StartBit: 0,
				EndBit:   31,
			},
		},
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

func TestMakeArrayNRegs(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		startBit  int64
		width     int64
		want      Access
	}{
		{0, 1, 0, 32,
			ArrayNRegs{
				Strategy:  "ArrayNRegs",
				RegCount:  1,
				ItemCount: 1,
				ItemWidth: 32,
				StartAddr: 0,
				StartBit:  0,
			},
		},
		{1, 4, 0, 5,
			ArrayNRegs{
				Strategy:  "ArrayNRegs",
				RegCount:  1,
				ItemCount: 4,
				ItemWidth: 5,
				StartAddr: 1,
				StartBit:  0,
			},
		},
		{2, 2, 20, 23,
			ArrayNRegs{
				Strategy:  "ArrayNRegs",
				RegCount:  3,
				ItemCount: 2,
				ItemWidth: 23,
				StartAddr: 2,
				StartBit:  20,
			},
		},
		{3, 2, 20, 22,
			ArrayNRegs{
				Strategy:  "ArrayNRegs",
				RegCount:  2,
				ItemCount: 2,
				ItemWidth: 22,
				StartAddr: 3,
				StartBit:  20,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNRegs(test.count, test.startAddr, test.startBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayNInReg(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{2, 4, 16,
			ArrayNInReg{
				Strategy:   "ArrayNInReg",
				RegCount:   2,
				ItemCount:  4,
				ItemWidth:  16,
				ItemsInReg: 2,
				StartAddr:  2,
				StartBit:   0,
			},
		},
		{4, 8, 8,
			ArrayNInReg{
				Strategy:   "ArrayNInReg",
				RegCount:   2,
				ItemCount:  8,
				ItemWidth:  8,
				ItemsInReg: 4,
				StartAddr:  4,
				StartBit:   0,
			},
		},
		{5, 12, 7,
			ArrayNInReg{
				Strategy:   "ArrayNInReg",
				RegCount:   3,
				ItemCount:  12,
				ItemWidth:  7,
				ItemsInReg: 4,
				StartAddr:  5,
				StartBit:   0,
			},
		},
		{6, 50, 3,
			ArrayNInReg{
				Strategy:   "ArrayNInReg",
				RegCount:   5,
				ItemCount:  50,
				ItemWidth:  3,
				ItemsInReg: 10,
				StartAddr:  6,
				StartBit:   0,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNInReg(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
