package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Package struct {
	IntConsts     map[string]int64
	IntListConsts map[string][]int64
	StrConsts     map[string]string
}

func (p Package) HasConsts() bool {
	if len(p.IntConsts) != 0 || len(p.IntListConsts) != 0 {
		return true
	}
	if len(p.StrConsts) != 0 {
		return true
	}

	return false
}

func ConstifyPackages(packages prs.Packages) map[string]Package {
	cPkgs := map[string]Package{}

	// TODO: Resolve name conflicts.
	for name, pkgs := range packages {
		if len(pkgs) > 1 {
			panic("not yet implemented")
		}
		for _, pkg := range pkgs {
			cPkgs[name] = constifyPkg(pkg)
		}
	}

	return cPkgs
}

func constifyPkg(pkg *prs.Package) Package {
	p := Package{
		IntConsts:     map[string]int64{},
		IntListConsts: map[string][]int64{},
		StrConsts:     map[string]string{},
	}

	for _, s := range pkg.Symbols {
		if c, ok := s.(*prs.Const); ok {
			v, err := c.Value.Eval()
			if err != nil {
				panic("not yet implemented")
			}
			switch v.(type) {
			case val.BitStr:
				panic("not yet implemented")
			case val.Bool:
				panic("not yet implemented")
			case val.Int:
				p.IntConsts[c.Name()] = int64(v.(val.Int))
			case val.List:
				switch v.(val.List)[0].(type) {
				case val.BitStr:
					panic("not yet implemented")
				case val.Bool:
					panic("not yet implemented")
				case val.Int:
					p.IntListConsts[c.Name()] = constifyIntList(v.(val.List))
				case val.Str:
					panic("not yet implemented")
				default:
					panic("should never happen")
				}
			case val.Str:
				p.StrConsts[c.Name()] = string(v.(val.Str))
			default:
				panic("should never happen")
			}
		}
	}

	return p
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
