package elem

import (
	fbdl "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

type Groupable interface {
	fbdl.Element
	Groups() []string
}
