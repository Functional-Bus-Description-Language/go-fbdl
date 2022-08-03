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
		"log2": true,
	}

	validArgCount := map[string]int{
		"log2": 1,
	}

	if ok := validFuncNames[c.funcName]; !ok {
		return fmt.Errorf("unknown function '%s'", c.funcName)
	}

	count := validArgCount[c.funcName]
	if len(c.args) != count {
		return fmt.Errorf(
			"function '%s' takes %d arguments, but %d were given",
			c.funcName, count, len(c.args),
		)
	}

	return nil
}

func evalLog2(c Call) (val.Value, error) {
	arg, err := c.args[0].Eval()
	if err != nil {
		return nil, fmt.Errorf("log2 argument evaluation: %v", err)
	}

	argType := "unknown"
	f := float64(0.0)

	switch arg.(type) {
	case val.Int:
		argType = "int"
		f = float64(arg.(val.Int))
	case val.Float:
		argType = "float"
		f = float64(arg.(val.Float))
	}

	if argType != "int" && argType != "float" {
		return nil, fmt.Errorf("cannot evaluate log2 for argument of %s type", argType)
	}

	r := math.Log2(f)
	if r == float64(int64(r)) {
		return val.Int(int64(r)), nil
	}

	panic("not yet implement, needs float point type")
}
