package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regAtomicConfig registerifies an atomic Config functionality.
func regAtomicConfig(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray {
		return regAtomicConfigArray(cfg, addr, gp)
	}
	return regAtomicConfigSingle(cfg, addr, gp)
}

func regAtomicConfigArray(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	var acs access.Access

	if cfg.Count*cfg.Width <= busWidth {
		acs = access.MakeArrayOneReg(cfg.Count, addr, 0, cfg.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth/2 < cfg.Width && cfg.Width <= busWidth {
		acs = access.MakeArrayOneInReg(cfg.Count, addr, 0, cfg.Width)
		// TODO: This is a place for adding a potential Gap.
	} else if busWidth%cfg.Width == 0 || cfg.Count <= busWidth/cfg.Width || cfg.Width < busWidth/2 {
		acs = access.MakeArrayNInReg(cfg.Count, addr, cfg.Width)
		// TODO: This is a place for adding a potential Gap.
	} else {
		panic("unimplemented")
	}

	addr += acs.GetRegCount()

	cfg.Access = acs

	return addr
}

func regAtomicConfigSingle(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	acs := access.MakeSingle(addr, 0, cfg.Width)
	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.GetEndAddr(),
			EndAddr:   acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.GetRegCount()

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
	if acs.GetEndBit() < busWidth-1 {
		gp.Add(gap.Gap{
			StartAddr: acs.GetEndAddr(),
			EndAddr:   acs.GetEndAddr(),
			StartBit:  acs.GetEndBit() + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.GetRegCount()

	cfg.Access = acs

	return addr
}
