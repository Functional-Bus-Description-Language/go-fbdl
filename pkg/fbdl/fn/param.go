package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Param struct {
	Func

	Range types.Range
	Width int64

	Access types.Access
}

func (p Param) Type() string { return "param" }
