// Package ins implements code responsible for instantiation.
package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
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
		return fmt.Errorf("cannot evaluate 'Main' bus 'width' property")
	}

	if v, ok := v.(val.Int); ok {
		busWidth = int64(v)
	} else {
		log.Fatalf(
			"%s: line %d: 'Main' bus 'width' property must be of type 'integer'",
			main.File().Path, prop.LineNum,
		)
	}

	return nil
}

func Instantiate(packages prs.Packages, zeroTimestamp bool) *Element {
	main, ok := packages["main"][0].Symbols.Get("Main", prs.ElemInst)
	if !ok {
		log.Println("instantiation: there is no 'Main' bus; returning nil")
		return nil
	}

	err := setBusWidth(main)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	err = resolveArgLists(packages)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	var mainBus *Element

	for pkgName, pkgs := range packages {
		for _, pkg := range pkgs {
			for _, symbol := range pkg.Symbols {
				name := symbol.Name()
				e, ok := symbol.(prs.Element)
				if !ok {
					continue
				}

				if name != "Main" && util.IsBaseType(e.Type()) {
					continue
				}

				elem := instantiateElement(e)

				if pkgName == "main" && name == "Main" {
					mainBus = elem
				}
			}
		}
	}

	if _, exists := mainBus.Elems.Get("X_ID_X"); exists {
		panic("X_ID_X is reserved element name")
	}

	id := x_id_x()
	hash := int64(mainBus.hash())
	if busWidth < 32 {
		hash = hash & ((1 << busWidth) - 1)
	}
	// Ignore error, the value has been trimmed to the proper width.
	dflt, _ := val.BitStrFromInt(val.Int(hash), busWidth)
	id.Props["default"] = dflt
	mainBus.Elems.Add(id)
	if _, ok := mainBus.Consts["X_ID_X"]; ok {
		log.Fatalf("Main bus cannot have constant named 'X_ID_X'")
	}
	// ID will never be wider than 64 bits, so use val.Int.
	mainBus.Consts["X_ID_X"] = val.Int(hash)

	if _, exists := mainBus.Elems.Get("X_TIMESTAMP_X"); exists {
		panic("X_TIMESTAMP_X is reserved element name")
	}
	mainBus.Elems.Add(x_timestamp_x(zeroTimestamp))

	return mainBus
}

func instantiateElement(e prs.Element) *Element {
	typeChain := resolveToBaseType(e)
	elem, err := instantiateTypeChain(typeChain)
	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.File().Path, e.LineNum(), e.Name(), err,
		)
	}

	if elem.Count < 0 {
		log.Fatalf(
			"%s: line %d: negative size (%d) of '%s' array",
			e.File().Path, e.LineNum(), elem.Count, e.Name(),
		)
	}

	fillProps(elem)

	if err = elem.makeGrps(); err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.File().Path, e.LineNum(), e.Name(), err,
		)
	}

	err = elem.processDflt()
	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			e.File().Path, e.LineNum(), e.Name(), err,
		)
	}

	return elem
}

func resolveToBaseType(e prs.Element) []prs.Element {
	typeChain := []prs.Element{}

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
		Props:  map[string]val.Value{},
		Consts: map[string]val.Value{},
		Elems:  ElemContainer{},
	}

	for i, t := range tc {
		resolvedArgs := make(map[string]prs.Expr)
		if (i+1) < len(tc) && tc[i+1].ResolvedArgs() != nil {
			resolvedArgs = tc[i+1].ResolvedArgs()
		}
		err := inst.applyType(t, resolvedArgs)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
	}

	if tc[len(tc)-1].Doc() != "" {
		inst.Doc = tc[len(tc)-1].Doc()
	}

	return inst, nil
}
