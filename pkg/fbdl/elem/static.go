package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Static struct {
	Elem

	Default val.BitStr
	Groups  []string
	Once    bool
	Width   int64

	Access access.Access
}
