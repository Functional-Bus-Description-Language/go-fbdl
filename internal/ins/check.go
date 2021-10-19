package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/val"
)

// Check property value type, value and conflicting properties within element.
func checkProperty(name string, prop prs.Property) error {
	pv, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate expression: %v", err)
	}

	invalidTypeMsg := `'%s' property must be of type '%s', current type '%s'`

	switch name {
	case "atomic", "once":
		if _, ok := pv.(val.Bool); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "bool", pv.Type())
		}
	case "doc":
		if _, ok := pv.(val.Str); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
	case "masters":
		v, ok := pv.(val.Int)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "integer", pv.Type())
		}
		if v.V < 1 {
			return fmt.Errorf("'masters' property must be positive, current value (%d)", v.V)
		}
	case "width":
		v, ok := pv.(val.Int)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "integer", pv.Type())
		}
		if v.V < 0 {
			return fmt.Errorf("'width' property must be natural, current value (%d)", v.V)
		}
	default:
		msg := `checkProperty() for property '%s' not yet implemented`
		msg = fmt.Sprintf(msg, name)
		panic(msg)
	}

	return nil
}
