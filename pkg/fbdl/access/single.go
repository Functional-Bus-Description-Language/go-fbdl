package access

import (
	"encoding/json"
	"fmt"
)

// SingleSingle describes an access to a single functionality placed within single register.
type SingleSingle struct {
	Addr     int64
	startBit int64
	endBit   int64
}

func (ss SingleSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy string
		Addr     int64
		StartBit int64
		EndBit   int64
	}{
		Strategy: "Single",
		Addr:     ss.Addr,
		StartBit: ss.startBit,
		EndBit:   ss.endBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (ss SingleSingle) GetRegCount() int64   { return 1 }
func (ss SingleSingle) GetStartAddr() int64  { return ss.Addr }
func (ss SingleSingle) GetEndAddr() int64    { return ss.Addr }
func (ss SingleSingle) GetStartBit() int64   { return ss.startBit }
func (ss SingleSingle) GetEndBit() int64     { return ss.endBit }
func (ss SingleSingle) GetWidth() int64      { return ss.endBit - ss.startBit + 1 }
func (ss SingleSingle) StartRegWidth() int64 { return ss.GetWidth() }
func (ss SingleSingle) EndRegWidth() int64   { return ss.GetWidth() }

func MakeSingleSingle(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make SingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return SingleSingle{
		Addr:     addr,
		startBit: startBit,
		endBit:   startBit + width - 1,
	}
}

// SingleContinuous describes an access to a single functionality placed within multiple continuous registers.
type SingleContinuous struct {
	regCount int64

	startAddr int64 // Address of the first register.
	startBit  int64
	endBit    int64
}

func (sc SingleContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		StartAddr int64
		StartBit  int64
		EndBit    int64
	}{
		Strategy:  "Continuous",
		RegCount:  sc.regCount,
		StartAddr: sc.startAddr,
		StartBit:  sc.startBit,
		EndBit:    sc.endBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (sc SingleContinuous) GetRegCount() int64   { return sc.regCount }
func (sc SingleContinuous) GetStartAddr() int64  { return sc.startAddr }
func (sc SingleContinuous) GetEndAddr() int64    { return sc.startAddr + sc.regCount - 1 }
func (sc SingleContinuous) GetStartBit() int64   { return sc.startBit }
func (sc SingleContinuous) GetEndBit() int64     { return sc.endBit }
func (sc SingleContinuous) StartRegWidth() int64 { return busWidth - sc.startBit }
func (sc SingleContinuous) EndRegWidth() int64   { return sc.endBit + 1 }

func (sc SingleContinuous) GetWidth() int64 {
	w := busWidth - sc.startBit + sc.endBit + 1
	if sc.regCount > 2 {
		w += busWidth * (sc.regCount - 2)
	}
	return w
}

// IsEndRegWider returns true if end register is wider than the start one.
func (sc SingleContinuous) IsEndRegWider() bool {
	return sc.endBit > busWidth-sc.startBit
}

func MakeSingleContinuous(addr, startBit, width int64) Access {
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

	return SingleContinuous{
		regCount:  regCount,
		startAddr: addr,
		startBit:  startBit,
		endBit:    endBit,
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
