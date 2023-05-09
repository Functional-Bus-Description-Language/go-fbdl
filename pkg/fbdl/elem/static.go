package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Static struct {
	Elem

	Groups     []string
	InitValue  val.BitStr
	ReadValue  val.BitStr
	ResetValue val.BitStr
	Width      int64

	Access access.Access
}
