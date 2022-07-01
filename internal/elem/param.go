package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type param struct {
	Elem

	// Properties
	// TODO: Should Default be supported for param?
	// Default BitStr
	//Range Range
	Groups []string
	Width  int64

	Access access.Access
}

func (p *Param) Type() string { return "param" }

func (p *Param) SetGroups(g []string) { p.param.Groups = g }
func (p *Param) Groups() []string     { return p.param.Groups }

func (p *Param) SetWidth(w int64) { p.param.Width = w }
func (p *Param) Width() int64     { return p.param.Width }

func (p *Param) SetAccess(a access.Access) { p.param.Access = a }
func (p *Param) Access() access.Access     { return p.param.Access }

// Param represents param element.
type Param struct {
	param
}

func (p *Param) Hash() int64 {
	return 0
}
