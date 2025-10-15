package access

import (
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
func MakeSingleOneReg(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make SingleOneReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return Access{
		Type:          "SingleOneReg",
		RegCount:      1,
		RegWidth:      busWidth,
		ItemCount:     1,
		ItemWidth:     width,
		StartAddr:     addr,
		EndAddr:       addr,
		StartBit:      startBit,
		EndBit:        startBit + width - 1,
		StartRegWidth: width,
		EndRegWidth:   width,
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

	return Access{
		Type:          "SingleNRegs",
		RegCount:      regCount,
		RegWidth:      busWidth,
		ItemCount:     1,
		ItemWidth:     width,
		StartAddr:     addr,
		EndAddr:       addr + regCount - 1,
		StartBit:      startBit,
		EndBit:        endBit,
		StartRegWidth: busWidth - startBit,
		EndRegWidth:   endBit + 1,
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
