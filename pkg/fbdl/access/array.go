package access

import (
	"fmt"
	"math"
)

// ArrayOneReg describes an access to an array of functionalities
// with all items placed in one register.
//
//	Example:
//
//	s [4]status; width = 7
//
//	                   Reg N
//	--------------------------------------------
//	|| s[0] | s[1] | s[2] | s[3] | 4 bits gap ||
//	--------------------------------------------
type ArrayOneReg struct {
	Type      string
	Addr      int64
	StartBit  int64
	ItemWidth int64
	ItemCount int64
}

func (aor ArrayOneReg) GetRegCount() int64      { return 1 }
func (aor ArrayOneReg) GetStartAddr() int64     { return aor.Addr }
func (aor ArrayOneReg) GetEndAddr() int64       { return aor.Addr }
func (aor ArrayOneReg) GetStartBit() int64      { return aor.StartBit }
func (aor ArrayOneReg) GetEndBit() int64        { return aor.StartBit*aor.ItemCount*aor.ItemWidth - 1 }
func (aor ArrayOneReg) GetWidth() int64         { return aor.ItemWidth }
func (aor ArrayOneReg) GetStartRegWidth() int64 { return aor.ItemCount * aor.ItemWidth }
func (aor ArrayOneReg) GetEndRegWidth() int64   { return aor.ItemCount * aor.ItemWidth }

func MakeArrayOneReg(itemCount, addr, startBit, width int64) ArrayOneReg {
	if startBit+(width*itemCount) > busWidth {
		panic(
			fmt.Sprintf(
				"cannot make ArrayOneReg, startBit + (width * itemCount) > busWidth, (%d + (%d * %d) > %d)",
				startBit, width, itemCount, busWidth,
			),
		)
	}

	return ArrayOneReg{
		Type:      "ArrayOneReg",
		Addr:      addr,
		StartBit:  startBit,
		ItemCount: itemCount,
		ItemWidth: width,
	}
}

// ArraySingle describes an access to an array of functionalities
// with single item placed within single register.
//
//	Example:
//
//	c [3]config; width = 25
//
//	         Reg N                  Reg N+1                Reg N+2
//	----------------------- ----------------------- -----------------------
//	|| c[0] | 7 bits gap || || c[1] | 7 bits gap || || c[2] | 7 bits gap ||
//	----------------------- ----------------------- -----------------------
type ArrayOneInReg struct {
	Type      string
	RegCount  int64
	StartAddr int64
	StartBit  int64
	EndBit    int64
}

func (aoir ArrayOneInReg) GetRegCount() int64      { return aoir.RegCount }
func (aoir ArrayOneInReg) GetStartAddr() int64     { return aoir.StartAddr }
func (aoir ArrayOneInReg) GetEndAddr() int64       { return aoir.StartAddr + aoir.RegCount - 1 }
func (aoir ArrayOneInReg) GetStartBit() int64      { return aoir.StartBit }
func (aoir ArrayOneInReg) GetEndBit() int64        { return aoir.EndBit }
func (aoir ArrayOneInReg) GetWidth() int64         { return aoir.EndBit - aoir.StartBit + 1 }
func (aoir ArrayOneInReg) GetStartRegWidth() int64 { return aoir.GetWidth() }
func (aoir ArrayOneInReg) GetEndRegWidth() int64   { return aoir.GetWidth() }

func MakeArrayOneInReg(itemCount, addr, startBit, width int64) ArrayOneInReg {
	if startBit+width > busWidth {
		msg := `cannot make ArrayOneInReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return ArrayOneInReg{
		Type:      "ArrayOneInReg",
		RegCount:  itemCount,
		StartAddr: addr,
		StartBit:  startBit,
		EndBit:    startBit + width - 1,
	}
}

// ArrayNRegs describes an access to an array of functionalities
// with single functionality placed within multiple continuous registers.
//
//	Example:
//
//	p [4]param; width = 14
//
//	           Reg N                        Reg N+1
//	--------------------------- ---------------------------------
//	|| p[0] | p[1] | p[2](0) || || p[2](1) | p[3] | 8 bits gap ||
//	--------------------------- ---------------------------------
type ArrayNRegs struct {
	Type      string
	RegCount  int64
	ItemCount int64
	ItemWidth int64
	StartAddr int64
	StartBit  int64
}

func (anr ArrayNRegs) GetRegCount() int64      { return anr.RegCount }
func (anr ArrayNRegs) GetStartAddr() int64     { return anr.StartAddr }
func (anr ArrayNRegs) GetEndAddr() int64       { return anr.StartAddr + anr.RegCount - 1 }
func (anr ArrayNRegs) GetWidth() int64         { return anr.ItemWidth }
func (anr ArrayNRegs) GetStartBit() int64      { return anr.StartBit }
func (anr ArrayNRegs) GetStartRegWidth() int64 { return busWidth - anr.StartBit }
func (anr ArrayNRegs) GetEndRegWidth() int64   { return anr.GetEndBit() + 1 }

func (anr ArrayNRegs) GetEndBit() int64 {
	return ((anr.StartBit + anr.RegCount*anr.ItemWidth - 1) % busWidth)
}

func MakeArrayNRegs(itemCount, startAddr, startBit, width int64) Access {
	anr := ArrayNRegs{
		Type:      "ArrayNRegs",
		ItemCount: itemCount,
		ItemWidth: width,
		StartAddr: startAddr,
		StartBit:  startBit,
	}

	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit

	anr.RegCount = int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1

	return anr
}

// ArrayNInReg describes an access to an array of functionalities
// with multiple functionalities placed within single register.
//
//	Example:
//
//	c [6]config; width = 15
//
//	            Reg N                         Reg N+1                        Reg N+2
//	------------------------------ ------------------------------ ------------------------------
//	|| c[0] | c[1] | 2 bits gap || || c[2] | c[3] | 2 bits gap || || c[4] | c[5] | 2 bits gap ||
//	------------------------------ ------------------------------ ------------------------------
type ArrayNInReg struct {
	Type       string
	RegCount   int64
	ItemCount  int64
	ItemWidth  int64
	ItemsInReg int64
	StartAddr  int64
	StartBit   int64
}

func (anir ArrayNInReg) GetRegCount() int64      { return anir.RegCount }
func (anir ArrayNInReg) GetStartAddr() int64     { return anir.StartAddr }
func (anir ArrayNInReg) GetEndAddr() int64       { return anir.StartAddr + anir.RegCount - 1 }
func (anir ArrayNInReg) GetWidth() int64         { return anir.ItemWidth }
func (anir ArrayNInReg) GetStartBit() int64      { return anir.StartBit }
func (anir ArrayNInReg) GetEndBit() int64        { return anir.StartBit + anir.ItemsInReg*anir.ItemWidth - 1 }
func (anir ArrayNInReg) GetStartRegWidth() int64 { return anir.GetWidth() }
func (anir ArrayNInReg) GetEndRegWidth() int64   { return anir.GetWidth() }

// MakeArrayNInReg makes ArrayNInReg starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayNInReg(itemCount, startAddr, width int64) Access {
	itemsInReg := busWidth / width

	if itemCount%itemsInReg != 0 {
		panic(
			fmt.Sprintf(
				"cannot make ArrayNInReg, itemCount %% itemsInReg != 0, %d %% %d != 0",
				itemCount, itemsInReg,
			),
		)
	}

	anir := ArrayNInReg{
		Type:       "ArrayNInReg",
		RegCount:   itemCount / itemsInReg,
		ItemCount:  itemCount,
		ItemWidth:  width,
		ItemsInReg: itemsInReg,
		StartAddr:  startAddr,
		StartBit:   0,
	}

	return anir
}

// ArrayNInRegMInEndReg describes an access to an array of functionalities
// with multiple functionalities placed within single register.
//
//	Example:
//
//	c [5]config; width = 15
//
//	            Reg N                         Reg N+1                     Reg N+2
//	------------------------------ ------------------------------ ------------------------
//	|| c[0] | c[1] | 2 bits gap || || c[2] | c[3] | 2 bits gap || || c[4] | 17 bits gap ||
//	------------------------------ ------------------------------ ------------------------
type ArrayNInRegMInEndReg struct {
	Type          string
	RegCount      int64
	ItemCount     int64
	ItemWidth     int64
	ItemsInReg    int64
	ItemsInEndReg int64
	StartAddr     int64
	StartBit      int64
}

func (anm ArrayNInRegMInEndReg) GetRegCount() int64      { return anm.RegCount }
func (anm ArrayNInRegMInEndReg) GetStartAddr() int64     { return anm.StartAddr }
func (anm ArrayNInRegMInEndReg) GetEndAddr() int64       { return anm.StartAddr + anm.RegCount - 1 }
func (anm ArrayNInRegMInEndReg) GetWidth() int64         { return anm.ItemWidth }
func (anm ArrayNInRegMInEndReg) GetStartBit() int64      { return anm.StartBit }
func (anm ArrayNInRegMInEndReg) GetStartRegWidth() int64 { return anm.ItemsInReg * anm.ItemWidth }
func (anm ArrayNInRegMInEndReg) GetEndRegWidth() int64   { return anm.ItemsInEndReg * anm.ItemWidth }
func (anm ArrayNInRegMInEndReg) GetEndBit() int64 {
	return anm.StartBit + anm.ItemsInEndReg*anm.ItemWidth - 1
}

// MakeArrayNInRegMInEndReg makes ArrayNInRegMInEndReg starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayNInRegMInEndReg(itemCount, startAddr, width int64) Access {
	itemsInReg := busWidth / width
	itemsInEndReg := itemCount % itemsInReg

	if itemsInEndReg == 0 {
		panic("itemsInEndReg = 0, use ArrayNInReg")
	}

	anm := ArrayNInRegMInEndReg{
		Type:          "ArrayNInRegMInEndReg",
		RegCount:      int64(math.Ceil(float64(itemCount) / float64(itemsInReg))),
		ItemCount:     itemCount,
		ItemWidth:     width,
		ItemsInReg:    itemsInReg,
		ItemsInEndReg: itemsInEndReg,
		StartAddr:     startAddr,
		StartBit:      0,
	}

	return anm
}
