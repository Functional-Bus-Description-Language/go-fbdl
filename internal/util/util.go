package util

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func IsBaseType(t string) bool {
	baseTypes := [...]string{
		"block", "bus", "config", "func", "mask", "param", "return", "status", "stream",
	}

	for i, _ := range baseTypes {
		if t == baseTypes[i] {
			return true
		}
	}

	return false
}

// IsValidProperty returns true if given property is valid for given base type.
func IsValidProperty(p string, t string) error {
	validProps := map[string][]string{
		"block":  []string{},
		"bus":    []string{"masters", "width"},
		"config": []string{"atomic", "default", "groups", "range", "once", "width"},
		"func":   []string{},
		"mask":   []string{"atomic", "default", "groups", "once", "width"},
		// TODO: Decide if "default" should be possible for param.
		// It creates some problems as not all programming languges support it.
		"param":  []string{"range", "width"},
		"return": []string{"width"},
		"status": []string{"atomic", "groups", "once", "width"},
		"stream": []string{},
	}

	if list, ok := validProps[t]; ok {
		for i, _ := range list {
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
		for i, _ := range list {
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
		"block":  []string{"block", "config", "func", "mask", "status", "stream"},
		"bus":    []string{"block", "config", "func", "mask", "status", "stream"},
		"config": []string{},
		"func":   []string{"param", "return"},
		"mask":   []string{},
		"param":  []string{},
		"status": []string{},
		"stream": []string{"param", "return"},
	}

	if list, ok := validTypes[ot]; ok {
		for i, _ := range list {
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

// IsValidQualifiedIdentifier returns an error if given qualified identifier is not valid.
// For example, if symbol name starts with lower case letter.
func IsValidQualifiedIdentifier(qi string) error {
	aux := strings.Split(qi, ".")
	pkg := aux[0]
	sym := aux[1]
	if unicode.IsUpper([]rune(sym)[0]) == false {
		return fmt.Errorf(
			"symbol '%s' imported from package '%s' starts with lower case letter",
			sym, pkg,
		)
	}

	return nil
}
