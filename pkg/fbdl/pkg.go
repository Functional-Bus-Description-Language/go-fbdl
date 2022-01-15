package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
)

type Package struct {
	constContainer
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
	p := Package{}

	for _, s := range pkg.Symbols {
		if c, ok := s.(*prs.Const); ok {
			v, err := c.Value.Eval()
			if err != nil {
				panic("not yet implemented")
			}
			p.addConst(c.Name(), v)
		}
	}

	return p
}
