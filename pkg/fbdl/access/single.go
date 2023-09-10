package access

import "fmt"

// SingleOneReg describes an access to a single functionality placed within single register.
//
//	Example:
//
//	s status; width = 23
//
//	       Reg N
//	--------------------
//	|| s | 9 bits gap ||
//	--------------------
type SingleOneReg struct {
	Strategy string
	Addr     int64
	StartBit int64
	EndBit   int64
}

func (sor SingleOneReg) GetRegCount() int64      { return 1 }
func (sor SingleOneReg) GetStartAddr() int64     { return sor.Addr }
func (sor SingleOneReg) GetEndAddr() int64       { return sor.Addr }
func (sor SingleOneReg) GetStartBit() int64      { return sor.StartBit }
func (sor SingleOneReg) GetEndBit() int64        { return sor.EndBit }
func (sor SingleOneReg) GetWidth() int64         { return sor.EndBit - sor.StartBit + 1 }
func (sor SingleOneReg) GetStartRegWidth() int64 { return sor.GetWidth() }
func (sor SingleOneReg) GetEndRegWidth() int64   { return sor.GetWidth() }

func MakeSingleOneReg(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make SingleOneReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return SingleOneReg{
		Strategy: "SingleOneReg",
		Addr:     addr,
		StartBit: startBit,
		EndBit:   startBit + width - 1,
	}
}

// SingleNRegs describes an access to a single functionality placed within multiple continuous registers.
//
//	Example:
//
//	c config; width = 72
//
//	  Reg N     Reg N+1           Reg N+2
//	---------- ---------- ------------------------
//	|| c(0) || || c(1) || || c(2) | 24 bits gap ||
//	---------- ---------- ------------------------
type SingleNRegs struct {
	Strategy  string
	RegCount  int64
	StartAddr int64 // Address of the first register.
	StartBit  int64
	EndBit    int64
}

func (snr SingleNRegs) GetRegCount() int64      { return snr.RegCount }
func (snr SingleNRegs) GetStartAddr() int64     { return snr.StartAddr }
func (snr SingleNRegs) GetEndAddr() int64       { return snr.StartAddr + snr.RegCount - 1 }
func (snr SingleNRegs) GetStartBit() int64      { return snr.StartBit }
func (snr SingleNRegs) GetEndBit() int64        { return snr.EndBit }
func (snr SingleNRegs) GetStartRegWidth() int64 { return busWidth - snr.StartBit }
func (snr SingleNRegs) GetEndRegWidth() int64   { return snr.EndBit + 1 }

func (snr SingleNRegs) GetWidth() int64 {
	w := busWidth - snr.StartBit + snr.EndBit + 1
	if snr.RegCount > 2 {
		w += busWidth * (snr.RegCount - 2)
	}
	return w
}

// IsEndRegWider returns true if end register is wider than the start one.
func (snr SingleNRegs) IsEndRegWider() bool {
	return snr.EndBit > busWidth-snr.StartBit
}

func MakeSingleNRegs(addr, startBit, width int64) Access {
	regCount := int64(1)

	endBit := int64(0)
	w := busWidth - startBit
	for {
		regCount += 1
		if w+busWidth < width {
			w += busWidth
		} else {
			endBit = width - w - 1
			break
		}
	}

	return SingleNRegs{
		Strategy:  "SingleNRegs",
		RegCount:  regCount,
		StartAddr: addr,
		StartBit:  startBit,
		EndBit:    endBit,
	}
}

// MakeSingle makes SingleOneReg or SingleNRegs depending on the argument values.
func MakeSingle(addr, startBit, width int64) Access {
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return MakeSingleOneReg(addr, startBit, width)
	} else {
		return MakeSingleNRegs(addr, startBit, width)
	}
}
