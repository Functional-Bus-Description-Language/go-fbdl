package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regStatic registerifies Static functionality.
func regStatic(st *fn.Static, addr *address, gp *gap.Pool) {
	if st.IsArray {
		regStaticArray(st, addr, gp)
	} else {
		regStaticSingle(st, addr, gp)
	}
}

func regStaticSingle(st *fn.Static, addr *address, gp *gap.Pool) {
	/*
		var acs access.Access
		if g, ok := gp.GetSingle(st.Width, false); ok {
			acs = access.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width)
		} else {
	*/

	acs := access.MakeSingle(addr.value, 0, st.Width)
	addr.inc(acs.GetRegCount())

	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Single{
			Addr:      acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = acs
}

func regStaticArray(st *fn.Static, addr *address, gp *gap.Pool) {
	var acs access.Access

	// TODO: In all below branches a potential gap can be added.
	if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = access.MakeArrayOneInReg(st.Count, addr.value, 0, st.Width)
	} else if st.Width <= busWidth/2 && st.Count%(busWidth/st.Width) == 0 {
		acs = access.MakeArrayNInReg(st.Count, addr.value, st.Width)
	} else if st.Width <= busWidth/2 {
		acs = access.MakeArrayNInRegMInEndReg(st.Count, addr.value, st.Width)
	} else {
		panic("unimplemented")
	}
	addr.inc(acs.GetRegCount())

	st.Access = acs
}
