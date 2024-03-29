package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/pkg"
)

// Compile compiles functional bus description for a main bus named mainName located in the file which path is provided as mainPath.
// If noTimestamp is true, then the bus timestamp is not generated.
func Compile(mainPath, mainName string, addTimestamp bool) (*fn.Block, map[string]*pkg.Package, error) {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	bus, insPkgs, err := ins.Instantiate(packages, mainName)
	if err != nil {
		return nil, nil, err
	}

	// Below loop is needed, as map of concrete type cannot be by default treated
	// as map of interfaces, even if the concrete type meets the interface requirements.
	pkgs := map[string]*pkg.Package{}
	for k, v := range insPkgs {
		pkgs[k] = v
	}

	reg.Registerify(bus, addTimestamp)

	return bus, pkgs, nil
}
