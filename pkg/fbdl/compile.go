package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// Compile compiles functional bus description for a Main bus located in the file which path is provided as mainPath.
func Compile(mainPath string) (*elem.Block, map[string]elem.Package) {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	insBus := ins.Instantiate(packages, false)

	regBus := elem.Registerify(insBus)

	pkgsConsts := elem.ConstifyPackages(packages)

	return regBus, pkgsConsts
}
