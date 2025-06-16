package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Irq struct {
	Func

	AddEnable        bool
	Clear            string
	EnableInitValue  val.BitStr
	EnableResetValue val.BitStr
	InTrigger        string
	OutTrigger       string

	Access access.Access
}

func (i Irq) Type() string { return "irq" }
