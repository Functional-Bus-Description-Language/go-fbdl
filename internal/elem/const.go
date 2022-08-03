package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type cc struct {
	BoolConsts     map[string]bool
	BoolListConsts map[string][]bool
	FloatConsts    map[string]float64
	IntConsts      map[string]int64
	IntListConsts  map[string][]int64
	StrConsts      map[string]string
}

type ConstContainer struct {
	cc
}

func (c ConstContainer) BoolConsts() map[string]bool       { return c.cc.BoolConsts }
func (c ConstContainer) BoolListConsts() map[string][]bool { return c.cc.BoolListConsts }
func (c ConstContainer) IntConsts() map[string]int64       { return c.cc.IntConsts }
func (c ConstContainer) IntListConsts() map[string][]int64 { return c.cc.IntListConsts }
func (c ConstContainer) StrConsts() map[string]string      { return c.cc.StrConsts }

// HasConst returns true if container already has constant with given name.
func (c ConstContainer) HasConst(name string) bool {
	if _, ok := c.cc.BoolConsts[name]; ok {
		return true
	}
	if _, ok := c.cc.BoolListConsts[name]; ok {
		return true
	}
	if _, ok := c.cc.IntConsts[name]; ok {
		return true
	}
	if _, ok := c.cc.IntListConsts[name]; ok {
		return true
	}
	if _, ok := c.cc.StrConsts[name]; ok {
		return true
	}

	return false
}

func (c ConstContainer) HasConsts() bool {
	if len(c.cc.BoolConsts) != 0 || len(c.cc.BoolListConsts) != 0 {
		return true
	}
	if len(c.cc.IntConsts) != 0 || len(c.cc.IntListConsts) != 0 {
		return true
	}
	if len(c.cc.StrConsts) != 0 {
		return true
	}

	return false
}

func (cc *ConstContainer) AddConst(name string, v val.Value) {
	switch v.(type) {
	case val.BitStr:
		panic("not yet implemented")
	case val.Bool:
		cc.addBoolConst(name, v)
	case val.Float:
		cc.addFloatConst(name, v)
	case val.Int:
		cc.addIntConst(name, v)
	case val.List:
		switch v.(val.List)[0].(type) {
		case val.BitStr:
			panic("not yet implemented")
		case val.Bool:
			cc.addBoolListConst(name, v)
		case val.Int:
			cc.addIntListConst(name, v)
		case val.Str:
			panic("not yet implemented")
		default:
			panic("should never happen")
		}
	case val.Str:
		cc.addStrConst(name, v)
	default:
		panic("should never happen")
	}
}

func (c *ConstContainer) addBoolConst(name string, v val.Value) {
	b := bool(v.(val.Bool))
	if c.cc.BoolConsts == nil {
		c.cc.BoolConsts = map[string]bool{name: b}
	}
	c.cc.BoolConsts[name] = b
}

func (c *ConstContainer) addFloatConst(name string, v val.Value) {
	f := float64(v.(val.Float))
	if c.cc.FloatConsts == nil {
		c.cc.FloatConsts = map[string]float64{name: f}
	}
	c.cc.FloatConsts[name] = f
}

func (c *ConstContainer) addBoolListConst(name string, v val.Value) {
	l := constifyBoolList(v.(val.List))
	if l == nil {
		return
	}

	if c.cc.BoolListConsts == nil {
		c.cc.BoolListConsts = map[string][]bool{name: l}
	}
	c.cc.BoolListConsts[name] = l
}

func (c *ConstContainer) addIntConst(name string, v val.Value) {
	i := int64(v.(val.Int))
	if c.cc.IntConsts == nil {
		c.cc.IntConsts = map[string]int64{name: i}
	}
	c.cc.IntConsts[name] = i
}

func (c *ConstContainer) addIntListConst(name string, v val.Value) {
	l := constifyIntList(v.(val.List))
	if l == nil {
		return
	}

	if c.cc.IntListConsts == nil {
		c.cc.IntListConsts = map[string][]int64{name: l}
	}
	c.cc.IntListConsts[name] = l
}

func (c *ConstContainer) addStrConst(name string, v val.Value) {
	s := string(v.(val.Str))
	if c.cc.StrConsts == nil {
		c.cc.StrConsts = map[string]string{name: s}
	}
	c.cc.StrConsts[name] = s
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
