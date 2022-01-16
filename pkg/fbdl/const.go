package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type ConstContainer struct {
	BoolConsts     map[string]bool
	BoolListConsts map[string][]bool
	IntConsts      map[string]int64
	IntListConsts  map[string][]int64
	StrConsts      map[string]string
}

func (cc ConstContainer) HasConsts() bool {
	if len(cc.BoolConsts) != 0 || len(cc.BoolListConsts) != 0 {
		return true
	}
	if len(cc.IntConsts) != 0 || len(cc.IntListConsts) != 0 {
		return true
	}
	if len(cc.StrConsts) != 0 {
		return true
	}

	return false
}

func (cc *ConstContainer) addConst(name string, v val.Value) {
	switch v.(type) {
	case val.BitStr:
		panic("not yet implemented")
	case val.Bool:
		cc.addBoolConst(name, v)
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

func (cc *ConstContainer) addBoolConst(name string, v val.Value) {
	b := bool(v.(val.Bool))
	if cc.BoolConsts == nil {
		cc.BoolConsts = map[string]bool{name: b}
	}
	cc.BoolConsts[name] = b
}

func (cc *ConstContainer) addBoolListConst(name string, v val.Value) {
	l := constifyBoolList(v.(val.List))
	if l == nil {
		return
	}

	if cc.BoolListConsts == nil {
		cc.BoolListConsts = map[string][]bool{name: l}
	}
	cc.BoolListConsts[name] = l
}

func (cc *ConstContainer) addIntConst(name string, v val.Value) {
	i := int64(v.(val.Int))
	if cc.IntConsts == nil {
		cc.IntConsts = map[string]int64{name: i}
	}
	cc.IntConsts[name] = i
}

func (cc *ConstContainer) addIntListConst(name string, v val.Value) {
	l := constifyIntList(v.(val.List))
	if l == nil {
		return
	}

	if cc.IntListConsts == nil {
		cc.IntListConsts = map[string][]int64{name: l}
	}
	cc.IntListConsts[name] = l
}

func (cc *ConstContainer) addStrConst(name string, v val.Value) {
	s := string(v.(val.Str))
	if cc.StrConsts == nil {
		cc.StrConsts = map[string]string{name: s}
	}
	cc.StrConsts[name] = s
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
