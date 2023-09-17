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
				Type:     "SingleOneReg",
				Addr:     0,
				StartBit: 0,
				EndBit:   0,
			},
		},
		{1, 31, 2,
			SingleNRegs{
				Type:      "SingleNRegs",
				RegCount:  2,
				StartAddr: 1,
				StartBit:  31,
				EndBit:    0,
			},
		},
		{2, 30, 57,
			SingleNRegs{
				Type:      "SingleNRegs",
				RegCount:  3,
				StartAddr: 2,
				StartBit:  30,
				EndBit:    22,
			},
		},
		{3, 0, 32,
			SingleOneReg{
				Type:     "SingleOneReg",
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
				Type:      "ArrayNRegs",
				RegCount:  1,
				ItemCount: 1,
				ItemWidth: 32,
				StartAddr: 0,
				StartBit:  0,
			},
		},
		{1, 4, 0, 5,
			ArrayNRegs{
				Type:      "ArrayNRegs",
				RegCount:  1,
				ItemCount: 4,
				ItemWidth: 5,
				StartAddr: 1,
				StartBit:  0,
			},
		},
		{2, 2, 20, 23,
			ArrayNRegs{
				Type:      "ArrayNRegs",
				RegCount:  3,
				ItemCount: 2,
				ItemWidth: 23,
				StartAddr: 2,
				StartBit:  20,
			},
		},
		{3, 2, 20, 22,
			ArrayNRegs{
				Type:      "ArrayNRegs",
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
				Type:       "ArrayNInReg",
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
				Type:       "ArrayNInReg",
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
				Type:       "ArrayNInReg",
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
				Type:       "ArrayNInReg",
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

func TestMakeArrayNInRegMInEndReg(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 5, 7,
			ArrayNInRegMInEndReg{
				Type:          "ArrayNInRegMInEndReg",
				RegCount:      2,
				ItemCount:     5,
				ItemWidth:     7,
				ItemsInReg:    4,
				ItemsInEndReg: 1,
				StartAddr:     0,
				StartBit:      0,
			},
		},
		{1, 66, 1,
			ArrayNInRegMInEndReg{
				Type:          "ArrayNInRegMInEndReg",
				RegCount:      3,
				ItemCount:     66,
				ItemWidth:     1,
				ItemsInReg:    32,
				ItemsInEndReg: 2,
				StartAddr:     1,
				StartBit:      0,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNInRegMInEndReg(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayOneInNRegs(t *testing.T) {
	var tests = []struct {
		startAddr    int64
		count        int64
		width        int64
		want         Access
		wantRegCount int64
		wantEndBit   int64
	}{
		{0, 2, 33,
			ArrayOneInNRegs{
				Type:      "ArrayOneInNRegs",
				ItemCount: 2,
				ItemWidth: 33,
				StartAddr: 0,
			},
			4, 0,
		},
		{1, 3, 64,
			ArrayOneInNRegs{
				Type:      "ArrayOneInNRegs",
				ItemCount: 3,
				ItemWidth: 64,
				StartAddr: 1,
			},
			6, 31,
		},
	}

	for i, test := range tests {
		got := MakeArrayOneInNRegs(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}

		if got.GetRegCount() != test.wantRegCount {
			t.Errorf("[%d] got %d, want %d", i, got.GetRegCount(), test.wantRegCount)
		}

		if got.GetEndBit() != test.wantEndBit {
			t.Errorf("[%d] got %d, want %d", i, got.GetEndBit(), test.wantEndBit)
		}
	}
}
