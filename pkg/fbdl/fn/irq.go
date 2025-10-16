package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Irq struct {
	Func

	AddEnable        bool
	Clear            string
	EnableInitValue  types.BitStr
	EnableResetValue types.BitStr
	InTrigger        string
	OutTrigger       string

	Access       access.Access
	EnableAccess access.Access // Access to the irq enable register

	// Address that must be written to generate a strobe signal for explicit clear.
	// The outer user logic must correctly handle the strobe clear signal.
	// Otherwise, the irq will not be cleared.
	ClearAddr *int64
}

func (i Irq) Type() string { return "irq" }
