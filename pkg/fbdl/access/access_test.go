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
			SingleContinuous{
				regCount:  2,
				startAddr: 1,
				startBit:  31,
				endBit:    0,
			},
		},
		{2, 30, 57,
			SingleContinuous{
				regCount:  3,
				startAddr: 2,
				startBit:  30,
				endBit:    22,
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

func TestMakeArrayContinuous(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		startBit  int64
		width     int64
		want      Access
	}{
		{0, 1, 0, 32,
			ArrayContinuous{
				regCount:  1,
				ItemCount: 1,
				ItemWidth: 32,
				startAddr: 0,
				startBit:  0,
			},
		},
		{1, 4, 0, 5,
			ArrayContinuous{
				regCount:  1,
				ItemCount: 4,
				ItemWidth: 5,
				startAddr: 1,
				startBit:  0,
			},
		},
		{2, 2, 20, 23,
			ArrayContinuous{
				regCount:  3,
				ItemCount: 2,
				ItemWidth: 23,
				startAddr: 2,
				startBit:  20,
			},
		},
		{3, 2, 20, 22,
			ArrayContinuous{
				regCount:  2,
				ItemCount: 2,
				ItemWidth: 22,
				startAddr: 3,
				startBit:  20,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayContinuous(test.count, test.startAddr, test.startBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayMultiplePacked(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 1, 32,
			ArrayMultiple{
				regCount:    1,
				ItemCount:   1,
				ItemWidth:   32,
				ItemsPerReg: 1,
				startAddr:   0,
				startBit:    0,
			},
		},
		{1, 4, 8,
			ArrayMultiple{
				regCount:    1,
				ItemCount:   4,
				ItemWidth:   8,
				ItemsPerReg: 4,
				startAddr:   1,
				startBit:    0,
			},
		},
		{2, 3, 16,
			ArrayMultiple{
				regCount:    2,
				ItemCount:   3,
				ItemWidth:   16,
				ItemsPerReg: 2,
				startAddr:   2,
				startBit:    0,
			},
		},
		{3, 4, 4,
			ArrayMultiple{
				regCount:    1,
				ItemCount:   4,
				ItemWidth:   4,
				ItemsPerReg: 8,
				startAddr:   3,
				startBit:    0,
			},
		},
		{4, 5, 8,
			ArrayMultiple{
				regCount:    2,
				ItemCount:   5,
				ItemWidth:   8,
				ItemsPerReg: 4,
				startAddr:   4,
				startBit:    0,
			},
		},
		{5, 10, 7,
			ArrayMultiple{
				regCount:    3,
				ItemCount:   10,
				ItemWidth:   7,
				ItemsPerReg: 4,
				startAddr:   5,
				startBit:    0,
			},
		},
		{6, 50, 3,
			ArrayMultiple{
				regCount:    5,
				ItemCount:   50,
				ItemWidth:   3,
				ItemsPerReg: 10,
				startAddr:   6,
				startBit:    0,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayMultiplePacked(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
