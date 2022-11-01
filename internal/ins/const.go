package ins

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util/constContainer"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func constifyPackages(packages prs.Packages) map[string]*elem.Package {
	cPkgs := map[string]*elem.Package{}

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

func constifyPkg(pkg *prs.Package) *elem.Package {
	p := elem.Package{}

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
