package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Static struct {
	Elem

	InitValue val.BitStr
	Groups    []string
	Width     int64

	Access access.Access
}
