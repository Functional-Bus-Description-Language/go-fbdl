package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Irq struct {
	Elem

	AddEnable        bool
	Clear            string
	EnableInitValue  val.BitStr
	EnableResetValue val.BitStr
	Groups           []string
	InTrigger        string
	OutTrigger       string

	Access access.Access
}

func (i *Irq) GroupNames() []string { return i.Groups }
