package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type st struct {
	Elem

	// Properties
	Atomic  bool
	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64

	Access access.Access
}

// Status represents status element.
type Status struct {
	st
}

func (s *Status) Type() string { return "status" }

func (c *Status) SetAtomic(a bool) { c.st.Atomic = a }
func (c *Status) Atomic() bool     { return c.st.Atomic }

func (c *Status) SetDefault(d val.BitStr) { c.st.Default = d }
func (c *Status) Default() val.BitStr     { return c.st.Default }

func (c *Status) SetGroups(g []string) { c.st.Groups = g }
func (c *Status) Groups() []string     { return c.st.Groups }

func (c *Status) SetOnce(a bool) { c.st.Once = a }
func (c *Status) Once() bool     { return c.st.Once }

func (c *Status) SetWidth(w int64) { c.st.Width = w }
func (c *Status) Width() int64     { return c.st.Width }

func (c *Status) SetAccess(a access.Access) { c.st.Access = a }
func (c *Status) Access() access.Access     { return c.st.Access }

// HasDecreasingAccessOrder returns true if status must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic statuses.
// If the end register is wider, then starting reading from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (s *Status) HasDecreasingAccessOrder() bool {
	if !s.st.Atomic {
		return false
	}

	if asc, ok := s.st.Access.(access.SingleContinuous); ok {
		if asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}

func (s *Status) Hash() int64 {
	return 0
}
