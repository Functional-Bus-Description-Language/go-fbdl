package inst

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parse"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

type Element struct {
	name       string
	baseType   string
	Properties map[string]val.Value
	Constants  map[string]val.Value
	Elements   map[string]*Element
}

func (elem *Element) applyType(type_ parse.Element, resolvedArgs map[string]val.Value) error {
	if elem.baseType == "" {
		if !util.IsBaseType(type_.Type()) {
			return fmt.Errorf("cannot start element instantiation from non base type '%s'", type_.Type())
		}

		elem.baseType = type_.Type()
	}

	if def, ok := type_.(*parse.ElementDefinition); ok {
		elem.name = def.Name()
	}

	if resolvedArgs != nil {
		type_.SetResolvedArgs(resolvedArgs)
	}

	for name, prop := range type_.Properties() {
		if err := util.IsValidProperty(name, elem.baseType); err != nil {
			panic("implement me")
		}
		err := checkProperty(name, prop)
		if err != nil {
			return fmt.Errorf("\n  %s: line %d: %v", type_.FilePath(), prop.LineNumber, err)
		}
		if _, exist := elem.Properties[name]; exist {
			return fmt.Errorf(
				"cannot set property '%s', property is already set in one of ancestor types",
				name,
			)
		}
		v, err := prop.Value.Eval()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}
		elem.Properties[name] = v
	}

	for _, s := range type_.Symbols() {
		pe, ok := s.(*parse.ElementDefinition)
		if !ok {
			continue
		}

		e := instantiateElement(pe)

		if util.IsValidType(elem.baseType, e.baseType) == false {
			return fmt.Errorf(
				"element '%s' of base type '%s' cannot be instantiated in element of base type '%s'",
				e.name, e.baseType, elem.baseType,
			)
		}

		if _, ok := elem.Elements[e.name]; ok {
			return fmt.Errorf(
				"cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types",
				e.name,
			)
		}

		elem.Elements[e.name] = e
	}

	return nil
}
