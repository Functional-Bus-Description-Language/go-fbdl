package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// regStatus registerifies a Status element.
func regStatus(st *elem.Status, addr int64, gp *gap.Pool) int64 {
	if st.IsArray() {
		return regStatusArray(st, addr, gp)
	} else {
		return regStatusSingle(st, addr, gp)
	}
}

func regStatusSingle(st *elem.Status, addr int64, gp *gap.Pool) int64 {
	var a access.Access

	if g, ok := gp.GetSingle(st.Width(), false); ok {
		a = access.MakeSingleSingle(g.EndAddr, g.StartBit, st.Width())
	} else {
		a = access.MakeSingle(addr, 0, st.Width())
		addr += a.RegCount()
	}
	if a.EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: a.EndAddr(),
			EndAddr:   a.EndAddr(),
			StartBit:  a.EndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: true,
		})
	}

	st.SetAccess(a)

	return addr
}

func regStatusArray(st *elem.Status, addr int64, gp *gap.Pool) int64 {
	var a access.Access

	if busWidth/2 < st.Width() && st.Width() <= busWidth {
		a = access.MakeArraySingle(st.Count(), addr, 0, st.Width())
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width() == 0 || st.Count() <= busWidth/st.Width() || st.Width() < busWidth/2 {
		a = access.MakeArrayMultiplePacked(st.Count(), addr, st.Width())
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("not yet implemented")
	}
	addr += a.RegCount()

	st.SetAccess(a)

	return addr
}
