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
