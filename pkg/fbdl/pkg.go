package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Package struct {
	IntConsts map[string]int64
	StrConsts map[string]string
}

func (p Package) HasConsts() bool {
	if len(p.IntConsts) != 0 {
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
			cPkgs[name] = Package{
				IntConsts: map[string]int64{},
				StrConsts: map[string]string{},
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
						cPkgs[name].IntConsts[c.Name()] = int64(v.(val.Int))
					case val.List:
						panic("not yet implemented")
					case val.Str:
						cPkgs[name].StrConsts[c.Name()] = string(v.(val.Str))
					default:
						panic("should never happen")
					}
				}
			}
		}
	}

	return cPkgs
}
