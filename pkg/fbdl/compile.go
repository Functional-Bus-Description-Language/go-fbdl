package fbdl

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// Compile compiles functional bus description for a main bus named mainName located in the file which path is provided as mainPath.
func Compile(mainPath, mainName string) (elem.Block, map[string]elem.Package, error) {
	packages := prs.DiscoverPackages(mainPath)
	prs.ParsePackages(packages)

	bus, insPkgs, err := ins.Instantiate(packages, mainName, false)
	if err != nil {
		return nil, nil, fmt.Errorf("instantiation: %v", err)
	}

	// Below loop is needed, as map of concrete type cannot be by default treated
	// as map of interfaces, even if the concrete type meets the interface requirements.
	pkgs := map[string]elem.Package{}
	for k, v := range insPkgs {
		pkgs[k] = v
	}

	reg.Registerify(bus)

	return bus, pkgs, nil
}
