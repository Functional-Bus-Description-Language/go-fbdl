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
	Strategy  string
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
		msg := `cannot make ArrayOneReg, startBit + (width * itemCount) > busWidth, (%d + (%d * %d) > %d)`
		panic(fmt.Sprintf(msg, startBit, width, itemCount, busWidth))
	}

	return ArrayOneReg{
		Strategy:  "ArrayOneReg",
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
	Strategy  string
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
		Strategy:  "ArrayOneInReg",
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
	Strategy  string
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
		Strategy:  "ArrayNRegs",
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

// ArrayMultiple describes an access to an array of functionalities
// with multiple functionalities placed within single register.
type ArrayMultiple struct {
	regCount int64

	ItemCount   int64
	ItemWidth   int64
	ItemsPerReg int64
	startAddr   int64
	startBit    int64
}

func (am ArrayMultiple) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy    string
		RegCount    int64
		ItemCount   int64
		ItemWidth   int64
		ItemsPerReg int64
		StartAddr   int64
		StartBit    int64
	}{
		Strategy:    "Multiple",
		RegCount:    am.regCount,
		ItemCount:   am.ItemCount,
		ItemWidth:   am.ItemWidth,
		ItemsPerReg: am.ItemsPerReg,
		StartAddr:   am.startAddr,
		StartBit:    am.startBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (am ArrayMultiple) GetRegCount() int64  { return am.regCount }
func (am ArrayMultiple) GetStartAddr() int64 { return am.startAddr }
func (am ArrayMultiple) GetEndAddr() int64   { return am.startAddr + am.regCount - 1 }
func (am ArrayMultiple) GetWidth() int64     { return am.ItemWidth }
func (am ArrayMultiple) GetStartBit() int64  { return am.startBit }

func (am ArrayMultiple) GetStartRegWidth() int64 {
	if am.ItemCount < am.ItemsPerReg {
		return am.ItemCount * am.ItemWidth
	}
	return am.ItemsPerReg * am.ItemWidth
}

func (am ArrayMultiple) GetEndRegWidth() int64 {
	itemsInEndReg := am.ItemCount % am.ItemsPerReg
	if itemsInEndReg == 0 {
		itemsInEndReg = am.ItemsPerReg
	}
	return itemsInEndReg * am.ItemWidth
}

func (am ArrayMultiple) GetEndBit() int64 {
	if am.regCount == 1 {
		return am.startBit + am.ItemCount*am.ItemWidth - 1
	} else if am.ItemCount%am.ItemsPerReg == 0 {
		return am.startBit + am.ItemsPerReg*am.ItemWidth - 1
	} else {
		itemsInLast := am.ItemCount % am.ItemsPerReg
		return am.startBit + itemsInLast*am.ItemWidth - 1
	}
}

func (am ArrayMultiple) ItemsInLastReg() int64 {
	inLastReg := am.ItemCount % am.ItemsPerReg
	if inLastReg == 0 {
		inLastReg = am.ItemsPerReg
	}
	return inLastReg
}

// MakeArrayMultiplePacked makes ArrayMultiple starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayMultiplePacked(itemCount, startAddr, width int64) Access {
	am := ArrayMultiple{
		ItemCount:   itemCount,
		ItemWidth:   width,
		ItemsPerReg: busWidth / width,
		startAddr:   startAddr,
		startBit:    0,
	}

	if itemCount <= am.ItemsPerReg {
		am.regCount = 1
	} else {
		am.regCount = int64(math.Ceil(float64(itemCount) / float64(am.ItemsPerReg)))
	}

	return am
}
