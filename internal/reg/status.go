package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regAtomicStatus registerifies an atomic Status functionality.
func regAtomicStatus(st *fn.Status, addr *address, gp *gap.Pool) {
	if st.IsArray {
		regAtomicStatusArray(st, addr, gp)
	}
	regAtomicStatusSingle(st, addr, gp)
}

func regAtomicStatusSingle(st *fn.Status, addr *address, gp *gap.Pool) {
	var acs access.Access

	if g, ok := gp.GetSingle(st.Width, false); ok {
		acs = access.MakeSingle(g.Addr, g.StartBit, st.Width)
	} else {
		acs = access.MakeSingle(addr.value, 0, st.Width)
		addr.inc(acs.GetRegCount())
	}

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

func regAtomicStatusArray(st *fn.Status, addr *address, gp *gap.Pool) {
	var acs access.Access

	// TODO: In all below branches a potential gap can be added.
	if st.Count*st.Width <= busWidth {
		acs = access.MakeArrayOneReg(st.Count, addr.value, 0, st.Width)
	} else if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = access.MakeArrayOneInReg(st.Count, addr.value, 0, st.Width)
	} else if st.Width <= busWidth/2 && st.Count%(busWidth/st.Width) == 0 {
		acs = access.MakeArrayNInReg(st.Count, addr.value, st.Width)
	} else if st.Width <= busWidth/2 {
		acs = access.MakeArrayNInRegMInEndReg(st.Count, addr.value, st.Width)
	} else if st.Width > busWidth {
		acs = access.MakeArrayOneInNRegs(st.Count, addr.value, st.Width)
	} else {
		panic("unimplemented")
	}

	addr.inc(acs.GetRegCount())

	st.Access = acs
}

// regNonAtomicStatus registerifies a Non-Atomic Status functionality.
func regNonAtomicStatus(st *fn.Status, addr *address, gp *gap.Pool) {
	if st.IsArray {
		regNonAtomicStatusArray(st, addr, gp)
	}
	regNonAtomicStatusSingle(st, addr, gp)
}

func regNonAtomicStatusSingle(st *fn.Status, addr *address, gp *gap.Pool) {
	var acs access.Access

	if g, ok := gp.GetSingle(st.Width, false); ok {
		acs = access.MakeSingle(g.Addr, g.StartBit, st.Width)
	} else {
		acs = access.MakeSingle(addr.value, 0, st.Width)
		addr.inc(acs.GetRegCount())
	}

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

func regNonAtomicStatusArray(st *fn.Status, addr *address, gp *gap.Pool) {
	var acs access.Access

	if st.Count*st.Width <= busWidth {
		acs = access.MakeArrayOneReg(st.Count, addr.value, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = access.MakeArrayOneInReg(st.Count, addr.value, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width == 0 || st.Count <= busWidth/st.Width || st.Width < busWidth/2 {
		acs = access.MakeArrayNInReg(st.Count, addr.value, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if st.Width > busWidth {
		acs = access.MakeArrayOneInNRegs(st.Count, addr.value, st.Width)
	} else {
		panic("unimplemented")
	}
	addr.inc(acs.GetRegCount())

	st.Access = acs
}
