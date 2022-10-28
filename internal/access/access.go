package access

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

type Access interface {
	access.Access

	StartAddr() int64
	EndAddr() int64

	EndBit() int64

	Hash() uint32
}

type Single interface {
	Access

	access.Single
}
