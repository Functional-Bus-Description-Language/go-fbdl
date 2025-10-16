package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

type Return struct {
	Func

	Width int64

	Access types.Access
}

func (r Return) Type() string { return "return" }
