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
func MakeArrayOneReg(itemCount, addr, startBit, width int64) Access {
	if startBit+(width*itemCount) > busWidth {
		panic(
			fmt.Sprintf(
				"cannot make ArrayOneReg, startBit + (width * itemCount) > busWidth, (%d + (%d * %d) > %d)",
				startBit, width, itemCount, busWidth,
			),
		)
	}

	return Access{
		Type:          "ArrayOneReg",
		RegCount:      1,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     addr,
		EndAddr:       addr,
		StartBit:      startBit,
		EndBit:        startBit + itemCount*width - 1,
		StartRegWidth: itemCount * width,
		EndRegWidth:   itemCount * width,
	}
}

// ArrayOneInReg describes an access to an array of functionalities
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
func MakeArrayOneInReg(itemCount, addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make ArrayOneInReg, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return Access{
		Type:          "ArrayOneInReg",
		RegCount:      itemCount,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     addr,
		EndAddr:       addr + itemCount - 1,
		StartBit:      startBit,
		EndBit:        startBit + width - 1,
		StartRegWidth: width,
		EndRegWidth:   width,
	}
}

// ArrayNRegs describes an access to an array of functionalities
// with single functionality placed continuously within registers.
//
//	Example:
//
//	p [4]param; width = 14
//
//	           Reg N                        Reg N+1
//	--------------------------- ---------------------------------
//	|| p[0] | p[1] | p[2](0) || || p[2](1) | p[3] | 8 bits gap ||
//	--------------------------- ---------------------------------
func MakeArrayNRegs(itemCount, startAddr, startBit, width int64) Access {
	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit

	regCount := int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1
	endBit := (startBit + itemCount*width - 1) % busWidth

	startRegWidth := busWidth - startBit
	if regCount == 1 {
		startRegWidth = itemCount*width - startBit
	}

	return Access{
		Type:          "ArrayNRegs",
		RegCount:      regCount,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     startAddr,
		EndAddr:       startAddr + regCount - 1,
		StartBit:      startBit,
		EndBit:        endBit,
		StartRegWidth: startRegWidth,
		EndRegWidth:   endBit + 1,
	}
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
//
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

	regCount := itemCount / itemsInReg

	return Access{
		Type:          "ArrayNInReg",
		RegCount:      regCount,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     startAddr,
		EndAddr:       startAddr + regCount - 1,
		StartBit:      0,
		EndBit:        itemsInReg*width - 1,
		StartRegWidth: width * itemsInReg,
		EndRegWidth:   width * itemsInReg,
	}
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
//
// MakeArrayNInRegMInEndReg makes ArrayNInRegMInEndReg starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayNInRegMInEndReg(itemCount, startAddr, width int64) Access {
	itemsInReg := busWidth / width
	itemsInEndReg := itemCount % itemsInReg

	if itemsInEndReg == 0 {
		panic("itemsInEndReg = 0, use ArrayNInReg")
	}

	regCount := int64(math.Ceil(float64(itemCount) / float64(itemsInReg)))

	return Access{
		Type:          "ArrayNInRegMInEndReg",
		RegCount:      regCount,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     startAddr,
		EndAddr:       startAddr + regCount - 1,
		StartBit:      0,
		EndBit:        itemsInEndReg*width - 1,
		StartRegWidth: width * itemsInReg,
		EndRegWidth:   width * itemsInEndReg,
	}
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
//
// MakeArrayOneInNRegs makes ArrayNInRegMInEndReg starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayOneInNRegs(itemCount, startAddr, width int64) Access {
	if width <= busWidth {
		panic(fmt.Sprintf("width <= busWidth, %d <= %d", width, busWidth))
	}

	regsPerItem := width / busWidth
	if width%busWidth != 0 {
		regsPerItem++
	}
	regCount := itemCount * regsPerItem

	endBit := width - (width/busWidth)*busWidth - 1
	if width%busWidth == 0 {
		endBit = busWidth - 1
	}

	return Access{
		Type:          "ArrayOneInNRegs",
		RegCount:      regCount,
		RegWidth:      busWidth,
		ItemCount:     itemCount,
		ItemWidth:     width,
		StartAddr:     startAddr,
		EndAddr:       startAddr + regCount - 1,
		StartBit:      0,
		EndBit:        endBit,
		StartRegWidth: busWidth,
		EndRegWidth:   endBit + 1,
	}
}
