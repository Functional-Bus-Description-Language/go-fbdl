package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

type Element struct {
	Name       string
	LineNumber uint32
	BaseType   string
	IsArray    bool
	Count      int64
	Properties map[string]val.Value
	Constants  map[string]val.Value
	Elements   ElementContainer
}

func (elem *Element) applyType(type_ prs.Element, resolvedArgs map[string]prs.Expression) error {
	if elem.BaseType == "" {
		if !util.IsBaseType(type_.Type()) {
			return fmt.Errorf("cannot start element instantiation from non base type '%s'", type_.Type())
		}

		elem.BaseType = type_.Type()
	}

	if def, ok := type_.(*prs.ElementDefinition); ok {
		elem.Name = def.Name()
		elem.LineNumber = def.LineNumber()
	}

	if resolvedArgs != nil {
		type_.SetResolvedArgs(resolvedArgs)
	}

	for name, prop := range type_.Properties() {
		if err := util.IsValidProperty(name, elem.BaseType); err != nil {
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
		err = checkPropertyConflict(elem, name)
		if err != nil {
			return fmt.Errorf("line %d: %v", prop.LineNumber, err)
		}
		elem.Properties[name] = v
	}

	for _, s := range type_.Symbols() {
		pe, ok := s.(*prs.ElementDefinition)
		if !ok {
			continue
		}

		e := instantiateElement(pe)

		if util.IsValidType(elem.BaseType, e.BaseType) == false {
			return fmt.Errorf(
				"element '%s' of base type '%s' cannot be instantiated in element of base type '%s'",
				e.Name, e.BaseType, elem.BaseType,
			)
		}

		if !elem.Elements.Add(e) {
			return fmt.Errorf(
				"cannot instantiate element '%s', element with such name is already instantiated in one of ancestor types",
				e.Name,
			)
		}
	}

	if ed, ok := type_.(*prs.ElementDefinition); ok {
		if elem.IsArray {
			panic("should never happen")
		}
		if ed.IsArray {
			elem.IsArray = true
			count, err := ed.Count.Eval()

			if count.Type() != "integer" {
				return fmt.Errorf("size of array must be of 'integer' type, current type '%s'", count.Type())
			}

			if err != nil {
				return fmt.Errorf("applying type '%s': %v", type_.Name(), err)
			}
			elem.Count = int64(count.(val.Int))
		} else {
			elem.Count = int64(1)
		}
	}

	return nil
}
