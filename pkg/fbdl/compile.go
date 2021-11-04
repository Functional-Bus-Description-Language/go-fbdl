package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"
)

// Compile compiles functional bus description for a main bus located in the file which path is provided as mainPath.
func Compile(mainPath string) *reg.BlockElement {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	insBus := ins.Instantiate(packages)

	regBus := reg.Registerify(insBus)

	return regBus
}
