package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
)

// regStatic registerifies Static element.
func regStatic(st *elem.Static, addr int64, gp *gap.Pool) int64 {
	if st.IsArray() {
		return regStaticArray(st, addr, gp)
	} else {
		return regStaticSingle(st, addr, gp)
	}
}

func regStaticSingle(st *elem.Static, addr int64, gp *gap.Pool) int64 {
	if g, ok := gp.GetSingle(st.Width(), false); ok {
		a := access.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width())
		if a.Mask().End() == busWidth {
			addr += 1
		}
		st.SetAccess(a)
	} else if st.Width() <= busWidth {
		a := access.MakeSingle(addr, 0, st.Width())
		addr += a.RegCount()
		st.SetAccess(a)
	} else {
		a := access.MakeSingleContinuous(addr, 0, st.Width())
		st.SetAccess(a)
	}

	return addr
}

func regStaticArray(st *elem.Static, addr int64, gp *gap.Pool) int64 {
	if busWidth/2 < st.Width() && st.Width() <= busWidth {
		a := access.MakeArraySingle(st.Count(), addr, 0, st.Width())
		addr += a.RegCount()
		st.SetAccess(a)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width() == 0 || st.Count() <= busWidth/st.Width() || st.Width() < busWidth/2 {
		a := access.MakeArrayMultiplePacked(st.Count(), addr, st.Width())
		addr += a.RegCount()
		st.SetAccess(a)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("not yet implemented")
	}

	return addr
}
