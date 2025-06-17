package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
)

func regIrq(irq *fn.Irq, addr int64, gp *gap.Pool) int64 {
	if irq.IsArray {
		return regIrqArray(irq, addr)
	}
	return regIrqSingle(irq, addr, gp)
}

// Irq is potentially put into a gap only if it has no enable register and is explicitly cleared.
// In all other cases saving some address space is not worth the extra complexity.
//
// As irqs are registerified as the last ones, the function doesn't add any gap to the pool.
func regIrqSingle(irq *fn.Irq, addr int64, gp *gap.Pool) int64 {
	if !irq.AddEnable && irq.Clear == "Explicit" {
		if g, ok := gp.GetSingle(1, true); ok {
			irq.Access = access.MakeSingle(g.Addr, g.StartBit, 1)
			clrAddr := g.Addr
			irq.ClearAddr = &clrAddr
			return addr
		}
	}

	// Handle all remaining cases.

	irq.Access = access.MakeSingle(addr, 0, 1)
	if irq.AddEnable {
		irq.EnableAccess = access.MakeSingle(addr, 1, 1)
		addr++
	}
	if irq.Clear == "Explicit" {
		clrAddr := addr
		irq.ClearAddr = &clrAddr
	}

	addr++

	return addr
}

func regIrqArray(irq *fn.Irq, addr int64) int64 {
	panic("unimplemented")
}
