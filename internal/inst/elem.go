package inst

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parse"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/value"
)

type Element struct {
	name       string
	baseType   string
	Properties map[string]value.Value
	Constants  map[string]value.Value
	Elements   map[string]*Element
}

func (e *Element) applyType(t parse.Element, resolvedArgs map[string]value.Value) error {
	if e.baseType == "" {
		if !util.IsBaseType(t.Type()) {
			return fmt.Errorf("cannot start element instantiation from non base type '%s'", t.Type())
		}

		e.baseType = t.Type()
	}

	if def, ok := t.(*parse.ElementDefinition); ok {
		e.name = def.Name()
	}

	if resolvedArgs != nil {
		t.SetResolvedArgs(resolvedArgs)
	}

	for name, prop := range t.Properties() {
		if util.IsValidProperty(e.baseType, name) == false {
			panic("implement me")
		}
		err := checkProperty(name, prop)
		if err != nil {
			return fmt.Errorf("some message: %v", err)
		}
		if _, exist := e.Properties[name]; exist {
			return fmt.Errorf(
				"cannot set property '%s', property is already set in one of ancestor types",
				name,
			)
		}
		v, err := prop.Value.Value()
		if err != nil {
			return fmt.Errorf("cannot evaluate expression")
		}
		e.Properties[name] = v
	}

	return nil
}
