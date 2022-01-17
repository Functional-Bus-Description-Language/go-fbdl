package fbdl

import (
	"encoding/json"
	"fmt"
	"math"
)

type AccessMask struct {
	Upper, Lower int64
}

func (am AccessMask) Width() int64 { return am.Upper - am.Lower + 1 }

type Access interface {
	RegCount() int64 // RegCount returns the number of occupied registers.
	IsArray() bool
	StartAddr() int64
	EndAddr() int64
	EndBit() int64
	Width() int64 // Width returns total width of single element.
}

// AccessSingleSingle describes access to ...
type AccessSingleSingle struct {
	Addr int64
	Mask AccessMask
}

func (ass AccessSingleSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy string
		Count    int64
		Addr     int64
		Mask     AccessMask
	}{
		Strategy: "Single",
		Count:    1,
		Addr:     ass.Addr,
		Mask:     ass.Mask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (ass AccessSingleSingle) RegCount() int64  { return 1 }
func (ass AccessSingleSingle) IsArray() bool    { return false }
func (ass AccessSingleSingle) StartAddr() int64 { return ass.Addr }
func (ass AccessSingleSingle) EndAddr() int64   { return ass.Addr }
func (ass AccessSingleSingle) EndBit() int64    { return ass.Mask.Upper }
func (ass AccessSingleSingle) Width() int64     { return ass.Mask.Width() }

func makeAccessSingleSingle(addr, startBit, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make AccessSingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return AccessSingleSingle{
		Addr: addr,
		Mask: AccessMask{Upper: startBit + width - 1, Lower: startBit},
	}
}

// AccessSingleSingle describes access to ...
type AccessSingleContinuous struct {
	regCount int64

	startAddr int64      // Address of the first register.
	StartMask AccessMask // Mask for the first register.
	EndMask   AccessMask // Mask for the last register.
}

func (asc AccessSingleContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		StartAddr int64
		StartMask AccessMask
		EndMask   AccessMask
	}{
		Strategy:  "Continuous",
		RegCount:  asc.regCount,
		StartAddr: asc.startAddr,
		StartMask: asc.StartMask,
		EndMask:   asc.EndMask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (asc AccessSingleContinuous) RegCount() int64  { return asc.regCount }
func (asc AccessSingleContinuous) IsArray() bool    { return false }
func (asc AccessSingleContinuous) StartAddr() int64 { return asc.startAddr }
func (asc AccessSingleContinuous) EndAddr() int64   { return asc.startAddr + asc.regCount - 1 }
func (asc AccessSingleContinuous) EndBit() int64    { return asc.EndMask.Upper }

func (asc AccessSingleContinuous) Width() int64 {
	w := asc.StartMask.Width() + asc.EndMask.Width()
	if asc.regCount > 2 {
		w += busWidth * (asc.regCount - 2)
	}
	return w
}

// IsEndRegWider returns true if end register is wider than the start one.
func (asc AccessSingleContinuous) IsEndMaskWider() bool {
	if asc.EndMask.Width() > asc.StartMask.Width() {
		return true
	}
	return false
}

func makeAccessSingleContinuous(addr, startBit, width int64) Access {
	startMask := AccessMask{Upper: busWidth - 1, Lower: startBit}
	regCount := int64(1)

	var endMask AccessMask
	w := busWidth - startBit
	for {
		regCount += 1
		if w+busWidth < width {
			w += busWidth
		} else {
			endMask = AccessMask{Upper: (width - w) - 1, Lower: 0}
			break
		}
	}

	return AccessSingleContinuous{
		regCount:  regCount,
		startAddr: addr,
		StartMask: startMask,
		EndMask:   endMask,
	}
}

// makeAccessSingle makes AccessSingleSingle or AccessSingleContinuous depending on the argument values.
func makeAccessSingle(addr, startBit, width int64) Access {
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return makeAccessSingleSingle(addr, startBit, width)
	} else {
		return makeAccessSingleContinuous(addr, startBit, width)
	}
}

type AccessArraySingle struct {
	regCount int64

	startAddr int64
	Mask      AccessMask
}

func (aas AccessArraySingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		StartAddr int64
		Mask      AccessMask
	}{
		Strategy:  "Single",
		RegCount:  aas.regCount,
		StartAddr: aas.startAddr,
		Mask:      aas.Mask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (aas AccessArraySingle) RegCount() int64  { return aas.regCount }
func (aas AccessArraySingle) IsArray() bool    { return true }
func (aas AccessArraySingle) StartAddr() int64 { return aas.startAddr }
func (aas AccessArraySingle) EndAddr() int64   { return aas.startAddr + aas.regCount - 1 }
func (aas AccessArraySingle) EndBit() int64    { return aas.Mask.Upper }
func (aas AccessArraySingle) Width() int64     { return aas.Mask.Width() }

func makeAccessArraySingle(itemCount, addr, startBit, width int64) AccessArraySingle {
	if startBit+width > busWidth {
		msg := `cannot make AccessArraySingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return AccessArraySingle{
		regCount:  itemCount,
		startAddr: addr,
		Mask:      AccessMask{Upper: startBit + width - 1, Lower: startBit},
	}
}

type AccessArrayContinuous struct {
	regCount int64

	ItemCount int64
	ItemWidth int64
	startAddr int64
	StartBit  int64
}

func (aac AccessArrayContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		RegCount  int64
		ItemCount int64
		ItemWidth int64
		StartAddr int64
		StartBit  int64
	}{
		Strategy:  "Continuous",
		RegCount:  aac.regCount,
		ItemCount: aac.ItemCount,
		ItemWidth: aac.ItemWidth,
		StartAddr: aac.startAddr,
		StartBit:  aac.StartBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (aac AccessArrayContinuous) RegCount() int64  { return aac.regCount }
func (aac AccessArrayContinuous) IsArray() bool    { return true }
func (aac AccessArrayContinuous) StartAddr() int64 { return aac.startAddr }
func (aac AccessArrayContinuous) EndAddr() int64   { return aac.startAddr + aac.regCount - 1 }
func (aac AccessArrayContinuous) Width() int64     { return aac.ItemWidth }

func (aac AccessArrayContinuous) EndBit() int64 {
	return ((aac.StartBit + aac.regCount*aac.ItemWidth - 1) % busWidth)
}

func makeAccessArrayContinuous(itemCount, startAddr, startBit, width int64) Access {
	aac := AccessArrayContinuous{
		ItemCount: itemCount,
		ItemWidth: width,
		startAddr: startAddr,
		StartBit:  startBit,
	}

	totalWidth := itemCount * width
	firstRegWidth := busWidth - startBit

	aac.regCount = int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1

	return aac
}

type AccessArrayMultiple struct {
	regCount int64

	ItemCount      int64
	ItemWidth      int64
	ItemsPerAccess int64
	startAddr      int64
	StartBit       int64
}

func (aam AccessArrayMultiple) MarshalJSON() ([]byte, error) {
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
		RegCount:       aam.regCount,
		ItemCount:      aam.ItemCount,
		ItemWidth:      aam.ItemWidth,
		ItemsPerAccess: aam.ItemsPerAccess,
		StartAddr:      aam.startAddr,
		StartBit:       aam.StartBit,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (aam AccessArrayMultiple) RegCount() int64  { return aam.regCount }
func (aam AccessArrayMultiple) IsArray() bool    { return true }
func (aam AccessArrayMultiple) StartAddr() int64 { return aam.startAddr }
func (aam AccessArrayMultiple) EndAddr() int64   { return aam.startAddr + aam.regCount - 1 }
func (aam AccessArrayMultiple) Width() int64     { return aam.ItemWidth }

func (aam AccessArrayMultiple) EndBit() int64 {
	if aam.regCount == 1 {
		return aam.StartBit + aam.ItemCount*aam.ItemWidth - 1
	} else if aam.ItemCount%aam.ItemsPerAccess == 0 {
		return aam.StartBit + aam.ItemsPerAccess*aam.ItemWidth - 1
	} else {
		itemsInLast := aam.ItemCount % aam.ItemsPerAccess
		return aam.StartBit + itemsInLast*aam.ItemWidth - 1
	}
}

func (aam AccessArrayMultiple) ItemsInLastReg() int64 {
	return aam.ItemCount % aam.ItemsPerAccess
}

// makeAccessArrayMultiplePacked makes AccessArrayMultiple starting from bit 0,
// and placing as many items within single register as possible.
func makeAccessArrayMultiplePacked(itemCount, startAddr, width int64) Access {
	aam := AccessArrayMultiple{
		ItemCount:      itemCount,
		ItemWidth:      width,
		ItemsPerAccess: busWidth / width,
		startAddr:      startAddr,
		StartBit:       0,
	}

	if itemCount <= aam.ItemsPerAccess {
		aam.regCount = 1
	} else {
		aam.regCount = int64(math.Ceil(float64(itemCount) / float64(aam.ItemsPerAccess)))
	}

	return aam
}
