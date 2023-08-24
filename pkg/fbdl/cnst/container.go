package cnst

type Container struct {
	Bools     map[string]bool
	BoolLists map[string][]bool
	Floats    map[string]float64
	Ints      map[string]int64
	IntLists  map[string][]int64
	Strings   map[string]string
}

// Empty returns true if Container holds no constants.
func (c *Container) Empty() bool {
	if len(c.Bools) != 0 || len(c.BoolLists) != 0 {
		return false
	}
	if len(c.Floats) != 0 {
		return false
	}
	if len(c.Ints) != 0 || len(c.IntLists) != 0 {
		return false
	}
	if len(c.Strings) != 0 {
		return false
	}

	return true
}
