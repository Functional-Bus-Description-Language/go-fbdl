package ins

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/tok"
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
	case "access":
		v, ok := pv.(val.Str)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
		if v != "Read Write" && v != "Read Only" && v != "Write Only" {
			return fmt.Errorf(
				"'access' property must be \"Read Write\", \"Read Only\" or \"Write Only\", current value (%q)", v,
			)
		}
	case "add-enable", "atomic", "byte-write-enable":
		if _, ok := pv.(val.Bool); !ok {
			return fmt.Errorf(invalidTypeMsg, name, "bool", pv.Type())
		}
	case "clear":
		v, ok := pv.(val.Str)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
		if v != "Explicit" && v != "On Read" {
			return fmt.Errorf(
				"'clear' property must be \"Explicit\" or \"On Read\", current value (%q)", v,
			)
		}
	case "delay":
		switch pv.(type) {
		case val.Time:
			break
		default:
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "time", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "enable-init-value", "enable-reset-value", "init-value", "read-value", "reset-value":
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
	case "in-trigger", "out-trigger":
		v, ok := pv.(val.Str)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
		if v != "Edge" && v != "Level" {
			return tok.Error{
				Msg:  fmt.Sprintf("'%s' property must be \"Edge\" or \"Level\", current value %q", name, v),
				Toks: []tok.Token{prop.ValueTok},
			}
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
				return tok.Error{
					Msg:  fmt.Sprintf("'range' property value must be natural, value %d is negative", v),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
		case val.List:
			if len(v) == 0 {
				return tok.Error{
					Msg:  "empty 'range' property value list",
					Toks: []tok.Token{prop.ValueTok},
				}
			}
			if len(v)%2 != 0 {
				return tok.Error{
					Msg:  fmt.Sprintf("length of 'range' property value list must be even, current length %d", len(v)),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
			lower := true
			var lowerBound, upperBound int64
			for i, bound := range v {
				bound_val, ok := bound.(val.Int)
				if !ok {
					return tok.Error{
						Msg: fmt.Sprintf(
							"all values in 'range' property list must be of type 'integer', value with index %d is of type '%s'",
							i, bound.Type(),
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
				if bound_val < 0 {
					return tok.Error{
						Msg: fmt.Sprintf(
							"'range' property value must be natural, value with index %d is negative %d",
							i, bound_val,
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
				if lower {
					lowerBound = int64(bound_val)
					lower = false
				} else {
					upperBound = int64(bound_val)
					lower = true
				}
				if lower && lowerBound > upperBound {
					return tok.Error{
						Msg: fmt.Sprintf(
							"'range' property list, lower bound with index %d (%d) is greater than upper bound with index %d (%d)",
							i-1, lowerBound, i, upperBound,
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
			}
		default:
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "integer or [integer]", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "reset":
		v, ok := pv.(val.Str)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "string", pv.Type())
		}
		reset := string(v)
		if reset != "Sync" && reset != "Async" {
			return tok.Error{
				Msg:  fmt.Sprintf("'reset' property must be \"Sync\" or \"Async\", current value %q", reset),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "read-latency", "size", "width":
		v, ok := pv.(val.Int)
		if !ok {
			return fmt.Errorf(invalidTypeMsg, name, "integer", pv.Type())
		}
		if v < 0 {
			return tok.Error{
				Msg:  fmt.Sprintf("'%s' property must be natural, current value %d", prop.Name, v),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	default:
		panic(fmt.Sprintf("checkProp() for property '%s' not yet implemented", name))
	}

	return nil
}
