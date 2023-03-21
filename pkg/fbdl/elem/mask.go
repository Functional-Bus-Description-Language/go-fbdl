package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Mask struct {
	Elem

	Atomic    bool
	InitValue val.BitStr
	Groups    []string
	Width     int64

	Access access.Access
}

func (m *Mask) GroupNames() []string { return m.Groups }
