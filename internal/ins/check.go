package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Check property value type and value.
func checkProp(prop prs.Prop) error {
	pv, err := prop.Value.Eval()
	if err != nil {
		return fmt.Errorf("cannot evaluate expression: %v", err)
	}

	invalidTypeMsg := `'%s' property must be of type '%s', current type '%s'`

	name := prop.Name

	switch name {
	case "atomic":
		if _, ok := pv.(val.Bool); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "bool", pv.Type())
		}
	case "delay":
		switch pv.(type) {
		case val.Time:
			break
		default:
			return fmt.Errorf(invalidTypeMsg, name, "time", pv.Type())
		}
	case "init-value":
		switch pv.(type) {
		case val.Int, val.BitStr:
			break
		default:
			return fmt.Errorf(invalidTypeMsg, name, "integer or bit string", pv.Type())
		}
	case "groups":
		switch v := pv.(type) {
		case val.Str:
			break
		case val.List:
			groups := v
			if len(groups) == 0 {
				return fmt.Errorf("groups list of length 0 makes no sense")
			}
			for i, v := range groups {
				if _, ok := v.(val.Str); !ok {
					return fmt.Errorf("all values in groups list must be of type 'string', item %d is of type '%s'", i, v.Type())
				}
			}
			groupsMap := make(map[string]int)
			for i, v := range groups {
				g := v.(val.Str)
				if firstIdx, exists := groupsMap[string(g)]; exists {
					return fmt.Errorf("duplicated %q in groups list, first item %d, second item %d", g, firstIdx, i)
				}
				groupsMap[string(g)] = i
			}
		default:
			return fmt.Errorf(invalidTypeMsg, name, "string or [string]", pv.Type())
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
		switch v := pv.(type) {
		case val.Int:
			if v < 0 {
				return fmt.Errorf("'range' property value must be natural, value %d is negative", v)
			}
		case val.List:
			if len(v) == 0 {
				return fmt.Errorf("'range' property value list must not be empty")
			}
			if len(v)%2 != 0 {
				return fmt.Errorf("length of 'range' property value list must be even, current length %d", len(v))
			}
			lower := true
			var lowerBound, upperBound int64
			for i, bound := range v {
				bound_val, ok := bound.(val.Int)
				if !ok {
					return fmt.Errorf(
						"all values in 'range' property list must be of type 'integer', value with index %d is of type '%s'",
						i, bound.Type(),
					)
				}
				if bound_val < 0 {
					return fmt.Errorf(
						"'range' property value must be natural, value with index %d is negative (%d)", i, bound_val,
					)
				}
				if lower {
					lowerBound = int64(bound_val)
					lower = false
				} else {
					upperBound = int64(bound_val)
					lower = true
				}
				if lower && lowerBound > upperBound {
					return fmt.Errorf(
						"'range' property list, lower bound with index %d (%d) is greater than upper bound with index %d (%d)",
						i-1, lowerBound, i, upperBound,
					)
				}
			}
		default:
			return fmt.Errorf(invalidTypeMsg, name, "integer or [integer]", pv.Type())
		}
	case "reset":
		v, ok := pv.(val.Str)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
		reset := string(v)
		if reset != "Sync" && reset != "Async" {
			return fmt.Errorf("'reset' property must be \"Sync\" or \"Async\", current value (%q)", reset)
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
		panic(fmt.Sprintf("checkProp() for property '%s' not yet implemented", name))
	}

	return nil
}

/*
TODO: Implement it somewhere.
func checkPropConflict(elem *Element, prop string) error {
	msg := `cannot set '%s' property, because '%s' property is already set in one of ancestor types`

	if _, ok := elem.Props["width"]; ok {
		if prop == "range" {
			return fmt.Errorf(msg, "range", "width")
		}
	}

	if _, ok := elem.Props["range"]; ok {
		if prop == "width" {
			return fmt.Errorf(msg, "width", "range")
		}
	}

	return nil
}
*/
