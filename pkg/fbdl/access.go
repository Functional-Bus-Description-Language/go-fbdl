package fbdl

import (
	"math"
)

type Mask struct {
	Upper, Lower int64
}

type Access interface {
	// Count returns the number of occupied registers.
	Count() int64
	IsArray() bool
}

type AccessSingle struct {
	Strategy string

	Address int64 // Address is the base address - address of the first register.
	count   int64 // count is the number of occupied registers.
	Mask    Mask
}

func (as *AccessSingle) Count() int64 { return as.count }

func (as *AccessSingle) IsArray() bool { return false }

func makeAccessSingle(baseAddr int64, baseBit int64, width int64) *AccessSingle {
	as := AccessSingle{
		Address: baseAddr,
		count:   int64(math.Ceil(float64(width) / float64(busWidth))),
	}

	if width > busWidth {
		as.Strategy = "Linear"
		as.Mask = Mask{Upper: (width - 1) % busWidth, Lower: 0}
	} else {
		as.Strategy = "Single"
		as.Mask = Mask{Upper: width - 1, Lower: 0}
	}

	return &as
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

func (aa *AccessArray) Count() int64 { return aa.count }

func (aa *AccessArray) IsArray() bool { return true }

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
