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
	LastBitPos() int64
}

// AccessSingleSingle describes access to ...
type AccessSingleSingle struct {
	Address int64
	Mask    Mask
}

func (ass AccessSingleSingle) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy string
		Count    int64
		Address  int64
		Mask     Mask
	}{
		Strategy: "Single",
		Count:    1,
		Address:  ass.Address,
		Mask:     ass.Mask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (ass AccessSingleSingle) Count() int64      { return 1 }
func (ass AccessSingleSingle) IsArray() bool     { return false }
func (ass AccessSingleSingle) LastBitPos() int64 { return ass.Mask.Upper }

func makeAccessSingleSingle(addr int64, startBit int64, width int64) Access {
	if startBit+width > busWidth {
		msg := `cannot make AccessSingleSingle, startBit + width > busWidth, (%d + %d > %d)`
		panic(fmt.Sprintf(msg, startBit, width, busWidth))
	}

	return AccessSingleSingle{
		Address: addr,
		Mask:    Mask{Upper: startBit + width - 1, Lower: startBit},
	}
}

// AccessSingleSingle describes access to ...
type AccessSingleContinuous struct {
	count int64 // count is the number of occupied registers.

	StartAddress int64 // Address of the first register.
	StartMask    Mask  // Mask for the first register.
	EndMask      Mask  // Mask for the last register.
}

func (asc AccessSingleContinuous) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(struct {
		Strategy     string
		Count        int64
		StartAddress int64
		StartMask    Mask
		EndMask      Mask
	}{
		Strategy:     "Continuous",
		Count:        asc.count,
		StartAddress: asc.StartAddress,
		StartMask:    asc.StartMask,
		EndMask:      asc.EndMask,
	})

	if err != nil {
		return nil, err
	}

	return j, nil
}

func (asc AccessSingleContinuous) Count() int64      { return asc.count }
func (asc AccessSingleContinuous) IsArray() bool     { return false }
func (asc AccessSingleContinuous) LastBitPos() int64 { return asc.EndMask.Upper }

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
		count:        count,
		StartAddress: addr,
		StartMask:    startMask,
		EndMask:      endMask,
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

type AccessArray struct {
	Strategy string

	Address int64
	count   int64 // count is the number of occupied registers.

	AccessesPerItem int64
	ItemsPerAccess  int64

	BunchSize        int64
	AccessesPerBunch int64

	Mask Mask
}

func (aa *AccessArray) Count() int64      { return aa.count }
func (aa *AccessArray) IsArray() bool     { return true }
func (aa *AccessArray) LastBitPos() int64 { return 1 } // FIXME

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
