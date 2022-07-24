package access

import (
	"encoding/json"
	"fmt"
)

// SingleSingle describes an access to a single element placed within single register.
type SingleSingle struct {
	Addr int64
	Mask Mask
}

func (ss SingleSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy string
		Addr     int64
		Mask     Mask
	}{
		Strategy: "Single",
		Addr:     ss.Addr,
		Mask:     ss.Mask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (ss SingleSingle) RegCount() int64  { return 1 }
func (ss SingleSingle) StartAddr() int64 { return ss.Addr }
func (ss SingleSingle) EndAddr() int64   { return ss.Addr }
func (ss SingleSingle) EndBit() int64    { return ss.Mask.Upper }
func (ss SingleSingle) Width() int64     { return ss.Mask.Width() }

func MakeSingleSingle(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make SingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return SingleSingle{
		Addr: addr,
		Mask: Mask{Upper: startBit + width - 1, Lower: startBit},
	}
}

// SingleContinuous describes an access to a single element placed within multiple continuous registers.
type SingleContinuous struct {
	regCount int64

	startAddr int64 // Address of the first register.
	StartMask Mask  // Mask for the first register.
	EndMask   Mask  // Mask for the last register.
}

func (sc SingleContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		StartAddr int64
		StartMask Mask
		EndMask   Mask
	}{
		Strategy:  "Continuous",
		RegCount:  sc.regCount,
		StartAddr: sc.startAddr,
		StartMask: sc.StartMask,
		EndMask:   sc.EndMask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (sc SingleContinuous) RegCount() int64  { return sc.regCount }
func (sc SingleContinuous) StartAddr() int64 { return sc.startAddr }
func (sc SingleContinuous) EndAddr() int64   { return sc.startAddr + sc.regCount - 1 }
func (sc SingleContinuous) EndBit() int64    { return sc.EndMask.Upper }

func (sc SingleContinuous) Width() int64 {
	w := sc.StartMask.Width() + sc.EndMask.Width()
	if sc.regCount > 2 {
		w += busWidth * (sc.regCount - 2)
	}
	return w
}

// IsEndRegWider returns true if end register is wider than the start one.
func (sc SingleContinuous) IsEndMaskWider() bool {
	if sc.EndMask.Width() > sc.StartMask.Width() {
		return true
	}
	return false
}

func MakeSingleContinuous(addr, startBit, width int64) Access {
	startMask := Mask{Upper: busWidth - 1, Lower: startBit}
	regCount := int64(1)

	var endMask Mask
	w := busWidth - startBit
	for {
		regCount += 1
		if w+busWidth < width {
			w += busWidth
		} else {
			endMask = Mask{Upper: (width - w) - 1, Lower: 0}
			break
		}
	}

	return SingleContinuous{
		regCount:  regCount,
		startAddr: addr,
		StartMask: startMask,
		EndMask:   endMask,
	}
}

// MakeSingle makes SingleSingle or SingleContinuous depending on the argument values.
func MakeSingle(addr, startBit, width int64) Access {
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return MakeSingleSingle(addr, startBit, width)
	} else {
		return MakeSingleContinuous(addr, startBit, width)
	}
}
