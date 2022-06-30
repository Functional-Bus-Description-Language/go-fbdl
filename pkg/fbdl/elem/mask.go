package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// Mask represents mask element.
type Mask struct {
	Elem

	Access access.Access

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64
}

func (m *Mask) Type() string { return "mask" }

func (m *Mask) Hash() int64 {
	return 0
}

func (m *Mask) GroupNames() []string {
	return m.Groups
}
