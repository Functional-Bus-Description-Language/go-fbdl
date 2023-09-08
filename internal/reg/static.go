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
		var a access.Access
		if g, ok := gp.GetSingle(st.Width, false); ok {
			a = access.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width)
		} else {
	*/

	a := access.MakeSingle(addr, 0, st.Width)
	addr += a.RegCount()

	if a.EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: a.EndAddr(),
			EndAddr:   a.EndAddr(),
			StartBit:  a.EndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.Access = a

	return addr
}

func regStaticArray(st *fn.Static, addr int64, gp *gap.Pool) int64 {
	var a access.Access

	if busWidth/2 < st.Width && st.Width <= busWidth {
		a = access.MakeArraySingle(st.Count, addr, 0, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width == 0 || st.Count <= busWidth/st.Width || st.Width < busWidth/2 {
		a = access.MakeArrayMultiplePacked(st.Count, addr, st.Width)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("unimplemented")
	}
	addr += a.RegCount()

	st.Access = a

	return addr
}
