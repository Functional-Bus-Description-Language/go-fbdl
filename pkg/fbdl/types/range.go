package types

import "math"

// The Range interface represents range type from the FBDL specification.
//
// Check SingleRange and ArrayRange structs documentation for more details.
type Range interface {
	isRange()
	BitWidth() int64 // Returns bit width required to represent the range.
}

// The SingleRange struct represents possible single value range.
//
// This type is also used internally to represents address space.
// It also might be useful in generators implementation to represent,
// for example, address ranges.
type SingleRange struct {
	Start int64 // Left bound
	End   int64 // Right bound
}

func (sr SingleRange) isRange() {}

func (sr SingleRange) BitWidth() int64 {
	return int64(math.Ceil(math.Log2(float64(sr.End + 1))))
}

func (sr SingleRange) Shift(offset int64) SingleRange {
	sr.Start += offset
	sr.End += offset
	return sr
}

// ArrayRange represents possible multiple value ranges.
// For example, [1:3, 8:10] means that the value can equal 1, 2, 3, 8, 9 or 10.
type ArrayRange []SingleRange

func (ar ArrayRange) isRange() {}

// TODO: This function can potentially be removed.
func (ar ArrayRange) IsEmpty() bool {
	return len(ar) == 0
}

func (ar ArrayRange) BitWidth() int64 {
	max := int64(0)
	for _, sr := range ar {
		if sr.BitWidth() > max {
			max = sr.BitWidth()
		}
	}
	return max
}
