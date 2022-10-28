package access

import (
	"bytes"
	"fmt"
	"math"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/hash"
)

type as struct {
	Strategy  string
	StartAddr int64
	EndAddr   int64
	Mask      Mask
	ItemCount int64
	ItemWidth int64
}

type ArraySingle struct {
	as
}

func (as ArraySingle) RegCount() int64  { return as.as.EndAddr - as.as.StartAddr + 1 }
func (as ArraySingle) StartAddr() int64 { return as.as.StartAddr }
func (as ArraySingle) EndAddr() int64   { return as.as.EndAddr }
func (as ArraySingle) Mask() Mask       { return as.as.Mask }
func (as ArraySingle) ItemCount() int64 { return as.as.ItemCount }
func (as ArraySingle) ItemWidth() int64 { return as.as.ItemWidth }

func (as ArraySingle) EndBit() int64 { return as.Mask().End() }

func (as ArraySingle) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, as.StartAddr())
	hash.Write(&buf, as.EndAddr())
	hash.Write(&buf, as.Mask())
	hash.Write(&buf, as.ItemCount())
	hash.Write(&buf, as.ItemWidth())
	return hash.Hash(buf)
}

func MakeArraySingle(itemCount, addr, startBit, width int64) ArraySingle {
	if startBit+width > busWidth {
		msg := `cannot make ArraySingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return ArraySingle{
		as: as{
			Strategy:  "Single",
			StartAddr: addr,
			EndAddr:   addr + itemCount - 1,
			Mask:      makeMask(startBit, startBit+width-1),
			ItemCount: itemCount,
			ItemWidth: width,
		},
	}
}

type ac struct {
	Strategy  string
	StartAddr int64
	EndAddr   int64
	StartMask Mask
	EndMask   Mask
	ItemCount int64
	ItemWidth int64
}

type ArrayContinuous struct {
	ac
}

func (ac ArrayContinuous) RegCount() int64  { return ac.ac.StartAddr - ac.ac.EndAddr + 1 }
func (ac ArrayContinuous) StartAddr() int64 { return ac.ac.StartAddr }
func (ac ArrayContinuous) EndAddr() int64   { return ac.ac.EndAddr }
func (ac ArrayContinuous) StartMask() Mask  { return ac.ac.StartMask }
func (ac ArrayContinuous) EndMask() Mask    { return ac.ac.EndMask }
func (ac ArrayContinuous) ItemCount() int64 { return ac.ac.ItemCount }
func (ac ArrayContinuous) ItemWidth() int64 { return ac.ac.ItemWidth }

func (ac ArrayContinuous) EndBit() int64 { return ac.EndMask().End() }

func (ac ArrayContinuous) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, ac.StartAddr())
	hash.Write(&buf, ac.EndAddr())
	hash.Write(&buf, ac.StartMask())
	hash.Write(&buf, ac.EndMask())
	hash.Write(&buf, ac.ItemCount())
	hash.Write(&buf, ac.ItemWidth())
	return hash.Hash(buf)
}

func MakeArrayContinuous(itemCount, startAddr, startBit, width int64) ArrayContinuous {
	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit
	regCount := int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1
	endBit := (startBit + regCount*width - 1) % busWidth

	ac := ArrayContinuous{
		ac: ac{
			Strategy:  "Continuous",
			StartAddr: startAddr,
			EndAddr:   startAddr + regCount - 1,
			StartMask: makeMask(startBit, busWidth-1),
			EndMask:   makeMask(0, endBit),
			ItemCount: itemCount,
			ItemWidth: width,
		},
	}

	return ac
}

type am struct {
	Strategy       string
	StartAddr      int64
	EndAddr        int64
	StartMask      Mask
	EndMask        Mask
	ItemCount      int64
	ItemWidth      int64
	ItemsPerAccess int64
}

type ArrayMultiple struct {
	am
}

func (am ArrayMultiple) RegCount() int64       { return am.am.StartAddr - am.am.EndAddr + 1 }
func (am ArrayMultiple) StartAddr() int64      { return am.am.StartAddr }
func (am ArrayMultiple) EndAddr() int64        { return am.am.EndAddr }
func (am ArrayMultiple) StartMask() Mask       { return am.am.StartMask }
func (am ArrayMultiple) EndMask() Mask         { return am.am.EndMask }
func (am ArrayMultiple) ItemCount() int64      { return am.am.ItemCount }
func (am ArrayMultiple) ItemWidth() int64      { return am.am.ItemWidth }
func (am ArrayMultiple) ItemsPerAccess() int64 { return am.am.ItemsPerAccess }

func (am ArrayMultiple) EndBit() int64 { return am.EndMask().End() }

func (am ArrayMultiple) Hash() uint32 {
	buf := bytes.Buffer{}
	hash.Write(&buf, am.StartAddr())
	hash.Write(&buf, am.EndAddr())
	hash.Write(&buf, am.StartMask())
	hash.Write(&buf, am.EndMask())
	hash.Write(&buf, am.ItemCount())
	hash.Write(&buf, am.ItemWidth())
	hash.Write(&buf, am.ItemsPerAccess())
	return hash.Hash(buf)
}

// MakeArrayMultiplePacked makes ArrayMultiple starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayMultiplePacked(itemCount, startAddr, width int64) ArrayMultiple {
	itemsPerAccess := busWidth / width
	regCount := int64(1)
	if itemCount > itemsPerAccess {
		regCount = int64(math.Ceil(float64(itemCount) / float64(itemsPerAccess)))
	}

	var endBit int64
	if regCount == 1 {
		endBit = itemCount*width - 1
	} else if itemCount%itemsPerAccess == 0 {
		endBit = itemsPerAccess*width - 1
	} else {
		itemsInLast := itemCount % itemsPerAccess
		endBit = itemsInLast*width - 1
	}

	am := ArrayMultiple{
		am: am{
			Strategy:       "Multiple",
			StartAddr:      startAddr,
			EndAddr:        startAddr + regCount - 1,
			StartMask:      makeMask(0, width-1),
			EndMask:        makeMask(endBit-width+1, endBit),
			ItemCount:      itemCount,
			ItemWidth:      width,
			ItemsPerAccess: itemsPerAccess,
		},
	}

	return am
}
