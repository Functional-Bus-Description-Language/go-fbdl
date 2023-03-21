package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Status struct {
	Elem

	Atomic bool
	Groups []string
	Width  int64

	Access access.Access
}

func (s *Status) GroupNames() []string { return s.Groups }
