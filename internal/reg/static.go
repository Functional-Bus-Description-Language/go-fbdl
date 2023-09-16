package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regStatic registerifies Static functionality.
func regStatic(st *fn.Static, addr int64, gp *gap.Pool) int64 {
	if st.IsArray {
		return regStaticArray(st, addr, gp)
	} else {
		return regStaticSingle(st, addr, gp)
	}
}

func regStaticSingle(st *fn.Static, addr int64, gp *gap.Pool) int64 {
	/*
		var acs access.Access
		if g, ok := gp.GetSingle(st.Width, false); ok {
			acs = access.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width)
		} else {
	*/

	acs := access.MakeSingle(addr, 0, st.Width)
	addr += acs.GetRegCount()

	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Single{
			Addr:      acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = acs

	return addr
}

func regStaticArray(st *fn.Static, addr int64, gp *gap.Pool) int64 {
	var acs access.Access

	// TODO: In all below branches a potential gap can be added.
	if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = access.MakeArrayOneInReg(st.Count, addr, 0, st.Width)
	} else if st.Width <= busWidth/2 && st.Count%(busWidth/st.Width) == 0 {
		acs = access.MakeArrayNInReg(st.Count, addr, st.Width)
	} else if st.Width <= busWidth/2 {
		acs = access.MakeArrayNInRegMInEndReg(st.Count, addr, st.Width)
	} else {
		panic("unimplemented")
	}
	addr += acs.GetRegCount()

	st.Access = acs

	return addr
}
