package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
)

// Compile compiles functional bus description for a main bus located in the file which path is provided as mainPath.
func Compile(mainPath string) (*Block, map[string]Package) {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	insBus := ins.Instantiate(packages)

	regBus := Registerify(insBus)

	pkgsConsts := ConstifyPackages(packages)

	return regBus, pkgsConsts
}
