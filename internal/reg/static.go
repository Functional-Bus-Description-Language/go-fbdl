package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
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
		var acs types.Access
		if g, ok := gp.Single(st.Width, false); ok {
			acs = types.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width)
		} else {
	*/

	acs := types.MakeSingleAccess(addr, 0, st.Width)
	addr += acs.RegCount

	if acs.EndBit < busWidth-1 {
		gp.Add(gap.Single{
			Addr:      acs.EndAddr,
			StartBit:  acs.EndBit + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = acs

	return addr
}

func regStaticArray(st *fn.Static, addr int64, gp *gap.Pool) int64 {
	var acs types.Access

	// TODO: In all below branches a potential gap can be added.
	if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = types.MakeArrayOneInRegAccess(st.Count, addr, 0, st.Width)
	} else if st.Width <= busWidth/2 && st.Count%(busWidth/st.Width) == 0 {
		acs = types.MakeArrayNInRegAccess(st.Count, addr, st.Width)
	} else if st.Width <= busWidth/2 {
		acs = types.MakeArrayNInRegMInEndRegAccess(st.Count, addr, st.Width)
	} else {
		panic("unimplemented")
	}
	addr += acs.RegCount

	st.Access = acs

	return addr
}
