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
		return err
	}

	invalidTypeMsg := `%s property must be of type %s, current type %s`

	name := prop.Name

	switch name {
	case "access":
		v, ok := pv.(val.Str)
		if !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "string", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		if v != "Read Write" && v != "Read Only" && v != "Write Only" {
			return tok.Error{
				Msg: fmt.Sprintf(
					"access property must be \"Read Write\", \"Read Only\" or \"Write Only\", current value %q", v,
				),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "add-enable", "atomic", "byte-write-enable":
		if _, ok := pv.(val.Bool); !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "bool", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "clear":
		v, ok := pv.(val.Str)
		if !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "string", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		if v != "Explicit" && v != "On Read" {
			return tok.Error{
				Msg:  fmt.Sprintf("clear property must be \"Explicit\" or \"On Read\", current value %q", v),
				Toks: []tok.Token{prop.ValueTok},
			}
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
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "integer or bit string", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
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
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "string", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		if v != "Edge" && v != "Level" {
			return tok.Error{
				Msg:  fmt.Sprintf("%s property must be \"Edge\" or \"Level\", current value %q", name, v),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "masters":
		v, ok := pv.(val.Int)
		if !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "integer", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		if v < 1 {
			return tok.Error{
				Msg:  fmt.Sprintf("masters property must be positive, current value %d", v),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "range":
		switch v := pv.(type) {
		case val.Int:
			if v < 0 {
				return tok.Error{
					Msg:  fmt.Sprintf("range property value must be natural, value %d is negative", v),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
		case val.Range:
			if v.L < 0 {
				return tok.Error{
					Msg:  fmt.Sprintf("negative range left bound %d", v.L),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
			if v.R < 0 {
				return tok.Error{
					Msg:  fmt.Sprintf("negative range right bound %d", v.R),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
			if v.L > v.R {
				return tok.Error{
					Msg:  fmt.Sprintf("range left bound greater than right bound, %d > %d", v.L, v.R),
					Toks: []tok.Token{prop.ValueTok},
				}
			}
		case val.List:
			if len(v) == 0 {
				return tok.Error{
					Msg:  "empty range property value list",
					Toks: []tok.Token{prop.ValueTok},
				}
			}

			for i, rng := range v {
				r, ok := rng.(val.Range)
				if !ok {
					return tok.Error{
						Msg: fmt.Sprintf(
							"all values in range property list must be of type range, value with index %d is of type %s",
							i, rng.Type(),
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}

				if r.L < 0 {
					return tok.Error{
						Msg: fmt.Sprintf(
							"negative range left bound %d in value with index %d",
							r.L, i,
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
				if r.R < 0 {
					return tok.Error{
						Msg: fmt.Sprintf(
							"negative range right bound %d in value with index %d",
							r.R, i,
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
				if r.L > r.R {
					return tok.Error{
						Msg: fmt.Sprintf(
							"range left bound greater than right bound in value with index %d, %d > %d",
							i, r.L, r.R,
						),
						Toks: []tok.Token{prop.ValueTok},
					}
				}
			}
		default:
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "integer, range or [range]", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "reset":
		v, ok := pv.(val.Str)
		if !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "string", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		reset := string(v)
		if reset != "Sync" && reset != "Async" {
			return tok.Error{
				Msg:  fmt.Sprintf("reset property must be \"Sync\" or \"Async\", current value %q", reset),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	case "read-latency", "size", "width":
		v, ok := pv.(val.Int)
		if !ok {
			return tok.Error{
				Msg:  fmt.Sprintf(invalidTypeMsg, name, "integer", pv.Type()),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
		if v < 0 {
			return tok.Error{
				Msg:  fmt.Sprintf("%s property must be natural, current value %d", prop.Name, v),
				Toks: []tok.Token{prop.ValueTok},
			}
		}
	default:
		panic(fmt.Sprintf("checkProp() for property '%s' not yet implemented", name))
	}

	return nil
}
