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
				Mask: AccessMask{Upper: 0, Lower: 0},
			},
		},
		{1, 31, 2,
			AccessSingleContinuous{
				regCount:  2,
				startAddr: 1,
				StartMask: AccessMask{Upper: 31, Lower: 31},
				EndMask:   AccessMask{Upper: 0, Lower: 0},
			},
		},
		{2, 30, 57,
			AccessSingleContinuous{
				regCount:  3,
				startAddr: 2,
				StartMask: AccessMask{Upper: 31, Lower: 30},
				EndMask:   AccessMask{Upper: 22, Lower: 0},
			},
		},
		{3, 0, 32,
			AccessSingleSingle{
				Addr: 3,
				Mask: AccessMask{Upper: 31, Lower: 0},
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

func TestMakeAccessArrayContinuous(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		startBit  int64
		width     int64
		want      Access
	}{
		{0, 1, 0, 32,
			AccessArrayContinuous{
				regCount:  1,
				ItemCount: 1,
				ItemWidth: 32,
				startAddr: 0,
				StartBit:  0,
			},
		},
		{1, 4, 0, 5,
			AccessArrayContinuous{
				regCount:  1,
				ItemCount: 4,
				ItemWidth: 5,
				startAddr: 1,
				StartBit:  0,
			},
		},
		{2, 2, 20, 23,
			AccessArrayContinuous{
				regCount:  3,
				ItemCount: 2,
				ItemWidth: 23,
				startAddr: 2,
				StartBit:  20,
			},
		},
		{3, 2, 20, 22,
			AccessArrayContinuous{
				regCount:  2,
				ItemCount: 2,
				ItemWidth: 22,
				startAddr: 3,
				StartBit:  20,
			},
		},
	}

	for i, test := range tests {
		got := makeAccessArrayContinuous(test.count, test.startAddr, test.startBit, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeAccessArrayMultiplePacked(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 1, 32,
			AccessArrayMultiple{
				regCount:       1,
				ItemCount:      1,
				ItemWidth:      32,
				ItemsPerAccess: 1,
				startAddr:      0,
				StartBit:       0,
			},
		},
		{1, 4, 8,
			AccessArrayMultiple{
				regCount:       1,
				ItemCount:      4,
				ItemWidth:      8,
				ItemsPerAccess: 4,
				startAddr:      1,
				StartBit:       0,
			},
		},
		{2, 3, 16,
			AccessArrayMultiple{
				regCount:       2,
				ItemCount:      3,
				ItemWidth:      16,
				ItemsPerAccess: 2,
				startAddr:      2,
				StartBit:       0,
			},
		},
		{3, 4, 4,
			AccessArrayMultiple{
				regCount:       1,
				ItemCount:      4,
				ItemWidth:      4,
				ItemsPerAccess: 8,
				startAddr:      3,
				StartBit:       0,
			},
		},
		{4, 5, 8,
			AccessArrayMultiple{
				regCount:       2,
				ItemCount:      5,
				ItemWidth:      8,
				ItemsPerAccess: 4,
				startAddr:      4,
				StartBit:       0,
			},
		},
		{5, 10, 7,
			AccessArrayMultiple{
				regCount:       3,
				ItemCount:      10,
				ItemWidth:      7,
				ItemsPerAccess: 4,
				startAddr:      5,
				StartBit:       0,
			},
		},
		{6, 50, 3,
			AccessArrayMultiple{
				regCount:       5,
				ItemCount:      50,
				ItemWidth:      3,
				ItemsPerAccess: 10,
				startAddr:      6,
				StartBit:       0,
			},
		},
	}

	for i, test := range tests {
		got := makeAccessArrayMultiplePacked(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
