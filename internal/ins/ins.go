// Package ins implements code responsible for instantiation.
package ins

import (
	"fmt"
	"log"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	fbdlElem "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
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
			main.File().Path, prop.LineNum,
		)
	}

	return nil
}

// Instantiate main bus within given packages scope.
// MainName is the name of the main bus.
func Instantiate(packages prs.Packages, mainName string, zeroTimestamp bool) (*elem.Block, map[string]*elem.Package, error) {
	main, ok := packages["main"][0].Symbols.Get(mainName, prs.ElemInst)
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

	var mainBus *elem.Block

	for pkgName, pkgs := range packages {
		for _, pkg := range pkgs {
			for _, symbol := range pkg.Symbols {
				name := symbol.Name()
				prsElem, ok := symbol.(prs.Element)
				if !ok {
					continue
				}

				if name != mainName && util.IsBaseType(prsElem.Type()) {
					continue
				}

				e := insElement(prsElem)

				if pkgName == "main" && name == mainName {
					mainBus = e.(*elem.Block)
				}
			}
		}
	}

	if mainBus.HasElement("ID") {
		log.Fatalf("'ID' is reserved element name in Main bus")
	}

	id := id()
	hash := int64(mainBus.Hash())
	if busWidth < 32 {
		hash = hash & ((1 << busWidth) - 1)
	}
	// Ignore error, the value has been trimmed to the proper width.
	dflt, _ := val.BitStrFromInt(val.Int(hash), busWidth)
	id.SetDefault(fbdlVal.MakeBitStr(dflt))
	mainBus.AddStatus(id)

	if mainBus.HasElement("TIMESTAMP") {
		log.Fatalf("'TIMESTAMP' is reserved element name in Main bus")
	}
	mainBus.AddStatus(timestamp(zeroTimestamp))

	pkgs := constifyPackages(packages)

	return mainBus, pkgs, nil
}

func insElement(pe prs.Element) fbdlElem.Element {
	typeChain := resolveToBaseType(pe)

	var e fbdlElem.Element
	var err error

	typ := typeChain[0].Type()
	switch typ {
	case "block", "bus":
		e, err = insBlock(typeChain)
	case "config":
		e, err = insConfig(typeChain)
	case "func":
		e, err = insFunc(typeChain)
	case "mask":
		e, err = insMask(typeChain)
	case "param":
		e, err = insParam(typeChain)
	case "return":
		e, err = insReturn(typeChain)
	case "status":
		e, err = insStatus(typeChain)
	case "stream":
		e, err = insStream(typeChain)
	default:
		log.Fatalf(
			"%s: line %d: instantiating element '%s', "+
				"cannot start element instantiation from non base type '%s'",
			pe.File().Path, pe.LineNum(), pe.Name(), typ,
		)
	}

	if err != nil {
		log.Fatalf(
			"%s: line %d: instantiating element '%s': %v",
			pe.File().Path, pe.LineNum(), pe.Name(), err,
		)
	}

	return e
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

		typeChain = append(typeChain, resolveToBaseType(type_elem)...)
	}

	typeChain = append(typeChain, e)
	return typeChain
}
