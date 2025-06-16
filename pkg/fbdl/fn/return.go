package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Return struct {
	Func

	Width int64

	Access access.Access
}

func (r Return) Type() string { return "return" }
