// Package hash implements hash calculation for public types.
// The types are public, however their hash functions should not be public.
package hash

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/addrSpace"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

func write(buf io.Writer, data any) {
	err := binary.Write(buf, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}
}

func Hash(data any) uint32 {
	switch d := data.(type) {
	case addrSpace.AddrSpace:
		return hashAddrSpace(d)
	case access.Sizes:
		return hashAccessSizes(d)
	case access.Access:
		return hashAccessAccess(d)
	case *elem.Elem:
		return hashElem(d)
	case *elem.Block:
		return hashBlock(d)
	case *elem.Config:
		return hashConfig(d)
	case *elem.ConstContainer:
		return hashConstContainer(d)
	case *elem.Mask:
		return hashMask(d)
	case *elem.Proc:
		return hashProc(d)
	case *elem.Param:
		return hashParam(d)
	case *elem.Return:
		return hashReturn(d)
	case *elem.Static:
		return hashStatic(d)
	case *elem.Status:
		return hashStatus(d)
	case *elem.Stream:
		return hashStream(d)
	default:
		panic(
			fmt.Sprintf("Hash not implemented for %T\n", data),
		)
	}
}
