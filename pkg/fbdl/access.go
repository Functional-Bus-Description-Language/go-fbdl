package fbdl

import (
	"math"
)

type Mask struct {
	Upper, Lower uint
}

type Access interface {
	Count() uint
	IsArray() bool
}

type AccessSingle struct {
	Strategy string

	Address uint
	count   uint // count is the number of occupied registers.
	Mask    Mask
}

func (as *AccessSingle) Count() uint { return as.count }

func (as *AccessSingle) IsArray() bool { return false }

func MakeAccessSingle(baseAddr uint, width uint) *AccessSingle {
	as := AccessSingle{
		Address: baseAddr,
		count:   uint(math.Ceil(float64(width) / float64(busWidth))),
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

	Address uint
	count   uint // count is the number of occupied registers.

	AccessesPerItem uint
	ItemsPerAccess  uint

	BunchSize        uint
	AccessesPerBunch uint

	Mask Mask
}

func (aa *AccessArray) Count() uint { return aa.count }

func (aa *AccessArray) IsArray() bool { return true }

func MakeAccessArray(count uint, baseAddr uint, width uint) *AccessArray {
	aa := AccessArray{
		Address:         baseAddr,
		AccessesPerItem: uint(math.Ceil(float64(width) / float64(busWidth))),
		ItemsPerAccess:  busWidth / width,
	}

	if aa.AccessesPerItem == 1 && aa.ItemsPerAccess == 1 {
		aa.Strategy = "Single"
		aa.count = count
		aa.Mask = Mask{Upper: width - 1, Lower: 0}
	} else if aa.AccessesPerItem == 1 && aa.ItemsPerAccess > 1 {
		aa.Strategy = "Multiple"
		aa.count = uint(math.Ceil(float64(count) / float64(aa.ItemsPerAccess)))
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
