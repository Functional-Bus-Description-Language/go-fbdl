// Package ins implements code responsible for instantiation.
package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/pkg"
)

const dfltBusWidth int64 = 32

var busWidth int64

func setBusWidth(main *prs.Inst) error {
	prop, ok := main.Props().Get("width")
	if !ok {
		busWidth = dfltBusWidth
		return nil
	}

	v, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf(
			"%s:%s: cannot evaluate main bus 'width' property",
			main.File().Path, prop.Loc(),
		)
	}

	if vi, ok := v.(val.Int); ok {
		busWidth = int64(vi)
	} else {
		return fmt.Errorf(
			"%s:%s: main bus 'width' property must be of integer type, current type %s",
			main.File().Path, prop.Loc(), v.Type(),
		)
	}

	return nil
}

// Instantiate main bus within given packages scope.
// MainName is the name of the main bus.
func Instantiate(packages prs.Packages, mainName string) (*fn.Block, map[string]*pkg.Package, error) {
	main, err := packages["main"][0].GetInst(mainName)
	if err != nil {
		return nil, nil, fmt.Errorf("%v", err)
	}
	log.Printf("debug: instantiating '%s' as the main bus", mainName)

	err = setBusWidth(main)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = resolveArgLists(packages)
	if err != nil {
		log.Fatalf("%v", err)
	}

	var mainBus *fn.Block

	for pkgName, pkgs := range packages {
		for _, pkg := range pkgs {
			for _, symbol := range pkg.Symbols() {
				name := symbol.Name()
				prsFn, ok := symbol.(prs.Functionality)
				if !ok {
					continue
				}

				if name != mainName && util.IsBaseType(prsFn.Type()) {
					continue
				}

				f := insFunctionality(prsFn)

				if pkgName == "main" && name == mainName {
					mainBus = f.(*fn.Block)
				}
			}
		}
	}

	pkgs := constifyPackages(packages)

	return mainBus, pkgs, nil
}

func insFunctionality(pf prs.Functionality) fn.Functionality {
	typeChain := resolveToBaseType(pf)

	var f fn.Functionality
	var err error

	typ := typeChain[0].Type()
	switch typ {
	case "blackbox":
		f, err = insBlackbox(typeChain)
	case "block", "bus":
		f, err = insBlock(typeChain)
	case "config":
		f, err = insConfig(typeChain)
	case "irq":
		f, err = insIrq(typeChain)
	case "mask":
		f, err = insMask(typeChain)
	case "param":
		f, err = insParam(typeChain)
	case "proc":
		f, err = insProc(typeChain)
	case "return":
		f, err = insReturn(typeChain)
	case "static":
		f, err = insStatic(typeChain)
	case "status":
		f, err = insStatus(typeChain)
	case "stream":
		f, err = insStream(typeChain)
	default:
		panic("should never happen")
	}

	if err != nil {
		log.Fatalf("%v", err)
	}

	return f
}

func resolveToBaseType(f prs.Functionality) []prs.Functionality {
	typeChain := []prs.Functionality{}

	if !util.IsBaseType(f.Type()) {
		var s prs.Symbol
		var err error
		if f.Scope() != nil {
			s, err = f.Scope().GetType(f.Type())
		} else {
			s, err = f.File().GetType(f.Type())
		}
		if err != nil {
			log.Fatalf(
				"%s:%d:%d: %v",
				f.File().Path, f.Line(), f.Col(), err,
			)
		}
		typeFn := s.(prs.Functionality)

		typeChain = append(typeChain, resolveToBaseType(typeFn)...)
	}

	typeChain = append(typeChain, f)
	return typeChain
}
