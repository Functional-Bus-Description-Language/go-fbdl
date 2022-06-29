package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// Status represents status element.
type Status struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64
}

// HasDecreasingAccessOrder returns true if status must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic statuses.
// If the end register is wider, then starting reading from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (s *Status) HasDecreasingAccessOrder() bool {
	if !s.Atomic {
		return false
	}

	if asc, ok := s.Access.(access.SingleContinuous); ok {
		if asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}

func (s *Status) Hash() int64 {
	return 0
}
