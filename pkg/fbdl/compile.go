package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// Compile compiles functional bus description for a main bus named mainName located in the file which path is provided as mainPath.
func Compile(mainPath, mainName string) (elem.Block, map[string]elem.Package) {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	bus, insPkgs := ins.Instantiate(packages, mainName, false)

	// Below loop is needed, as map of concrete type cannot be by default treated
	// as map of interfaces, even if the concrete type meets the interface requirements.
	pkgs := map[string]elem.Package{}
	for k, v := range insPkgs {
		pkgs[k] = v
	}

	if bus == nil {
		return bus, pkgs
	}

	reg.Registerify(bus)

	return bus, pkgs
}
