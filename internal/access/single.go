package access

import (
	"bytes"
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
)

type ss struct {
	Strategy string
	Addr     int64
	Mask     Mask
}

type SingleSingle struct {
	ss
}

func (ss SingleSingle) RegCount() int64 { return 1 }
func (ss SingleSingle) Addr() int64     { return ss.ss.Addr }
func (ss SingleSingle) Mask() Mask      { return ss.ss.Mask }
func (ss SingleSingle) Width() int64    { return ss.ss.Mask.Width() }

func (ss SingleSingle) StartAddr() int64 { return ss.Addr() }
func (ss SingleSingle) EndAddr() int64   { return ss.Addr() }
func (ss SingleSingle) EndBit() int64    { return ss.Mask().End() }

func (ss SingleSingle) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, ss.Addr())
	hash.Write(&buf, ss.Mask())
	return hash.Hash(buf)
}

func MakeSingleSingle(addr, startBit, width int64) SingleSingle {
	if startBit+width > busWidth {
		msg := `cannot make SingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return SingleSingle{
		ss: ss{
			Strategy: "Single",
			Addr:     addr,
			Mask:     makeMask(startBit, startBit+width-1),
		},
	}
}

type sc struct {
	Strategy  string
	StartAddr int64
	EndAddr   int64
	StartMask Mask
	EndMask   Mask
}

type SingleContinuous struct {
	sc
}

func (sc SingleContinuous) RegCount() int64  { return sc.sc.EndAddr - sc.sc.StartAddr + 1 }
func (sc SingleContinuous) StartAddr() int64 { return sc.sc.StartAddr }
func (sc SingleContinuous) EndAddr() int64   { return sc.sc.EndAddr }
func (sc SingleContinuous) StartMask() Mask  { return sc.sc.StartMask }
func (sc SingleContinuous) EndMask() Mask    { return sc.sc.EndMask }
func (sc SingleContinuous) Width() int64 {
	w := sc.StartMask().Width() + sc.EndMask().Width()
	w += (sc.EndAddr() - sc.StartAddr() - 1) * busWidth
	return w
}

func (sc SingleContinuous) EndBit() int64 { return sc.EndMask().End() }

func (sc SingleContinuous) Hash() uint32 {
	buf := bytes.Buffer{}

	hash.Write(&buf, sc.StartAddr())
	hash.Write(&buf, sc.EndAddr())

	hash.Write(&buf, sc.StartMask())
	hash.Write(&buf, sc.EndMask())

	return hash.Hash(buf)
}

func MakeSingleContinuous(addr, startBit, width int64) SingleContinuous {
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
		sc: sc{
			Strategy:  "Continuous",
			StartAddr: addr,
			EndAddr:   addr + regCount + 1,
			StartMask: makeMask(startBit, busWidth-1),
			EndMask:   makeMask(0, endBit),
		},
	}
}

// MakeSingle makes SingleSingle or SingleContinuous depending on the argument values.
func MakeSingle(addr, startBit, width int64) Single {
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return MakeSingleSingle(addr, startBit, width)
	} else {
		return MakeSingleContinuous(addr, startBit, width)
	}
}
