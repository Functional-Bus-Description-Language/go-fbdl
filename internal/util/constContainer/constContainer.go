package constContainer

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// HasConst returns true if container already has constant with given name.
func HasConst(c fn.ConstContainer, name string) bool {
	if _, ok := c.BoolConsts[name]; ok {
		return true
	}
	if _, ok := c.BoolListConsts[name]; ok {
		return true
	}
	if _, ok := c.IntConsts[name]; ok {
		return true
	}
	if _, ok := c.IntListConsts[name]; ok {
		return true
	}
	if _, ok := c.StrConsts[name]; ok {
		return true
	}

	return false
}

func AddConst(c *fn.ConstContainer, name string, v val.Value) {
	switch v.(type) {
	case val.BitStr:
		panic("not yet implemented")
	case val.Bool:
		addBoolConst(c, name, v)
	case val.Float:
		addFloatConst(c, name, v)
	case val.Int:
		addIntConst(c, name, v)
	case val.List:
		switch v.(val.List)[0].(type) {
		case val.BitStr:
			panic("not yet implemented")
		case val.Bool:
			addBoolListConst(c, name, v)
		case val.Int:
			addIntListConst(c, name, v)
		case val.Str:
			panic("not yet implemented")
		default:
			panic("should never happen")
		}
	case val.Str:
		addStrConst(c, name, v)
	default:
		panic("should never happen")
	}
}

func addBoolConst(c *fn.ConstContainer, name string, v val.Value) {
	b := bool(v.(val.Bool))
	if c.BoolConsts == nil {
		c.BoolConsts = map[string]bool{name: b}
	}
	c.BoolConsts[name] = b
}

func addFloatConst(c *fn.ConstContainer, name string, v val.Value) {
	f := float64(v.(val.Float))
	if c.FloatConsts == nil {
		c.FloatConsts = map[string]float64{name: f}
	}
	c.FloatConsts[name] = f
}

func addBoolListConst(c *fn.ConstContainer, name string, v val.Value) {
	l := constifyBoolList(v.(val.List))
	if l == nil {
		return
	}

	if c.BoolListConsts == nil {
		c.BoolListConsts = map[string][]bool{name: l}
	}
	c.BoolListConsts[name] = l
}

func addIntConst(c *fn.ConstContainer, name string, v val.Value) {
	i := int64(v.(val.Int))
	if c.IntConsts == nil {
		c.IntConsts = map[string]int64{name: i}
	}
	c.IntConsts[name] = i
}

func addIntListConst(c *fn.ConstContainer, name string, v val.Value) {
	l := constifyIntList(v.(val.List))
	if l == nil {
		return
	}

	if c.IntListConsts == nil {
		c.IntListConsts = map[string][]int64{name: l}
	}
	c.IntListConsts[name] = l
}

func addStrConst(c *fn.ConstContainer, name string, v val.Value) {
	s := string(v.(val.Str))
	if c.StrConsts == nil {
		c.StrConsts = map[string]string{name: s}
	}
	c.StrConsts[name] = s
}

// constifyBoolList tries to constify list as an bool list.
// If any elemnt is of different type than val.Bool, then it returns nil.
func constifyBoolList(l val.List) []bool {
	bools := []bool{}

	for _, v := range l {
		if i, ok := v.(val.Bool); ok {
			bools = append(bools, bool(i))
		} else {
			return nil
		}
	}

	return bools
}

// constifyIntList tries to constify list as an int list.
// If any elemnt is of different type than val.Int, then it returns nil.
func constifyIntList(l val.List) []int64 {
	ints := []int64{}

	for _, v := range l {
		if i, ok := v.(val.Int); ok {
			ints = append(ints, int64(i))
		} else {
			return nil
		}
	}

	return ints
}
