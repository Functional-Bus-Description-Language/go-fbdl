package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

func constifyPackages(packages prs.Packages) map[string]*fn.Package {
	cPkgs := map[string]*fn.Package{}

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

func constifyPkg(pkg *prs.Package) *fn.Package {
	p := fn.Package{}

	for _, s := range pkg.Symbols {
		if c, ok := s.(*prs.Const); ok {
			v, err := c.Value.Eval()
			if err != nil {
				panic("not yet implemented")
			}
			constContainer.AddConst(&p.ConstContainer, c.Name(), v)
		}
	}

	return &p
}
