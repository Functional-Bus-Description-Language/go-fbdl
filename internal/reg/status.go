package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	_ "github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
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
	if g, ok := gp.GetSingle(st.Width(), false); ok {
		st.SetAccess(access.MakeSingleSingle(g.EndAddr, g.StartBit(), st.Width()))
	} else {
		st.SetAccess(access.MakeSingle(addr, 0, st.Width()))
		addr += st.Access().RegCount()
	}
	if st.Access().EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: st.Access().EndAddr(),
			EndAddr:   st.Access().EndAddr(),
			Mask:      access.Mask{Upper: busWidth - 1, Lower: st.Access().EndBit() + 1},
			WriteSafe: true,
		})
	}

	return addr
}

func regStatusArray(st *elem.Status, addr int64, gp *gap.Pool) int64 {
	if busWidth/2 < st.Width() && st.Width() <= busWidth {
		st.SetAccess(access.MakeArraySingle(st.Count(), addr, 0, st.Width()))
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%st.Width() == 0 || st.Count() <= busWidth/st.Width() || st.Width() < busWidth/2 {
		st.SetAccess(access.MakeArrayMultiplePacked(st.Count(), addr, st.Width()))
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("not yet implemented")
	}
	addr += st.Access().RegCount()

	return addr
}

func regStatusArraySingle(st *elem.Status, addr, startBit int64) int64 {
	st.SetAccess(access.MakeArraySingle(st.Count(), addr, startBit, st.Width()))

	addr += st.Access().RegCount()

	return addr
}
