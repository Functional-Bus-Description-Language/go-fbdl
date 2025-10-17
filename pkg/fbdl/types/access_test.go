package types

import (
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
			Access{
				Type:          "SingleOneReg",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     1,
				ItemWidth:     1,
				StartAddr:     0,
				EndAddr:       0,
				StartBit:      0,
				EndBit:        0,
				StartRegWidth: 1,
				EndRegWidth:   1,
			},
		},
		{1, 31, 2,
			Access{
				Type:          "SingleNRegs",
				RegCount:      2,
				RegWidth:      32,
				ItemCount:     1,
				ItemWidth:     2,
				StartAddr:     1,
				EndAddr:       2,
				StartBit:      31,
				EndBit:        0,
				StartRegWidth: 1,
				EndRegWidth:   1,
			},
		},
		{2, 30, 57,
			Access{
				Type:          "SingleNRegs",
				RegCount:      3,
				RegWidth:      32,
				ItemCount:     1,
				ItemWidth:     57,
				StartAddr:     2,
				EndAddr:       4,
				StartBit:      30,
				EndBit:        22,
				StartRegWidth: 2,
				EndRegWidth:   23,
			},
		},
		{3, 0, 32,
			Access{
				Type:          "SingleOneReg",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     1,
				ItemWidth:     32,
				StartAddr:     3,
				EndAddr:       3,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
	}

	for i, test := range tests {
		got := MakeSingleAccess(test.baseAddr, test.baseBit, test.width)

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
			Access{
				Type:          "ArrayNRegs",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     1,
				ItemWidth:     32,
				StartAddr:     0,
				EndAddr:       0,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
		{1, 4, 0, 5,
			Access{
				Type:          "ArrayNRegs",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     4,
				ItemWidth:     5,
				StartAddr:     1,
				EndAddr:       1,
				StartBit:      0,
				EndBit:        19,
				StartRegWidth: 20,
				EndRegWidth:   20,
			},
		},
		{2, 2, 20, 23,
			Access{
				Type:          "ArrayNRegs",
				RegCount:      3,
				RegWidth:      32,
				ItemCount:     2,
				ItemWidth:     23,
				StartAddr:     2,
				EndAddr:       4,
				StartBit:      20,
				EndBit:        1,
				StartRegWidth: 12,
				EndRegWidth:   2,
			},
		},
		{3, 2, 20, 22,
			Access{
				Type:          "ArrayNRegs",
				RegCount:      2,
				RegWidth:      32,
				ItemCount:     2,
				ItemWidth:     22,
				StartAddr:     3,
				EndAddr:       4,
				StartBit:      20,
				EndBit:        31,
				StartRegWidth: 12,
				EndRegWidth:   32,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNRegsAccess(test.count, test.startAddr, test.startBit, test.width)

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
			Access{
				Type:          "ArrayNInReg",
				RegCount:      2,
				RegWidth:      32,
				ItemCount:     4,
				ItemWidth:     16,
				StartAddr:     2,
				EndAddr:       3,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
		{4, 8, 8,
			Access{
				Type:          "ArrayNInReg",
				RegCount:      2,
				RegWidth:      32,
				ItemCount:     8,
				ItemWidth:     8,
				StartAddr:     4,
				EndAddr:       5,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
		{5, 12, 7,
			Access{
				Type:          "ArrayNInReg",
				RegCount:      3,
				RegWidth:      32,
				ItemCount:     12,
				ItemWidth:     7,
				StartAddr:     5,
				EndAddr:       7,
				StartBit:      0,
				EndBit:        27,
				StartRegWidth: 28,
				EndRegWidth:   28,
			},
		},
		{6, 50, 3,
			Access{
				Type:          "ArrayNInReg",
				RegCount:      5,
				RegWidth:      32,
				ItemCount:     50,
				ItemWidth:     3,
				StartAddr:     6,
				EndAddr:       10,
				StartBit:      0,
				EndBit:        29,
				StartRegWidth: 30,
				EndRegWidth:   30,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNInRegAccess(test.count, test.startAddr, test.width)

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
			Access{
				Type:          "ArrayNInRegMInEndReg",
				RegCount:      2,
				RegWidth:      32,
				ItemCount:     5,
				ItemWidth:     7,
				StartAddr:     0,
				EndAddr:       1,
				StartBit:      0,
				EndBit:        6,
				StartRegWidth: 28,
				EndRegWidth:   7,
			},
		},
		{1, 66, 1,
			Access{
				Type:          "ArrayNInRegMInEndReg",
				RegCount:      3,
				RegWidth:      32,
				ItemCount:     66,
				ItemWidth:     1,
				StartAddr:     1,
				EndAddr:       3,
				StartBit:      0,
				EndBit:        1,
				StartRegWidth: 32,
				EndRegWidth:   2,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayNInRegMInEndRegAccess(test.count, test.startAddr, test.width)

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayOneInNRegs(t *testing.T) {
	var tests = []struct {
		startAddr int64
		count     int64
		width     int64
		want      Access
	}{
		{0, 2, 33,
			Access{
				Type:          "ArrayOneInNRegs",
				RegCount:      4,
				RegWidth:      32,
				ItemCount:     2,
				ItemWidth:     33,
				StartAddr:     0,
				EndAddr:       3,
				StartBit:      0,
				EndBit:        0,
				StartRegWidth: 32,
				EndRegWidth:   1,
			},
		},
		{1, 3, 64,
			Access{
				Type:          "ArrayOneInNRegs",
				RegCount:      6,
				RegWidth:      32,
				ItemCount:     3,
				ItemWidth:     64,
				StartAddr:     1,
				EndAddr:       6,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayOneInNRegsAccess(test.count, test.startAddr, test.width)

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayOneReg(t *testing.T) {
	var tests = []struct {
		itemCount int64
		addr      int64
		startBit  int64
		width     int64
		want      Access
	}{
		{7, 2, 0, 3,
			Access{
				Type:          "ArrayOneReg",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     7,
				ItemWidth:     3,
				StartAddr:     2,
				EndAddr:       2,
				StartBit:      0,
				EndBit:        20,
				StartRegWidth: 21,
				EndRegWidth:   21,
			},
		},
		{4, 5, 0, 8,
			Access{
				Type:          "ArrayOneReg",
				RegCount:      1,
				RegWidth:      32,
				ItemCount:     4,
				ItemWidth:     8,
				StartAddr:     5,
				EndAddr:       5,
				StartBit:      0,
				EndBit:        31,
				StartRegWidth: 32,
				EndRegWidth:   32,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayOneRegAccess(
			test.itemCount, test.addr, test.startBit, test.width,
		)

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}

func TestMakeArrayOneInReg(t *testing.T) {
	var tests = []struct {
		itemCount int64
		addr      int64
		startBit  int64
		width     int64
		want      Access
	}{
		{7, 0, 0, 30,
			Access{
				Type:          "ArrayOneInReg",
				RegCount:      7,
				RegWidth:      32,
				ItemCount:     7,
				ItemWidth:     30,
				StartAddr:     0,
				EndAddr:       6,
				StartBit:      0,
				EndBit:        29,
				StartRegWidth: 30,
				EndRegWidth:   30,
			},
		},
		{3, 45, 5, 21,
			Access{
				Type:          "ArrayOneInReg",
				RegCount:      3,
				RegWidth:      32,
				ItemCount:     3,
				ItemWidth:     21,
				StartAddr:     45,
				EndAddr:       47,
				StartBit:      5,
				EndBit:        25,
				StartRegWidth: 21,
				EndRegWidth:   21,
			},
		},
	}

	for i, test := range tests {
		got := MakeArrayOneInRegAccess(
			test.itemCount, test.addr, test.startBit, test.width,
		)

		if got != test.want {
			t.Errorf("[%d] got %v, want %v", i, got, test.want)
		}
	}
}
