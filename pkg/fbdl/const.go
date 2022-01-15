package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type constContainer struct {
	IntConsts     map[string]int64
	IntListConsts map[string][]int64
	StrConsts     map[string]string
}

func (cc *constContainer) addConst(name string, v val.Value) {
	switch v.(type) {
	case val.BitStr:
		panic("not yet implemented")
	case val.Bool:
		panic("not yet implemented")
	case val.Int:
		cc.addIntConst(name, v)
	case val.List:
		switch v.(val.List)[0].(type) {
		case val.BitStr:
			panic("not yet implemented")
		case val.Bool:
			panic("not yet implemented")
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

func (cc *constContainer) addIntConst(name string, v val.Value) {
	i := int64(v.(val.Int))
	if cc.IntConsts == nil {
		cc.IntConsts = map[string]int64{name: i}
	}
	cc.IntConsts[name] = i
}

func (cc *constContainer) addIntListConst(name string, v val.Value) {
	l := constifyIntList(v.(val.List))
	if l == nil {
		return
	}

	if cc.IntListConsts == nil {
		cc.IntListConsts = map[string][]int64{name: l}
	}
	cc.IntListConsts[name] = l
}

func (cc *constContainer) addStrConst(name string, v val.Value) {
	s := string(v.(val.Str))
	if cc.StrConsts == nil {
		cc.StrConsts = map[string]string{name: s}
	}
	cc.StrConsts[name] = s
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
