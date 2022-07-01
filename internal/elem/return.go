package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Return represents return element.
type ret struct {
	Elem

	// Properties
	Groups []string
	Width  int64

	Access access.Access
}

// Return represents return element.
type Return struct {
	ret
}

func (r *Return) Type() string { return "return" }

func (r *Return) SetGroups(g []string) { r.ret.Groups = g }
func (r *Return) Groups() []string     { return r.ret.Groups }

func (r *Return) SetWidth(w int64) { r.ret.Width = w }
func (r *Return) Width() int64     { return r.ret.Width }

func (r *Return) SetAccess(a access.Access) { r.ret.Access = a }
func (r *Return) Access() access.Access     { return r.ret.Access }

func (r *Return) Hash() int64 {
	return 0
}
