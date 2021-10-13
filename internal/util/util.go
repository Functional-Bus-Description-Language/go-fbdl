package util

func IsBaseType(t string) bool {
	base_types := [...]string{"block", "bus", "config", "func", "mask", "param", "status"}

	for i, _ := range base_types {
		if t == base_types[i] {
			return true
		}
	}

	return false
}

// IsValidProperty returns true if given property is valid for given type.
func IsValidProperty(t string, p string) bool {
	validProps := map[string][]string{
		"block":  []string{"doc"},
		"bus":    []string{"doc", "masters", "width"},
		"config": []string{"atomic", "default", "doc", "groups", "range", "once", "width"},
		"func":   []string{"doc"},
		"mask":   []string{"atomic", "default", "doc", "groups", "width"},
		"param":  []string{"default", "doc", "range", "width"},
		"status": []string{"atomic", "doc", "groups", "once", "width"},
	}

	if list, ok := validProps[t]; ok {
		for i, _ := range list {
			if p == list[i] {
				return true
			}
		}
	} else {
		panic("should never happen")
	}

	return false
}

// IsValidType returns true if given inner type is valid for given outter type.
func IsValidType(ot string, it string) bool {
	validTypes := map[string][]string{
		"block":  []string{"block", "config", "func", "mask", "status"},
		"bus":    []string{"block", "config", "func", "mask", "status"},
		"config": []string{},
		"func":   []string{"param"},
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
