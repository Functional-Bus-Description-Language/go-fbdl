package val

import "math"

// Range represents possible value range.
// The range can consist of multiple subranges.
// For example, [1, 3, 8, 10] means that the value can equal 1, 2, 3, 8, 9 or 10.
type Range []int64

func (r Range) IsEmpty() bool {
	return len(r) == 0
}

// Width returns bit width needed to represent the range.
func (r Range) Width() int64 {
	max := int64(0)
	for _, b := range r {
		if b > max {
			max = b
		}
	}
	return int64(math.Ceil(math.Log2(float64(max + 1))))
}
