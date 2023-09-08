package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regAtomicConfig registerifies an atomic Config functionality.
func regAtomicConfig(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray {
		panic("unimplemented")
		/* Should it be implemented the same way as for Status?
		if width == busWidth {

		} else if busWidth%width == 0 || insCfg.Count < busWidth/width {
			cfg.Access = makeAccessArrayMultiple(cfg.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("not yet implemented")
		}
		*/
	}
	return regAtomicConfigSingle(cfg, addr, gp)
}

func regAtomicConfigSingle(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	acs := access.MakeSingle(addr, 0, cfg.Width)
	if acs.EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.EndAddr(),
			EndAddr:   acs.EndAddr(),
			StartBit:  acs.EndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.RegCount()

	cfg.Access = acs

	return addr
}

func regNonAtomicConfig(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray {
		panic("unimplemented")
	}
	return regNonAtomicConfigSingle(cfg, addr, gp)
}

func regNonAtomicConfigSingle(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	// TODO: Check if there is write-safe gap at the end that can be utilized.
	acs := access.MakeSingle(addr, 0, cfg.Width)
	if acs.EndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.EndAddr(),
			EndAddr:   acs.EndAddr(),
			StartBit:  acs.EndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.RegCount()

	cfg.Access = acs

	return addr
}
