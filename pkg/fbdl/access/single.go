package access

import (
	"encoding/json"
	"fmt"
)

// SingleOneReg describes an access to a single functionality placed within single register.
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

func (sc SingleContinuous) GetRegCount() int64      { return sc.regCount }
func (sc SingleContinuous) GetStartAddr() int64     { return sc.startAddr }
func (sc SingleContinuous) GetEndAddr() int64       { return sc.startAddr + sc.regCount - 1 }
func (sc SingleContinuous) GetStartBit() int64      { return sc.startBit }
func (sc SingleContinuous) GetEndBit() int64        { return sc.endBit }
func (sc SingleContinuous) GetStartRegWidth() int64 { return busWidth - sc.startBit }
func (sc SingleContinuous) GetEndRegWidth() int64   { return sc.endBit + 1 }

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

// MakeSingle makes SingleOneReg or SingleContinuous depending on the argument values.
func MakeSingle(addr, startBit, width int64) Access {
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return MakeSingleOneReg(addr, startBit, width)
	} else {
		return MakeSingleContinuous(addr, startBit, width)
	}
}
