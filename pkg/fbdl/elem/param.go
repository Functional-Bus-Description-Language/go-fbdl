package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Param struct {
	Elem

	Groups []string
	Range  val.Range
	Width  int64

	Access access.Access
}

func (p *Param) GroupNames() []string { return p.Groups }
