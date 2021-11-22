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
				startAddr: 1,
				StartMask: Mask{Upper: 31, Lower: 31},
				EndMask:   Mask{Upper: 0, Lower: 0},
			},
		},
		{2, 30, 57,
			AccessSingleContinuous{
				count:     3,
				startAddr: 2,
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
				count:     1,
				ItemCount: 1,
				ItemWidth: 32,
				startAddr: 0,
				StartBit:  0,
			},
		},
		{1, 4, 0, 5,
			AccessArrayContinuous{
				count:     1,
				ItemCount: 4,
				ItemWidth: 5,
				startAddr: 1,
				StartBit:  0,
			},
		},
		{2, 2, 20, 23,
			AccessArrayContinuous{
				count:     3,
				ItemCount: 2,
				ItemWidth: 23,
				startAddr: 2,
				StartBit:  20,
			},
		},
		{3, 2, 20, 22,
			AccessArrayContinuous{
				count:     2,
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

func TestMakeAccessArrayMultiple(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 1, 32,
			AccessArrayMultiple{
				count:     1,
				ItemCount: 1,
				ItemWidth: 32,
				startAddr: 0,
			},
		},
		{1, 4, 8,
			AccessArrayMultiple{
				count:     1,
				ItemCount: 4,
				ItemWidth: 8,
				startAddr: 1,
			},
		},
		{2, 3, 16,
			AccessArrayMultiple{
				count:     2,
				ItemCount: 3,
				ItemWidth: 16,
				startAddr: 2,
			},
		},
		{3, 4, 4,
			AccessArrayMultiple{
				count:     1,
				ItemCount: 4,
				ItemWidth: 4,
				startAddr: 3,
			},
		},
		{4, 5, 8,
			AccessArrayMultiple{
				count:     2,
				ItemCount: 5,
				ItemWidth: 8,
				startAddr: 4,
			},
		},
	}

	for i, test := range tests {
		got := makeAccessArrayMultiple(test.count, test.startAddr, test.width)

		if reflect.TypeOf(got) != reflect.TypeOf(test.want) {
			t.Errorf("[%d] invalid type, got %T, want %T", i, got, test.want)
		}

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
