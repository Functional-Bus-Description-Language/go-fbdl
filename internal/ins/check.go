package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl"
)

// Check property value type and value.
func checkProperty(name string, prop prs.Property) error {
	pv, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate expression: %v", err)
	}

	invalidTypeMsg := `'%s' property must be of type '%s', current type '%s'`

	switch name {
	case "atomic", "once":
		if _, ok := pv.(fbdl.Bool); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "bool", pv.Type())
		}
	case "doc":
		if _, ok := pv.(fbdl.Str); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
	case "masters":
		v, ok := pv.(fbdl.Int)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "integer", pv.Type())
		}
		if v.V < 1 {
			return fmt.Errorf("'masters' property must be positive, current value (%d)", v.V)
		}
	case "range":
		v, ok := pv.(fbdl.List)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "list", pv.Type())
		}
		if len(v.V) != 2 {
			return fmt.Errorf("length of 'range' property value list must equal 2, current length %d", len(v.V))
		}
		v0, ok := v.V[0].(fbdl.Int)
		if !ok {
			return fmt.Errorf(
				"first value in 'range' property value list must be of type 'integer', current type '%s'", v.V[0].Type(),
			)
		}
		v1, ok := v.V[1].(fbdl.Int)
		if !ok {
			return fmt.Errorf(
				"second value in 'range' property value list must be of type 'integer', current type '%s'", v.V[1].Type(),
			)
		}
		if v0.V >= v1.V {
			return fmt.Errorf("second value in 'range' property value list must be greater than first value")
		}
	case "width":
		v, ok := pv.(fbdl.Int)
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

func checkPropertyConflict(elem *Element, prop string) error {
	msg := `cannot set '%s' property, because '%s' property is already set in one of ancestor types`
	if _, ok := elem.Properties["width"]; ok {
		if prop == "range" {
			return fmt.Errorf(msg, "range", "width")
		}
	}

	return nil
}
