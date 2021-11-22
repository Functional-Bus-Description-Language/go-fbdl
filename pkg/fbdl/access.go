package fbdl

import (
	"encoding/json"
	"fmt"
	"math"
)

type Mask struct {
	Upper, Lower int64
}

type Access interface {
	Count() int64 // Count returns the number of occupied registers.
	IsArray() bool
	StartAddr() int64
	EndAddr() int64
	EndBit() int64
}

// AccessSingleSingle describes access to ...
type AccessSingleSingle struct {
	Addr int64
	Mask Mask
}

func (ass AccessSingleSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy string
		Count    int64
		Addr     int64
		Mask     Mask
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

func (ass AccessSingleSingle) Count() int64     { return 1 }
func (ass AccessSingleSingle) IsArray() bool    { return false }
func (ass AccessSingleSingle) StartAddr() int64 { return ass.Addr }
func (ass AccessSingleSingle) EndAddr() int64   { return ass.Addr }
func (ass AccessSingleSingle) EndBit() int64    { return ass.Mask.Upper }

func makeAccessSingleSingle(addr int64, startBit int64, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make AccessSingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return AccessSingleSingle{
		Addr: addr,
		Mask: Mask{Upper: startBit + width - 1, Lower: startBit},
	}
}

// AccessSingleSingle describes access to ...
type AccessSingleContinuous struct {
	count int64 // count is the number of occupied registers.

	startAddr int64 // Address of the first register.
	StartMask Mask  // Mask for the first register.
	EndMask   Mask  // Mask for the last register.
}

func (asc AccessSingleContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		Count     int64
		StartAddr int64
		StartMask Mask
		EndMask   Mask
	}{
		Strategy:  "Continuous",
		Count:     asc.count,
		StartAddr: asc.startAddr,
		StartMask: asc.StartMask,
		EndMask:   asc.EndMask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (asc AccessSingleContinuous) Count() int64     { return asc.count }
func (asc AccessSingleContinuous) IsArray() bool    { return false }
func (asc AccessSingleContinuous) StartAddr() int64 { return asc.startAddr }
func (asc AccessSingleContinuous) EndAddr() int64   { return asc.startAddr + asc.count - 1 }
func (asc AccessSingleContinuous) EndBit() int64    { return asc.EndMask.Upper }

func makeAccessSingleContinuous(addr int64, startBit int64, width int64) Access {
	startMask := Mask{Upper: busWidth - 1, Lower: startBit}
	count := int64(1)

	var endMask Mask
	w := busWidth - startBit
	for {
		count += 1
		if w+busWidth < width {
			w += busWidth
		} else {
			endMask = Mask{Upper: (width - w) - 1, Lower: 0}
			break
		}
	}

	return AccessSingleContinuous{
		count:     count,
		startAddr: addr,
		StartMask: startMask,
		EndMask:   endMask,
	}
}

// makeAccessSingle makes AccessSingleSingle or AccessSingleContinuous depending on the argument values.
func makeAccessSingle(addr int64, startBit int64, width int64) Access {
	//remainder := width % busWidth
	firstRegRemainder := busWidth - startBit

	if width <= firstRegRemainder {
		return makeAccessSingleSingle(addr, startBit, width)
	} else {
		return makeAccessSingleContinuous(addr, startBit, width)
	}
}

type AccessArraySingle struct {
	count int64 // count is the number of occupied registers.

	startAddr int64
	Mask      Mask
}

func (aas AccessArraySingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		Count     int64
		StartAddr int64
		Mask      Mask
	}{
		Strategy:  "Single",
		Count:     aas.count,
		StartAddr: aas.startAddr,
		Mask:      aas.Mask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (aas AccessArraySingle) Count() int64     { return aas.count }
func (aas AccessArraySingle) IsArray() bool    { return true }
func (aas AccessArraySingle) StartAddr() int64 { return aas.startAddr }
func (aas AccessArraySingle) EndAddr() int64   { return aas.startAddr + aas.count - 1 }
func (aas AccessArraySingle) EndBit() int64    { return aas.Mask.Upper }

func makeAccessArraySingle(count int64, addr int64, startBit int64, width int64) AccessArraySingle {
	if startBit+width > busWidth {
		msg := `cannot make AccessArraySingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return AccessArraySingle{
		count:     count,
		startAddr: addr,
		Mask:      Mask{Upper: startBit + width - 1, Lower: startBit},
	}
}

type AccessArrayContinuous struct {
	count int64

	ItemCount int64
	ItemWidth int64
	startAddr int64
	StartBit  int64
}

func (aac AccessArrayContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		Count     int64
		ItemCount int64
		ItemWidth int64
		StartAddr int64
		StartBit  int64
	}{
		Strategy:  "Continuous",
		Count:     aac.count,
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

func (aac AccessArrayContinuous) Count() int64     { return aac.count }
func (aac AccessArrayContinuous) IsArray() bool    { return true }
func (aac AccessArrayContinuous) StartAddr() int64 { return aac.startAddr }
func (aac AccessArrayContinuous) EndAddr() int64   { return aac.startAddr + aac.count - 1 }

func (aac AccessArrayContinuous) EndBit() int64 {
	return ((aac.StartBit + aac.count*aac.ItemWidth - 1) % busWidth)
}

func makeAccessArrayContinuous(count int64, startAddr int64, startBit int64, width int64) Access {
	aac := AccessArrayContinuous{
		ItemCount: count,
		ItemWidth: width,
		startAddr: startAddr,
		StartBit:  startBit,
	}

	totalWidth := count * width
	firstRegWidth := busWidth - startBit

	aac.count = int64(math.Ceil((float64(totalWidth)-float64(firstRegWidth))/float64(busWidth))) + 1

	return aac
}

type AccessArrayMultiple struct {
	count int64

	ItemCount int64
	ItemWidth int64
	startAddr int64
}

func (aam AccessArrayMultiple) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy  string
		Count     int64
		ItemCount int64
		ItemWidth int64
		StartAddr int64
	}{
		Strategy:  "Multiple",
		Count:     aam.count,
		ItemCount: aam.ItemCount,
		ItemWidth: aam.ItemWidth,
		StartAddr: aam.startAddr,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (aam AccessArrayMultiple) Count() int64     { return aam.count }
func (aam AccessArrayMultiple) IsArray() bool    { return true }
func (aam AccessArrayMultiple) StartAddr() int64 { return aam.startAddr }
func (aam AccessArrayMultiple) EndAddr() int64   { return aam.startAddr + aam.count - 1 }

func (aam AccessArrayMultiple) EndBit() int64 {
	return ((aam.ItemCount*aam.ItemWidth - 1) % busWidth)
}

func makeAccessArrayMultiple(count int64, startAddr int64, width int64) Access {
	aam := AccessArrayMultiple{
		ItemCount: count,
		ItemWidth: width,
		startAddr: startAddr,
	}

	aam.count = int64(math.Ceil(float64(count) * float64(width) / float64(busWidth)))

	return aam
}

/*
func (aa *AccessArray) Count() int64      { return aa.count }
func (aa *AccessArray) IsArray() bool     { return true }
func (aa *AccessArray) EndBit() int64 { return 1 } // FIXME

func makeAccessArray(count int64, baseAddr int64, width int64) *AccessArray {
	aa := AccessArray{
		Address:         baseAddr,
		AccessesPerItem: int64(math.Ceil(float64(width) / float64(busWidth))),
		ItemsPerAccess:  busWidth / width,
	}

	if aa.AccessesPerItem == 1 && aa.ItemsPerAccess == 1 {
		aa.Strategy = "Single"
		aa.count = count
		aa.Mask = Mask{Upper: width - 1, Lower: 0}
	} else if aa.AccessesPerItem == 1 && aa.ItemsPerAccess > 1 {
		aa.Strategy = "Multiple"
		aa.count = int64(math.Ceil(float64(count) / float64(aa.ItemsPerAccess)))
	} else {
		aa.Strategy = "Bunch"
		// TODO: Calculate it correctly.
		aa.count = 0
		// Number of items in bunch.
		if (width % busWidth) == 0 {
			aa.BunchSize = 1
		} else {
			aa.BunchSize = busWidth / (width % busWidth)
		}
		// Number of accesses for bunch transfer.
		aa.AccessesPerBunch = aa.BunchSize*width/busWidth + 1
	}

	return &aa
}
*/
