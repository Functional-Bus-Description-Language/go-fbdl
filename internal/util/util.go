package util

import (
	"fmt"
	"math"
)

func IsBaseType(t string) bool {
	baseTypes := [...]string{
		"block", "bus", "config", "irq", "mask", "memory", "param", "proc", "return", "static", "status", "stream",
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
		"block":  []string{"masters", "reset"},
		"bus":    []string{"masters", "reset", "width"},
		"config": []string{"atomic", "groups", "init-value", "range", "read-value", "reset-value", "width"},
		"irq":    []string{"add-enable", "clear", "enable-init-value", "enable-reset-value", "groups", "in-trigger", "out-trigger"},
		"mask":   []string{"atomic", "groups", "init-value", "read-value", "reset-value", "width"},
		"memory": []string{"access", "byte-write-enable", "read-latency", "size", "width"},
		"param":  []string{"groups", "range", "width"},
		"proc":   []string{"delay"},
		"return": []string{"groups", "width"},
		"static": []string{"groups", "init-value", "read-value", "reset-value", "width"},
		"status": []string{"atomic", "groups", "read-value", "width"},
		"stream": []string{"delay"},
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

	msg := "invalid property '%[1]s' for element of type '%[2]s', "

	if len(validProps[t]) == 0 {
		msg += "type '%[2]s' has no properties"
	} else {
		msg += "valid properties for element of type '%[2]s' are:"
		list := validProps[t]
		for i := range list {
			msg = msg + " '" + list[i] + "',"
		}
		msg = msg[:len(msg)-1]
	}

	msg = fmt.Sprintf(msg, p, t)

	return fmt.Errorf(msg)
}

// IsValidInnerType returns true if given inner type is valid for given outter type.
func IsValidInnerType(it string, ot string) bool {
	validTypes := map[string][]string{
		"block":  []string{"block", "config", "irq", "mask", "memory", "proc", "static", "status", "stream"},
		"bus":    []string{"block", "config", "irq", "mask", "memory", "proc", "static", "status", "stream"},
		"config": []string{},
		"irq":    []string{},
		"mask":   []string{},
		"memory": []string{},
		"param":  []string{},
		"proc":   []string{"param", "return"},
		"return": []string{},
		"static": []string{},
		"status": []string{},
		"stream": []string{"param", "return"},
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
