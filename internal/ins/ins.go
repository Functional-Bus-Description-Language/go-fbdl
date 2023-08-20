// Package ins implements code responsible for instantiation.
package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

const dfltBusWidth int64 = 32

var busWidth int64

func setBusWidth(main prs.Symbol) error {
	e, ok := main.(*prs.Inst)
	if !ok {
		panic("FIX ME")
	}

	prop, ok := e.Props().Get("width")
	if !ok {
		busWidth = dfltBusWidth
		return nil
	}

	v, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate main bus 'width' property")
	}

	if v, ok := v.(val.Int); ok {
		busWidth = int64(v)
	} else {
		log.Fatalf(
			"%s: line %d: main bus 'width' property must be of type 'integer'",
			main.File().Path, prop.Line,
		)
	}

	return nil
}

// Instantiate main bus within given packages scope.
// MainName is the name of the main bus.
func Instantiate(packages prs.Packages, mainName string) (*fn.Block, map[string]*fn.Package, error) {
	main, ok := packages["main"][0].Symbols.Get(mainName, prs.FuncInst)
	if !ok {
		return nil, nil, fmt.Errorf("'%s' bus not found", mainName)
	}
	log.Printf("debug: instantiating '%s' as the main bus", mainName)

	err := setBusWidth(main)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	err = resolveArgLists(packages)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	var mainBus *fn.Block

	for pkgName, pkgs := range packages {
		for _, pkg := range pkgs {
			for _, symbol := range pkg.Symbols {
				name := symbol.Name()
				prsElem, ok := symbol.(prs.Functionality)
				if !ok {
					continue
				}

				if name != mainName && util.IsBaseType(prsElem.Type()) {
					continue
				}

				e := insElement(prsElem)

				if pkgName == "main" && name == mainName {
					mainBus = e.(*fn.Block)
				}
			}
		}
	}

	pkgs := constifyPackages(packages)

	return mainBus, pkgs, nil
}

func insElement(pf prs.Functionality) fn.Functionality {
	typeChain := resolveToBaseType(pf)

	var f fn.Functionality
	var err error

	typ := typeChain[0].Type()
	switch typ {
	case "block", "bus":
		f, err = insBlock(typeChain)
	case "config":
		f, err = insConfig(typeChain)
	case "irq":
		f, err = insIrq(typeChain)
	case "mask":
		f, err = insMask(typeChain)
	case "memory":
		f, err = insMemory(typeChain)
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
		log.Fatalf(
			"%s: line %d: instantiating element '%s', "+
				"cannot start element instantiation from non base type '%s'",
			pf.File().Path, pf.Line(), pf.Name(), typ,
		)
	}

	if err != nil {
		log.Fatalf("%s:%v", pf.File().Path, err)
	}

	return f
}

func resolveToBaseType(e prs.Functionality) []prs.Functionality {
	typeChain := []prs.Functionality{}

	if !util.IsBaseType(e.Type()) {
		var s prs.Symbol
		var err error
		if e.Parent() != nil {
			s, err = e.Parent().GetSymbol(e.Type(), prs.TypeDef)
		} else {
			s, err = e.File().GetSymbol(e.Type(), prs.TypeDef)
		}
		if err != nil {
			log.Fatalf("cannot get symbol '%s': %v", e.Type(), err)
		}
		type_elem := s.(prs.Functionality)

		typeChain = append(typeChain, resolveToBaseType(type_elem)...)
	}

	typeChain = append(typeChain, e)
	return typeChain
}
