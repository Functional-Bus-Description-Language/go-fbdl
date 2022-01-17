package util

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func IsBaseType(t string) bool {
	baseTypes := [...]string{
		"block", "bus", "config", "func", "mask", "param", "return", "status",
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

	msg := "invalid property '%s' for element of type '%s', " +
		"valid properties for element of type '%[2]s' are:"

	list := validProps[t]
	for i, _ := range list {
		msg = msg + " '" + list[i] + "',"
	}

	msg = fmt.Sprintf(msg, p, t)
	msg = msg[:len(msg)-1]

	return fmt.Errorf(msg)
}

// IsValidType returns true if given inner type is valid for given outter type.
func IsValidType(ot string, it string) bool {
	validTypes := map[string][]string{
		"block":  []string{"block", "config", "func", "mask", "status"},
		"bus":    []string{"block", "config", "func", "mask", "status"},
		"config": []string{},
		"func":   []string{"param", "return"},
		"mask":   []string{},
		"param":  []string{},
		"status": []string{},
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
