package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
)

type Group struct {
	Func

	Virtual bool

	Consts cnst.Container

	Configs  []*Config
	Irqs     []*Irq
	Masks    []*Mask
	Params   []*Param
	Returns  []*Return
	Statics  []*Static
	Statuses []*Status
}

func (g Group) Type() string { return "group" }

// IsReadOnly returns true if group has only functionalities that are read-only.
//
// Irq group is read-only only if all irqs are clear on read and have no enable register.
func (g Group) IsReadOnly() bool {
	if len(g.Configs) > 0 || len(g.Masks) > 0 || len(g.Params) > 0 {
		return false
	}

	for _, irq := range g.Irqs {
		if irq.Clear == "Explicit" || irq.AddEnable {
			return false
		}
	}

	return true
}
