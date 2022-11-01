package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Return struct {
	Elem

	Groups []string
	Width  int64

	Access access.Access
}
