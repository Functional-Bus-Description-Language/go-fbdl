package util

import (
	"fmt"
	"math"
)

const RepoIssueUrl = "https://github.com/Functional-Bus-Description-Language/go-fbdl/issues"

func IsBaseType(t string) bool {
	baseTypes := [...]string{
		"blackbox", "block", "bus", "config", "group", "irq", "mask", "param", "proc", "return", "static", "status", "stream",
	}

	for i := range baseTypes {
		if t == baseTypes[i] {
			return true
		}
	}

	return false
}

// IsValidProperty returns true if given property is valid for given base type.
func IsValidProperty(p string, t string) error {
	validProps := map[string][]string{
		"blackbox": []string{"size"},
		"block":    []string{"align", "masters", "reset"},
		"bus":      []string{"align", "masters", "reset", "width"},
		"config":   []string{"atomic", "init-value", "range", "read-value", "reset-value", "width"},
		"group":    []string{"virtual"},
		"irq":      []string{"add-enable", "clear", "enable-init-value", "enable-reset-value", "in-trigger", "out-trigger"},
		"mask":     []string{"atomic", "init-value", "read-value", "reset-value", "width"},
		"param":    []string{"range", "width"},
		"proc":     []string{"delay"},
		"return":   []string{"width"},
		"static":   []string{"init-value", "read-value", "reset-value", "width"},
		"status":   []string{"atomic", "read-value", "width"},
		"stream":   []string{"delay"},
	}

	if list, ok := validProps[t]; ok {
		for i := range list {
			if p == list[i] {
				return nil
			}
		}
	} else {
		panic(fmt.Sprintf("invalid base type '%s'", t))
	}

	msg := "invalid property '%[1]s' for %[2]s functionality"

	if len(validProps[t]) == 0 {
		msg += "type '%[2]s' has no properties"
	} else {
		msg += "\nvalid properties for %[2]s are:"
		list := validProps[t]
		for i := range list {
			msg = msg + " '" + list[i] + "',"
		}
		msg = msg[:len(msg)-1]
	}

	msg = fmt.Sprintf(msg, p, t)

	return fmt.Errorf("%s", msg)
}

// IsValidInnerType returns true if given inner type is valid for given outter type.
func IsValidInnerType(it string, ot string) bool {
	validTypes := map[string][]string{
		"blackbox": []string{},
		"block":    []string{"blackbox", "block", "config", "group", "irq", "mask", "proc", "static", "status", "stream"},
		"bus":      []string{"blackbox", "block", "config", "group", "irq", "mask", "proc", "static", "status", "stream"},
		"config":   []string{},
		"group":    []string{"config", "irq", "mask", "param", "return", "static", "status"},
		"irq":      []string{},
		"mask":     []string{},
		"param":    []string{},
		"proc":     []string{"param", "return"},
		"return":   []string{},
		"static":   []string{},
		"status":   []string{},
		"stream":   []string{"param", "return"},
	}

	if list, ok := validTypes[ot]; ok {
		for i := range list {
			if it == list[i] {
				return true
			}
		}
	} else {
		panic("should never happen")
	}

	return false
}

func AlignToPowerOf2(n int64) int64 {
	return int64(math.Pow(2, math.Ceil(math.Log2(float64(n)))))
}
