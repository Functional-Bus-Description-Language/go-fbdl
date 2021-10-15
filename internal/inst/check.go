package inst

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/parse"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

// Check property value type, value and conflicting properties within element.
func checkProperty(name string, prop parse.Property) error {
	pv, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate expression: %v", err)
	}

	switch name {
	case "atomic", "once":
		if _, ok := pv.(val.Bool); !ok {
			return fmt.Errorf("'%s' must be of type 'bool'", name)
		}
	case "width":
		v, ok := pv.(val.Int)
		if !ok {
			return fmt.Errorf("'%s' property must be of type 'integer', current type '%s'", name, pv.Type())
		}
		if v.V < 0 {
			return fmt.Errorf("'width' property cannot have negative value (%d)", v.V)
		}
	default:
		return fmt.Errorf("unknown property name '%s'", name)
	}

	return nil
}
