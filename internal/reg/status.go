package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regAtomicStatus registerifies an atomic Status functionality.
func regAtomicStatus(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	if st.IsArray {
		return regAtomicStatusArray(st, addr, gp)
	}
	return regAtomicStatusSingle(st, addr, gp)
}

func regAtomicStatusSingle(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	acs := access.MakeSingle(addr, 0, st.Width)
	addr += acs.GetRegCount()

	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.GetEndAddr(),
			EndAddr:   acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = acs

	return addr
}

func regAtomicStatusArray(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	var acs access.Access

	// TODO: In all below branches a potential gap can be added.
	if st.Count*st.Width <= busWidth {
		acs = access.MakeArrayOneReg(st.Count, addr, 0, st.Width)
	} else if busWidth/2 < st.Width && st.Width <= busWidth {
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

// regNonAtomicStatus registerifies a Non-Atomic Status functionality.
func regNonAtomicStatus(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	if st.IsArray {
		return regNonAtomicStatusArray(st, addr, gp)
	}
	return regNonAtomicStatusSingle(st, addr, gp)
}

func regNonAtomicStatusSingle(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	// TODO: Check if there is gap at the end that can be utilized.
	acs := access.MakeSingle(addr, 0, st.Width)
	addr += acs.GetRegCount()

	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.GetEndAddr(),
			EndAddr:   acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = acs

	return addr
}

func regNonAtomicStatusArray(st *fn.Status, addr int64, gp *gap.Pool) int64 {
	var acs access.Access

	if st.Count*st.Width <= busWidth {
		acs = access.MakeArrayOneReg(st.Count, addr, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth/2 < st.Width && st.Width <= busWidth {
		acs = access.MakeArrayOneInReg(st.Count, addr, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width == 0 || st.Count <= busWidth/st.Width || st.Width < busWidth/2 {
		acs = access.MakeArrayNInReg(st.Count, addr, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("unimplemented")
	}
	addr += acs.GetRegCount()

	st.Access = acs

	return addr
}
