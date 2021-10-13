package inst

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parse"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/value"
	"log"
)

const defaultBusWidth uint = 32

var busWidth uint

func setBusWidth(main parse.Symbol) error {
	e, ok := main.(*parse.ElementDefinition)
	if !ok {
		panic("FIX ME")
	}

	prop, ok := e.Properties()["width"]
	if !ok {
		busWidth = defaultBusWidth
		return nil
	}

	v, err := prop.Value.Value()
	if err != nil {
		return fmt.Errorf("cannot evaluate main bus width")
	}

	if v, ok := v.(value.Integer); ok {
		busWidth = uint(v.V)
	} else {
		log.Fatalf(
			"%s: line %d: 'main' bus 'width' property must be of type 'integer'",
			main.FilePath(), prop.LineNumber,
		)
	}

	return nil
}

func Instantiate(packages parse.Packages) *Element {
	main, ok := packages["main"][0].Symbols["main"]
	if !ok {
		log.Println("instantiation: there is no main bus: returning empty dictionary")
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
				e, ok := symbol.(parse.Element)
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

func instantiateElement(e parse.Element) *Element {
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

func resolveToBaseType(e parse.Element) []parse.Element {
	typeChain := []parse.Element{}

	if !util.IsBaseType(e.Type()) {
		s, err := e.GetSymbol(e.Type())
		if err != nil {
			log.Fatalf("cannot get symbol '%s': %v", e.Type(), err)
		}
		type_elem := s.(parse.Element)

		for _, bt := range resolveToBaseType(type_elem) {
			typeChain = append(typeChain, bt)
		}
	}

	typeChain = append(typeChain, e)
	return typeChain
}

func instantiateTypeChain(tc []parse.Element) (*Element, error) {
	inst := &Element{
		Properties: map[string]value.Value{},
		Constants:  map[string]value.Value{},
		Elements:   map[string]*Element{},
	}

	for i, t := range tc {
		resolvedArgs := make(map[string]value.Value)
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
