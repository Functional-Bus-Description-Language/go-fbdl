package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
)

// regConfig registerifies a Config element.
func regConfig(cfg *elem.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray {
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
		cfg.Access = access.MakeSingle(addr, 0, cfg.Width)
		if cfg.Access.EndBit() < busWidth-1 {
			gp.Add(gap.Gap{
				StartAddr: cfg.Access.EndAddr(),
				EndAddr:   cfg.Access.EndAddr(),
				StartBit:  cfg.Access.EndBit() + 1,
				EndBit:    busWidth - 1,
				WriteSafe: false,
			})
		}
	}
	addr += cfg.Access.RegCount()

	return addr
}
