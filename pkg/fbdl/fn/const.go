package fn

type ConstContainer struct {
	BoolConsts     map[string]bool
	BoolListConsts map[string][]bool
	FloatConsts    map[string]float64
	IntConsts      map[string]int64
	IntListConsts  map[string][]int64
	StrConsts      map[string]string
}

// Empty returns true if ConstContainer holds no constants.
func (c *ConstContainer) Empty() bool {
	if len(c.BoolConsts) != 0 || len(c.BoolListConsts) != 0 {
		return false
	}
	if len(c.FloatConsts) != 0 {
		return false
	}
	if len(c.IntConsts) != 0 || len(c.IntListConsts) != 0 {
		return false
	}
	if len(c.StrConsts) != 0 {
		return false
	}

	return true
}
