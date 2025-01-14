package access

import (
	"encoding/json"
	"fmt"
)

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
	typ      string
	addr     int64
	startBit int64
	endBit   int64
}

func (sor SingleOneReg) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type     string
		Addr     int64
		StartBit int64
		EndBit   int64
	}{
		Type:     sor.typ,
		Addr:     sor.addr,
		StartBit: sor.startBit,
		EndBit:   sor.endBit,
	})
}

func (sor SingleOneReg) RegCount() int64      { return 1 }
func (sor SingleOneReg) StartAddr() int64     { return sor.addr }
func (sor SingleOneReg) EndAddr() int64       { return sor.addr }
func (sor SingleOneReg) StartBit() int64      { return sor.startBit }
func (sor SingleOneReg) EndBit() int64        { return sor.endBit }
func (sor SingleOneReg) Width() int64         { return sor.endBit - sor.startBit + 1 }
func (sor SingleOneReg) StartRegWidth() int64 { return sor.Width() }
func (sor SingleOneReg) EndRegWidth() int64   { return sor.Width() }

func MakeSingleOneReg(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make SingleOneReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return SingleOneReg{
		typ:      "SingleOneReg",
		addr:     addr,
		startBit: startBit,
		endBit:   startBit + width - 1,
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
	typ       string
	regCount  int64
	startAddr int64 // Address of the first register.
	startBit  int64
	endBit    int64
}

func (snr SingleNRegs) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type      string
		RegCount  int64
		StartAddr int64
		StartBit  int64
		EndBit    int64
	}{
		Type:      snr.typ,
		RegCount:  snr.regCount,
		StartAddr: snr.startAddr,
		StartBit:  snr.startBit,
		EndBit:    snr.endBit,
	})
}

func (snr SingleNRegs) RegCount() int64      { return snr.regCount }
func (snr SingleNRegs) StartAddr() int64     { return snr.startAddr }
func (snr SingleNRegs) EndAddr() int64       { return snr.startAddr + snr.regCount - 1 }
func (snr SingleNRegs) StartBit() int64      { return snr.startBit }
func (snr SingleNRegs) EndBit() int64        { return snr.endBit }
func (snr SingleNRegs) StartRegWidth() int64 { return busWidth - snr.startBit }
func (snr SingleNRegs) EndRegWidth() int64   { return snr.endBit + 1 }

func (snr SingleNRegs) Width() int64 {
	w := busWidth - snr.startBit + snr.endBit + 1
	if snr.regCount > 2 {
		w += busWidth * (snr.regCount - 2)
	}
	return w
}

// IsEndRegWider returns true if end register is wider than the start one.
func (snr SingleNRegs) IsEndRegWider() bool {
	return snr.endBit > busWidth-snr.startBit
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
		typ:       "SingleNRegs",
		regCount:  regCount,
		startAddr: addr,
		startBit:  startBit,
		endBit:    endBit,
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
