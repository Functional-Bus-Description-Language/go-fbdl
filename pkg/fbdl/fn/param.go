package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Param struct {
	Func

	Range val.Range
	Width int64

	Access access.Access
}

func (p Param) Type() string { return "param" }
