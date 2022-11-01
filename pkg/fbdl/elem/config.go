package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

type Config struct {
	Elem

	Atomic  bool
	Default val.BitStr
	Groups  []string
	//Range   Range
	Once  bool
	Width int64

	Access access.Access
}
