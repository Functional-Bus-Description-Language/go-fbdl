// Package hash implements hash calculation for public types.
// The types are public, however their hash functions should not be public.
package hash

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/cnst"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"
)

func write(buf io.Writer, data any) {
	var err error

	if str, ok := data.(string); ok {
		_, err = buf.Write([]byte(str))
	} else {
		err = binary.Write(buf, binary.LittleEndian, data)
	}

	if err != nil {
		panic(err)
	}
}

func Hash(data any) uint32 {
	switch d := data.(type) {
	case types.Range:
		return hashRange(d)
	case access.Sizes:
		return hashAccessSizes(d)
	case access.Access:
		return hashAccessAccess(d)
	case *fn.Func:
		return hashFunc(d)
	case *fn.Blackbox:
		return hashBlackbox(d)
	case *fn.Block:
		return hashBlock(d)
	case *fn.Config:
		return hashConfig(d)
	case *cnst.Container:
		return hashConstContainer(d)
	case *fn.Irq:
		return hashIrq(d)
	case *fn.Mask:
		return hashMask(d)
	case *fn.Proc:
		return hashProc(d)
	case *fn.Param:
		return hashParam(d)
	case *fn.Return:
		return hashReturn(d)
	case *fn.Static:
		return hashStatic(d)
	case *fn.Status:
		return hashStatus(d)
	case *fn.Stream:
		return hashStream(d)
	default:
		panic(
			fmt.Sprintf("Hash not implemented for %T\n", data),
		)
	}
}
