package constContainer

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
)

// HasConst returns true if container already has constant with given name.
func HasConst(c cnst.Container, name string) bool {
	if _, ok := c.Bools[name]; ok {
		return true
	}
	if _, ok := c.BoolLists[name]; ok {
		return true
	}
	if _, ok := c.Ints[name]; ok {
		return true
	}
	if _, ok := c.IntLists[name]; ok {
		return true
	}
	if _, ok := c.Strings[name]; ok {
		return true
	}

	return false
}

func AddConst(c *cnst.Container, name string, v types.Value) {
	switch v.(type) {
	case types.BitStr:
		panic("not yet implemented")
	case types.Bool:
		addBoolConst(c, name, v)
	case types.Float:
		addFloatConst(c, name, v)
	case types.Int:
		addIntConst(c, name, v)
	case types.List:
		switch v.(types.List)[0].(type) {
		case types.BitStr:
			panic("not yet implemented")
		case types.Bool:
			addBoolListConst(c, name, v)
		case types.Int:
			addIntListConst(c, name, v)
		case types.Str:
			panic("not yet implemented")
		default:
			panic("should never happen")
		}
	case types.Str:
		addStrConst(c, name, v)
	default:
		panic("should never happen")
	}
}

func addBoolConst(c *cnst.Container, name string, v types.Value) {
	b := bool(v.(types.Bool))
	if c.Bools == nil {
		c.Bools = map[string]bool{name: b}
	}
	c.Bools[name] = b
}

func addFloatConst(c *cnst.Container, name string, v types.Value) {
	f := float64(v.(types.Float))
	if c.Floats == nil {
		c.Floats = map[string]float64{name: f}
	}
	c.Floats[name] = f
}

func addBoolListConst(c *cnst.Container, name string, v types.Value) {
	l := constifyBoolList(v.(types.List))
	if l == nil {
		return
	}

	if c.BoolLists == nil {
		c.BoolLists = map[string][]bool{name: l}
	}
	c.BoolLists[name] = l
}

func addIntConst(c *cnst.Container, name string, v types.Value) {
	i := int64(v.(types.Int))
	if c.Ints == nil {
		c.Ints = map[string]int64{name: i}
	}
	c.Ints[name] = i
}

func addIntListConst(c *cnst.Container, name string, v types.Value) {
	l := constifyIntList(v.(types.List))
	if l == nil {
		return
	}

	if c.IntLists == nil {
		c.IntLists = map[string][]int64{name: l}
	}
	c.IntLists[name] = l
}

func addStrConst(c *cnst.Container, name string, v types.Value) {
	s := string(v.(types.Str))
	if c.Strings == nil {
		c.Strings = map[string]string{name: s}
	}
	c.Strings[name] = s
}

// constifyBoolList tries to constify list as an bool list.
// If any elemnt is of different type than types.Bool, then it returns nil.
func constifyBoolList(l types.List) []bool {
	bools := []bool{}

	for _, v := range l {
		if i, ok := v.(types.Bool); ok {
			bools = append(bools, bool(i))
		} else {
			return nil
		}
	}

	return bools
}

// constifyIntList tries to constify list as an int list.
// If any elemnt is of different type than types.Int, then it returns nil.
func constifyIntList(l types.List) []int64 {
	ints := []int64{}

	for _, v := range l {
		if i, ok := v.(types.Int); ok {
			ints = append(ints, int64(i))
		} else {
			return nil
		}
	}

	return ints
}
