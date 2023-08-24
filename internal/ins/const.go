package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/pkg"
)

func constifyPackages(packages prs.Packages) map[string]*pkg.Package {
	cPkgs := map[string]*pkg.Package{}

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

func constifyPkg(pp *prs.Package) *pkg.Package {
	p := pkg.Package{}

	for _, s := range pp.Symbols {
		if c, ok := s.(*prs.Const); ok {
			v, err := c.Value.Eval()
			if err != nil {
				panic("not yet implemented")
			}
			constContainer.AddConst(&p.Consts, c.Name(), v)
		}
	}

	return &p
}
