package ins

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
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
		if v < 1 {
			return fmt.Errorf("'masters' property must be positive, current value (%d)", v)
		}
	case "range":
		v, ok := pv.(val.List)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "list", pv.Type())
		}
		if len(v) != 2 {
			return fmt.Errorf("length of 'range' property value list must equal 2, current length %d", len(v))
		}
		v0, ok := v[0].(val.Int)
		if !ok {
			return fmt.Errorf(
				"first value in 'range' property value list must be of type 'integer', current type '%s'", v[0].Type(),
			)
		}
		v1, ok := v[1].(val.Int)
		if !ok {
			return fmt.Errorf(
				"second value in 'range' property value list must be of type 'integer', current type '%s'", v[1].Type(),
			)
		}
		if v0 >= v1 {
			return fmt.Errorf("second value in 'range' property value list must be greater than first value")
		}
	case "width":
		v, ok := pv.(val.Int)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "integer", pv.Type())
		}
		if v < 0 {
			return fmt.Errorf("'width' property must be natural, current value (%d)", v)
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
