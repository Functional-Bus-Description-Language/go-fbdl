package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Status struct {
	Elem

	Atomic bool
	Groups []string
	Once   bool
	Width  int64

	Access access.Access
}
