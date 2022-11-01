package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Block struct {
	Elem

	Masters int64
	Width   int64

	ConstContainer

	Configs   []*Config
	Funcs     []*Func
	Masks     []*Mask
	Statics   []*Static
	Statuses  []*Status
	Streams   []*Stream
	Subblocks []*Block

	Sizes     access.Sizes
	AddrSpace access.AddrSpace
}
