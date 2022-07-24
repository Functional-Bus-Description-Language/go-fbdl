package access

import (
	"encoding/json"
	"fmt"
	"math"
)

// ArraySingle describes an access to an array of elements with single element placed within single register.
type ArraySingle struct {
	regCount int64

	startAddr int64
	startBit  int64
	endBit    int64
}

func (as ArraySingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		StartAddr int64
		StartBit  int64
		EndBit    int64
	}{
		Strategy:  "Single",
		RegCount:  as.regCount,
		StartAddr: as.startAddr,
		StartBit:  as.startBit,
		EndBit:    as.endBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (as ArraySingle) RegCount() int64  { return as.regCount }
func (as ArraySingle) StartAddr() int64 { return as.startAddr }
func (as ArraySingle) EndAddr() int64   { return as.startAddr + as.regCount - 1 }
func (as ArraySingle) StartBit() int64  { return as.startBit }
func (as ArraySingle) EndBit() int64    { return as.endBit }
func (as ArraySingle) Width() int64     { return as.endBit - as.startBit + 1 }

func MakeArraySingle(itemCount, addr, startBit, width int64) ArraySingle {
	if startBit+width > busWidth {
		msg := `cannot make ArraySingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return ArraySingle{
		regCount:  itemCount,
		startAddr: addr,
		startBit:  startBit,
		endBit:    startBit + width - 1,
	}
}

// ArrayContinuous describes an access to an array of elements with single element placed within multiple continuous registers.
type ArrayContinuous struct {
	regCount int64

	ItemCount int64
	ItemWidth int64
	startAddr int64
	startBit  int64
}

func (ac ArrayContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		ItemCount int64
		ItemWidth int64
		StartAddr int64
		StartBit  int64
	}{
		Strategy:  "Continuous",
		RegCount:  ac.regCount,
		ItemCount: ac.ItemCount,
		ItemWidth: ac.ItemWidth,
		StartAddr: ac.startAddr,
		StartBit:  ac.startBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (ac ArrayContinuous) RegCount() int64  { return ac.regCount }
func (ac ArrayContinuous) StartAddr() int64 { return ac.startAddr }
func (ac ArrayContinuous) EndAddr() int64   { return ac.startAddr + ac.regCount - 1 }
func (ac ArrayContinuous) Width() int64     { return ac.ItemWidth }
func (ac ArrayContinuous) StartBit() int64  { return ac.startBit }

func (ac ArrayContinuous) EndBit() int64 {
	return ((ac.startBit + ac.regCount*ac.ItemWidth - 1) % busWidth)
}

func MakeArrayContinuous(itemCount, startAddr, startBit, width int64) Access {
	ac := ArrayContinuous{
		ItemCount: itemCount,
		ItemWidth: width,
		startAddr: startAddr,
		startBit:  startBit,
	}

	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit

	ac.regCount = int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1

	return ac
}

// ArrayMultiple describes an access to an array of elements with multiple elements placed within single register.
type ArrayMultiple struct {
	regCount int64

	ItemCount      int64
	ItemWidth      int64
	ItemsPerAccess int64
	startAddr      int64
	startBit       int64
}

func (am ArrayMultiple) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy       string
		RegCount       int64
		ItemCount      int64
		ItemWidth      int64
		ItemsPerAccess int64
		StartAddr      int64
		StartBit       int64
	}{
		Strategy:       "Multiple",
		RegCount:       am.regCount,
		ItemCount:      am.ItemCount,
		ItemWidth:      am.ItemWidth,
		ItemsPerAccess: am.ItemsPerAccess,
		StartAddr:      am.startAddr,
		StartBit:       am.startBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (am ArrayMultiple) RegCount() int64  { return am.regCount }
func (am ArrayMultiple) StartAddr() int64 { return am.startAddr }
func (am ArrayMultiple) EndAddr() int64   { return am.startAddr + am.regCount - 1 }
func (am ArrayMultiple) Width() int64     { return am.ItemWidth }
func (am ArrayMultiple) StartBit() int64  { return am.startBit }

func (am ArrayMultiple) EndBit() int64 {
	if am.regCount == 1 {
		return am.startBit + am.ItemCount*am.ItemWidth - 1
	} else if am.ItemCount%am.ItemsPerAccess == 0 {
		return am.startBit + am.ItemsPerAccess*am.ItemWidth - 1
	} else {
		itemsInLast := am.ItemCount % am.ItemsPerAccess
		return am.startBit + itemsInLast*am.ItemWidth - 1
	}
}

func (am ArrayMultiple) ItemsInLastReg() int64 {
	return am.ItemCount % am.ItemsPerAccess
}

// MakeArrayMultiplePacked makes ArrayMultiple starting from bit 0,
// and placing as many items within single register as possible.
func MakeArrayMultiplePacked(itemCount, startAddr, width int64) Access {
	am := ArrayMultiple{
		ItemCount:      itemCount,
		ItemWidth:      width,
		ItemsPerAccess: busWidth / width,
		startAddr:      startAddr,
		startBit:       0,
	}

	if itemCount <= am.ItemsPerAccess {
		am.regCount = 1
	} else {
		am.regCount = int64(math.Ceil(float64(itemCount) / float64(am.ItemsPerAccess)))
	}

	return am
}
