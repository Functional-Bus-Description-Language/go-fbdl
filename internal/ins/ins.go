// Package ins implements code responsible for instantiation.
package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"log"
)

const defaultBusWidth int64 = 32

var busWidth int64

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
		busWidth = int64(v)
	} else {
		log.Fatalf(
			"%s: line %d: 'main' bus 'width' property must be of type 'integer'",
			main.FilePath(), prop.LineNumber,
		)
	}

	return nil
}

func Instantiate(packages prs.Packages) *Element {
	main, ok := packages["main"][0].Symbols.Get("main")
	if !ok {
		log.Println("instantiation: there is no main bus; returning nil")
		return nil
	}

	err := setBusWidth(main)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	err = resolveArgumentLists(packages)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	var main_bus *Element

	for pkg_name, pkgs := range packages {
		for _, pkg := range pkgs {
			for _, symbol := range pkg.Symbols {
				name := symbol.Name()
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

	if _, exists := main_bus.Elements.Get("x_uuid_x"); exists {
		panic("x_uuid_x is reserved element name")
	}
	main_bus.Elements.Add(x_uuid_x())

	if _, exists := main_bus.Elements.Get("x_timestamp_x"); exists {
		panic("x_timestamp_x is reserved element name")
	}
	main_bus.Elements.Add(x_timestamp_x())

	return main_bus
}

func instantiateElement(e prs.Element) *Element {
	typeChain := resolveToBaseType(e)
	elem, err := instantiateTypeChain(typeChain)
	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.FilePath(), e.LineNumber(), e.Name(), err,
		)
	}

	if elem.Count < 0 {
		log.Fatalf(
			"%s: line %d: negative size (%d) of '%s' array",
			e.FilePath(), e.LineNumber(), elem.Count, e.Name(),
		)
	}

	fillProperties(elem)

	if err = elem.makeGroups(); err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.FilePath(), e.LineNumber(), e.Name(), err,
		)
	}

	err = elem.processDefault()
	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.FilePath(), e.LineNumber(), e.Name(), err,
		)
	}

	return elem
}

func resolveToBaseType(e prs.Element) []prs.Element {
	typeChain := []prs.Element{}

	if !util.IsBaseType(e.Type()) {
		s, err := e.Parent().GetSymbol(e.Type())
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
		Elements:   ElementContainer{},
	}

	for i, t := range tc {
		resolvedArgs := make(map[string]prs.Expression)
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
