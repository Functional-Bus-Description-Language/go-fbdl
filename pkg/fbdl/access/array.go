package access

import (
	"encoding/json"
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
	typ       string
	addr      int64
	startBit  int64
	itemWidth int64
	itemCount int64
}

func (aor ArrayOneReg) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string
		Addr      int64
		StartBit  int64
		ItemWidth int64
		ItemCount int64
	}{
		Type:      aor.typ,
		Addr:      aor.addr,
		StartBit:  aor.startBit,
		ItemWidth: aor.itemWidth,
		ItemCount: aor.itemCount,
	})
}

func (aor ArrayOneReg) GetRegCount() int64      { return 1 }
func (aor ArrayOneReg) GetStartAddr() int64     { return aor.addr }
func (aor ArrayOneReg) GetEndAddr() int64       { return aor.addr }
func (aor ArrayOneReg) GetStartBit() int64      { return aor.startBit }
func (aor ArrayOneReg) GetEndBit() int64        { return aor.startBit*aor.itemCount*aor.itemWidth - 1 }
func (aor ArrayOneReg) GetWidth() int64         { return aor.itemWidth }
func (aor ArrayOneReg) GetStartRegWidth() int64 { return aor.itemCount * aor.itemWidth }
func (aor ArrayOneReg) GetEndRegWidth() int64   { return aor.itemCount * aor.itemWidth }

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
		typ:       "ArrayOneReg",
		addr:      addr,
		startBit:  startBit,
		itemCount: itemCount,
		itemWidth: width,
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
	typ       string
	regCount  int64
	startAddr int64
	startBit  int64
	endBit    int64
}

func (aoir ArrayOneInReg) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string
		RegCount  int64
		StartAddr int64
		StartBit  int64
		EndBit    int64
	}{
		Type:      aoir.typ,
		RegCount:  aoir.regCount,
		StartAddr: aoir.startAddr,
		StartBit:  aoir.startBit,
		EndBit:    aoir.endBit,
	})
}

func (aoir ArrayOneInReg) GetRegCount() int64      { return aoir.regCount }
func (aoir ArrayOneInReg) GetStartAddr() int64     { return aoir.startAddr }
func (aoir ArrayOneInReg) GetEndAddr() int64       { return aoir.startAddr + aoir.regCount - 1 }
func (aoir ArrayOneInReg) GetStartBit() int64      { return aoir.startBit }
func (aoir ArrayOneInReg) GetEndBit() int64        { return aoir.endBit }
func (aoir ArrayOneInReg) GetWidth() int64         { return aoir.endBit - aoir.startBit + 1 }
func (aoir ArrayOneInReg) GetStartRegWidth() int64 { return aoir.GetWidth() }
func (aoir ArrayOneInReg) GetEndRegWidth() int64   { return aoir.GetWidth() }

func MakeArrayOneInReg(itemCount, addr, startBit, width int64) ArrayOneInReg {
	if startBit+width > busWidth {
		msg := `cannot make ArrayOneInReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return ArrayOneInReg{
		typ:       "ArrayOneInReg",
		regCount:  itemCount,
		startAddr: addr,
		startBit:  startBit,
		endBit:    startBit + width - 1,
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
	typ       string
	regCount  int64
	itemCount int64
	itemWidth int64
	startAddr int64
	startBit  int64
}

func (anr ArrayNRegs) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string
		RegCount  int64
		ItemCount int64
		ItemWidth int64
		StartAddr int64
		StartBit  int64
	}{
		Type:      anr.typ,
		RegCount:  anr.regCount,
		ItemCount: anr.itemCount,
		ItemWidth: anr.itemWidth,
		StartAddr: anr.startAddr,
		StartBit:  anr.startBit,
	})
}

func (anr ArrayNRegs) GetRegCount() int64      { return anr.regCount }
func (anr ArrayNRegs) GetStartAddr() int64     { return anr.startAddr }
func (anr ArrayNRegs) GetEndAddr() int64       { return anr.startAddr + anr.regCount - 1 }
func (anr ArrayNRegs) GetWidth() int64         { return anr.itemWidth }
func (anr ArrayNRegs) GetStartBit() int64      { return anr.startBit }
func (anr ArrayNRegs) GetStartRegWidth() int64 { return busWidth - anr.startBit }
func (anr ArrayNRegs) GetEndRegWidth() int64   { return anr.GetEndBit() + 1 }

func (anr ArrayNRegs) GetEndBit() int64 {
	return ((anr.startBit + anr.regCount*anr.itemWidth - 1) % busWidth)
}

func MakeArrayNRegs(itemCount, startAddr, startBit, width int64) Access {
	anr := ArrayNRegs{
		typ:       "ArrayNRegs",
		itemCount: itemCount,
		itemWidth: width,
		startAddr: startAddr,
		startBit:  startBit,
	}

	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit

	anr.regCount = int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1

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
	typ        string
	regCount   int64
	itemCount  int64
	itemWidth  int64
	itemsInReg int64
	startAddr  int64
	startBit   int64
}

func (anir ArrayNInReg) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type       string
		RegCount   int64
		ItemCount  int64
		ItemWidth  int64
		ItemsInReg int64
		StartAddr  int64
		StartBit   int64
	}{
		Type:       anir.typ,
		RegCount:   anir.regCount,
		ItemCount:  anir.itemCount,
		ItemWidth:  anir.itemWidth,
		ItemsInReg: anir.itemsInReg,
		StartAddr:  anir.startAddr,
		StartBit:   anir.startBit,
	})
}

func (anir ArrayNInReg) GetRegCount() int64      { return anir.regCount }
func (anir ArrayNInReg) GetStartAddr() int64     { return anir.startAddr }
func (anir ArrayNInReg) GetEndAddr() int64       { return anir.startAddr + anir.regCount - 1 }
func (anir ArrayNInReg) GetWidth() int64         { return anir.itemWidth }
func (anir ArrayNInReg) GetStartBit() int64      { return anir.startBit }
func (anir ArrayNInReg) GetEndBit() int64        { return anir.startBit + anir.itemsInReg*anir.itemWidth - 1 }
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
		typ:        "ArrayNInReg",
		regCount:   itemCount / itemsInReg,
		itemCount:  itemCount,
		itemWidth:  width,
		itemsInReg: itemsInReg,
		startAddr:  startAddr,
		startBit:   0,
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
	typ           string
	regCount      int64
	itemCount     int64
	itemWidth     int64
	itemsInReg    int64
	itemsInEndReg int64
	startAddr     int64
	startBit      int64
}

func (anm ArrayNInRegMInEndReg) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type          string
		RegCount      int64
		ItemCount     int64
		ItemWidth     int64
		ItemsInReg    int64
		ItemsInEndReg int64
		StartAddr     int64
		StartBit      int64
	}{
		Type:          anm.typ,
		RegCount:      anm.regCount,
		ItemCount:     anm.itemCount,
		ItemWidth:     anm.itemWidth,
		ItemsInReg:    anm.itemsInReg,
		ItemsInEndReg: anm.itemsInEndReg,
		StartAddr:     anm.startAddr,
		StartBit:      anm.startBit,
	})
}

func (anm ArrayNInRegMInEndReg) GetRegCount() int64      { return anm.regCount }
func (anm ArrayNInRegMInEndReg) GetStartAddr() int64     { return anm.startAddr }
func (anm ArrayNInRegMInEndReg) GetEndAddr() int64       { return anm.startAddr + anm.regCount - 1 }
func (anm ArrayNInRegMInEndReg) GetWidth() int64         { return anm.itemWidth }
func (anm ArrayNInRegMInEndReg) GetStartBit() int64      { return anm.startBit }
func (anm ArrayNInRegMInEndReg) GetStartRegWidth() int64 { return anm.itemsInReg * anm.itemWidth }
func (anm ArrayNInRegMInEndReg) GetEndRegWidth() int64   { return anm.itemsInEndReg * anm.itemWidth }
func (anm ArrayNInRegMInEndReg) GetEndBit() int64 {
	return anm.startBit + anm.itemsInEndReg*anm.itemWidth - 1
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
		typ:           "ArrayNInRegMInEndReg",
		regCount:      int64(math.Ceil(float64(itemCount) / float64(itemsInReg))),
		itemCount:     itemCount,
		itemWidth:     width,
		itemsInReg:    itemsInReg,
		itemsInEndReg: itemsInEndReg,
		startAddr:     startAddr,
		startBit:      0,
	}

	return anm
}

// ArrayOneInNRegs describes an access to an array of functionalities
// with one functionality placed in N registers.
// Start bit is always 0.
//
//	Example:
//
//	c [2]config; width = 33
//
//	    Reg N               Reg N+1              Reg N+2              Reg N+3
//	------------- --------------------------- ------------- ---------------------------
//	|| c[0](0) || || c[0](1) | 31 bits gap || || c[1](0) || || c[1](1) | 31 bits gap ||
//	------------- --------------------------- ------------- ---------------------------
type ArrayOneInNRegs struct {
	typ       string
	itemCount int64
	itemWidth int64
	startAddr int64
}

func (aoinr ArrayOneInNRegs) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string
		ItemCount int64
		ItemWidth int64
		StartAddr int64
	}{
		Type:      aoinr.typ,
		ItemCount: aoinr.itemCount,
		ItemWidth: aoinr.itemWidth,
		StartAddr: aoinr.startAddr,
	})
}

func (aoinr ArrayOneInNRegs) GetRegCount() int64 {
	if aoinr.itemWidth%busWidth == 0 {
		return aoinr.itemCount * aoinr.itemWidth / busWidth
	}
	return aoinr.itemCount * (aoinr.itemWidth/busWidth + 1)
}
func (aoinr ArrayOneInNRegs) GetStartAddr() int64     { return aoinr.startAddr }
func (aoinr ArrayOneInNRegs) GetEndAddr() int64       { return aoinr.startAddr + aoinr.GetRegCount() - 1 }
func (aoinr ArrayOneInNRegs) GetWidth() int64         { return aoinr.itemWidth }
func (aoinr ArrayOneInNRegs) GetStartBit() int64      { return 0 }
func (aoinr ArrayOneInNRegs) GetStartRegWidth() int64 { return busWidth }
func (aoinr ArrayOneInNRegs) GetEndRegWidth() int64   { return aoinr.GetEndBit() + 1 }
func (aoinr ArrayOneInNRegs) GetEndBit() int64 {
	if aoinr.itemWidth%busWidth == 0 {
		return busWidth - 1
	}
	return aoinr.itemWidth - (aoinr.itemWidth/busWidth)*busWidth - 1
}
func (aoinr ArrayOneInNRegs) GetRegsPerItem() int64 {
	if aoinr.itemWidth%busWidth == 0 {
		return aoinr.itemWidth / busWidth
	}
	return aoinr.itemWidth/busWidth + 1
}

// MakeArrayOneInNRegs makes ArrayNInRegMInEndReg starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayOneInNRegs(itemCount, startAddr, width int64) Access {
	if width <= busWidth {
		panic(fmt.Sprintf("width <= busWidth, %d <= %d", width, busWidth))
	}

	aoinr := ArrayOneInNRegs{
		typ:       "ArrayOneInNRegs",
		itemCount: itemCount,
		itemWidth: width,
		startAddr: startAddr,
	}

	return aoinr
}
