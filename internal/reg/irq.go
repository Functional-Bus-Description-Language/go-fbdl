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

func regIrqSingle(irq *fn.Irq, addr int64, gp *gap.Pool) int64 {
	// Irq can be put into a gap only if it is explicitly cleared.
	if irq.Clear == "Explicit" {
		var acs access.Access
		if g, ok := gp.GetSingle(1, true); ok {
			acs = access.MakeSingle(g.Addr, g.StartBit, 1)
		} else {
			acs = access.MakeSingle(addr, 0, 1)
			addr++
		}
		irq.Access = acs

		if acs.EndBit() < busWidth-1 {
			gp.Add(gap.Single{
				Addr:      acs.EndAddr(),
				StartBit:  acs.EndBit() + 1,
				EndBit:    busWidth - 1,
				WriteSafe: false,
			})
		}
	} else {
		irq.Access = access.MakeSingle(addr, 0, 1)
		addr++
	}

	return addr
}

func regIrqArray(irq *fn.Irq, addr int64) int64 {
	panic("unimplemented")
}
