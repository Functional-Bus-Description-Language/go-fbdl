package prs

import (
	"fmt"
	"math"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// assertCall asserts that function name for given call is supported,
// and that the number of arguments for given call is valid.
func assertCall(c Call) error {
	validFuncNames := map[string]bool{
		"bool": true, "ceil": true, "floor": true, "log2": true, "log10": true,
	}

	validArgCount := map[string]int{
		"bool": 1, "ceil": 1, "floor": 1, "log2": 1, "log10": 1,
	}

	if ok := validFuncNames[c.funcName]; !ok {
		return fmt.Errorf("unknown function '%s'", c.funcName)
	}

	count := validArgCount[c.funcName]
	if len(c.args) != count {
		return fmt.Errorf(
			"function '%s' takes %d arguments, but %d were provided",
			c.funcName, count, len(c.args),
		)
	}

	return nil
}

func evalBool(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("bool argument evaluation: %v", err)
	}

	switch arg := arg.(type) {
	case val.Int:
		if arg != 0 {
			return val.Bool(true), nil
		}
		return val.Bool(false), nil
	}

	return nil, fmt.Errorf("invalid argument type '%s' for bool function", arg.Type())
}

func evalCeil(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("ceil argument evaluation: %v", err)
	}

	f := float64(0.0)

	switch arg := arg.(type) {
	case val.Int:
		return arg, nil
	case val.Float:
		f = float64(arg)
	}

	return val.Int(int64(math.Ceil(f))), nil
}

func evalFloor(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("floor argument evaluation: %v", err)
	}

	f := float64(0.0)

	switch arg := arg.(type) {
	case val.Int:
		return arg, nil
	case val.Float:
		f = float64(arg)
	}

	return val.Int(int64(math.Floor(f))), nil
}

func evalLog2(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("log2 argument evaluation: %v", err)
	}

	argType := "unknown"
	f := float64(0.0)

	switch arg := arg.(type) {
	case val.Int:
		argType = "int"
		f = float64(arg)
	case val.Float:
		argType = "float"
		f = float64(arg)
	}

	if argType != "int" && argType != "float" {
		return nil, fmt.Errorf("cannot evaluate log2 for argument of %s type", argType)
	}

	r := math.Log2(f)
	if r == float64(int64(r)) {
		return val.Int(int64(r)), nil
	}

	return val.Float(r), nil
}

func evalLog10(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("log10 argument evaluation: %v", err)
	}

	argType := "unknown"
	f := float64(0.0)

	switch arg := arg.(type) {
	case val.Int:
		argType = "int"
		f = float64(arg)
	case val.Float:
		argType = "float"
		f = float64(arg)
	}

	if argType != "int" && argType != "float" {
		return nil, fmt.Errorf("cannot evaluate log10 for argument of %s type", argType)
	}

	r := math.Log10(f)
	if r == float64(int64(r)) {
		return val.Int(int64(r)), nil
	}

	return val.Float(r), nil
}
