package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
	"log"
)

const defaultBusWidth uint = 32

var busWidth uint

func setBusWidth(main prs.Symbol) error {
	e, ok := main.(*prs.ElementDefinition)
	if !ok {
		panic("FIX ME")
	}

	prop, ok := e.Properties()["width"]
	if !ok {
		busWidth = defaultBusWidth
		return nil
	}

	v, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate main bus width")
	}

	if v, ok := v.(val.Int); ok {
		busWidth = uint(v.V)
	} else {
		log.Fatalf(
			"%s: line %d: 'main' bus 'width' property must be of type 'integer'",
			main.FilePath(), prop.LineNumber,
		)
	}

	return nil
}

func Instantiate(packages prs.Packages) *Element {
	main, ok := packages["main"][0].Symbols["main"]
	if !ok {
		log.Println("instantiation: there is no main bus: returning nil")
		return nil
	}

	setBusWidth(main)

	err := resolveArgumentLists(packages)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	var main_bus *Element

	for pkg_name, pkgs := range packages {
		for _, pkg := range pkgs {
			for name, symbol := range pkg.Symbols {
				e, ok := symbol.(prs.Element)
				if !ok {
					continue
				}

				if name != "main" && util.IsBaseType(e.Type()) {
					continue
				}

				elem := instantiateElement(e)

				if pkg_name == "main" && name == "main" {
					main_bus = elem
				}
			}
		}
	}

	return main_bus
}

func instantiateElement(e prs.Element) *Element {
	typeChain := resolveToBaseType(e)
	instance, err := instantiateTypeChain(typeChain)
	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.FilePath(), e.LineNumber(), e.Name(), err,
		)
	}

	return instance
}

func resolveToBaseType(e prs.Element) []prs.Element {
	typeChain := []prs.Element{}

	if !util.IsBaseType(e.Type()) {
		s, err := e.GetSymbol(e.Type())
		if err != nil {
			log.Fatalf("cannot get symbol '%s': %v", e.Type(), err)
		}
		type_elem := s.(prs.Element)

		for _, bt := range resolveToBaseType(type_elem) {
			typeChain = append(typeChain, bt)
		}
	}

	typeChain = append(typeChain, e)
	return typeChain
}

func instantiateTypeChain(tc []prs.Element) (*Element, error) {
	inst := &Element{
		Properties: map[string]val.Value{},
		Constants:  map[string]val.Value{},
		Elements:   map[string]*Element{},
	}

	for i, t := range tc {
		resolvedArgs := make(map[string]val.Value)
		if (i+1) < len(tc) && tc[i+1].ResolvedArgs() != nil {
			resolvedArgs = tc[i+1].ResolvedArgs()
		}
		err := inst.applyType(t, resolvedArgs)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	return inst, nil
}
