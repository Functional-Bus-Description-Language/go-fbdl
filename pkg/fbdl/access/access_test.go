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
				typ:      "SingleOneReg",
				addr:     0,
				startBit: 0,
				endBit:   0,
			},
		},
		{1, 31, 2,
			SingleNRegs{
				typ:       "SingleNRegs",
				regCount:  2,
				startAddr: 1,
				startBit:  31,
				endBit:    0,
			},
		},
		{2, 30, 57,
			SingleNRegs{
				typ:       "SingleNRegs",
				regCount:  3,
				startAddr: 2,
				startBit:  30,
				endBit:    22,
			},
		},
		{3, 0, 32,
			SingleOneReg{
				typ:      "SingleOneReg",
				addr:     3,
				startBit: 0,
				endBit:   31,
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
				typ:       "ArrayNRegs",
				regCount:  1,
				itemCount: 1,
				itemWidth: 32,
				startAddr: 0,
				startBit:  0,
			},
		},
		{1, 4, 0, 5,
			ArrayNRegs{
				typ:       "ArrayNRegs",
				regCount:  1,
				itemCount: 4,
				itemWidth: 5,
				startAddr: 1,
				startBit:  0,
			},
		},
		{2, 2, 20, 23,
			ArrayNRegs{
				typ:       "ArrayNRegs",
				regCount:  3,
				itemCount: 2,
				itemWidth: 23,
				startAddr: 2,
				startBit:  20,
			},
		},
		{3, 2, 20, 22,
			ArrayNRegs{
				typ:       "ArrayNRegs",
				regCount:  2,
				itemCount: 2,
				itemWidth: 22,
				startAddr: 3,
				startBit:  20,
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
				typ:        "ArrayNInReg",
				regCount:   2,
				itemCount:  4,
				itemWidth:  16,
				itemsInReg: 2,
				startAddr:  2,
				startBit:   0,
			},
		},
		{4, 8, 8,
			ArrayNInReg{
				typ:        "ArrayNInReg",
				regCount:   2,
				itemCount:  8,
				itemWidth:  8,
				itemsInReg: 4,
				startAddr:  4,
				startBit:   0,
			},
		},
		{5, 12, 7,
			ArrayNInReg{
				typ:        "ArrayNInReg",
				regCount:   3,
				itemCount:  12,
				itemWidth:  7,
				itemsInReg: 4,
				startAddr:  5,
				startBit:   0,
			},
		},
		{6, 50, 3,
			ArrayNInReg{
				typ:        "ArrayNInReg",
				regCount:   5,
				itemCount:  50,
				itemWidth:  3,
				itemsInReg: 10,
				startAddr:  6,
				startBit:   0,
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
				typ:           "ArrayNInRegMInEndReg",
				regCount:      2,
				itemCount:     5,
				itemWidth:     7,
				itemsInReg:    4,
				itemsInEndReg: 1,
				startAddr:     0,
				startBit:      0,
			},
		},
		{1, 66, 1,
			ArrayNInRegMInEndReg{
				typ:           "ArrayNInRegMInEndReg",
				regCount:      3,
				itemCount:     66,
				itemWidth:     1,
				itemsInReg:    32,
				itemsInEndReg: 2,
				startAddr:     1,
				startBit:      0,
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
				typ:       "ArrayOneInNRegs",
				itemCount: 2,
				itemWidth: 33,
				startAddr: 0,
			},
			4, 0,
		},
		{1, 3, 64,
			ArrayOneInNRegs{
				typ:       "ArrayOneInNRegs",
				itemCount: 3,
				itemWidth: 64,
				startAddr: 1,
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
