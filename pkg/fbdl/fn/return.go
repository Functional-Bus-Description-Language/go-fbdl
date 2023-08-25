package fn

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Return struct {
	Func

	Groups []string
	Width  int64

	Access access.Access
}

func (r Return) Type() string { return "return" }

func (r *Return) GroupNames() []string { return r.Groups }
