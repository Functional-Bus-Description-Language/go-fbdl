package value

import "math"

type Range interface {
	isRange()
	Width() int64
}

// SingleRange represents possible single value range.
type SingleRange struct {
	Left  int64 // Left bound
	Right int64 // Right bound
}

func (sr SingleRange) isRange() {}

// Width returns bit width required to represent the range.
func (sr SingleRange) Width() int64 {
	return int64(math.Ceil(math.Log2(float64(sr.Right + 1))))
}

// MultiRange represents possible multiple value ranges.
// For example, [1:3, 8:10] means that the value can equal 1, 2, 3, 8, 9 or 10.
type MultiRange []SingleRange

func (mr MultiRange) isRange() {}

// TODO: This function can potentially be removed.
func (ml MultiRange) IsEmpty() bool {
	return len(ml) == 0
}

// Width returns bit width required to represent the range list.
func (ml MultiRange) Width() int64 {
	max := int64(0)
	for _, r := range ml {
		if r.Width() > max {
			max = r.Width()
		}
	}
	return max
}
