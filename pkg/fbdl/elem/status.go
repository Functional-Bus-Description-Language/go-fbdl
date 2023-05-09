package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Status struct {
	Elem

	Atomic    bool
	Groups    []string
	ReadValue val.BitStr
	Width     int64

	Access access.Access
}

func (s *Status) GroupNames() []string { return s.Groups }
