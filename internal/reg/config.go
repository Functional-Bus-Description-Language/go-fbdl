package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
)

// regConfig registerifies a Config element.
func regConfig(cfg *elem.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray() {
		panic("not yet implemented")
		/* Should it be implemented the same way as for Status?
		if width == busWidth {

		} else if busWidth%width == 0 || insCfg.Count < busWidth/width {
			cfg.Access = makeAccessArrayMultiple(cfg.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("not yet implemented")
		}
		*/
	} else {
		return regConfigSingle(cfg, addr, gp)
	}
}

func regConfigSingle(cfg *elem.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.Width() <= busWidth {
		a := access.MakeSingleSingle(addr, 0, cfg.Width())
		cfg.SetAccess(a)
		if a.Mask().End() < busWidth-1 {
			gp.Add(gap.Gap{
				StartAddr: a.Addr(),
				EndAddr:   a.Addr(),
				StartBit:  a.Mask().End() + 1,
				EndBit:    busWidth - 1,
				WriteSafe: false,
			})
		}
	} else {
		a := access.MakeSingleContinuous(addr, 0, cfg.Width())
		cfg.SetAccess(a)
		if a.EndMask().End() < busWidth-1 {
			gp.Add(gap.Gap{
				StartAddr: a.StartAddr(),
				EndAddr:   a.EndAddr(),
				StartBit:  a.EndMask().End() + 1,
				EndBit:    busWidth - 1,
				WriteSafe: false,
			})
		}
	}

	addr += cfg.Access().RegCount()

	return addr
}
