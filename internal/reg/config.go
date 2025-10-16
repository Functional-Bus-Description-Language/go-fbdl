package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/types"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
)

// regAtomicConfig registerifies an atomic Config functionality.
func regAtomicConfig(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	if cfg.IsArray {
		return regAtomicConfigArray(cfg, addr, gp)
	}
	return regAtomicConfigSingle(cfg, addr, gp)
}

func regAtomicConfigArray(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	var acs types.Access

	// TODO: In all below branches a potential gap can be added.
	if cfg.Count*cfg.Width <= busWidth {
		acs = types.MakeArrayOneRegAccess(cfg.Count, addr, 0, cfg.Width)
	} else if busWidth/2 < cfg.Width && cfg.Width <= busWidth {
		acs = types.MakeArrayOneInRegAccess(cfg.Count, addr, 0, cfg.Width)
	} else if cfg.Width <= busWidth/2 && cfg.Count%(busWidth/cfg.Width) == 0 {
		acs = types.MakeArrayNInRegAccess(cfg.Count, addr, cfg.Width)
	} else if cfg.Width <= busWidth/2 {
		acs = types.MakeArrayNInRegMInEndRegAccess(cfg.Count, addr, cfg.Width)
	} else {
		panic("unimplemented")
	}

	addr += acs.RegCount

	cfg.Access = acs

	return addr
}

func regAtomicConfigSingle(cfg *fn.Config, addr int64, gp *gap.Pool) int64 {
	acs := types.MakeSingleAccess(addr, 0, cfg.Width)
	if acs.EndBit < busWidth-1 {
		gp.Add(gap.Single{
			Addr:      acs.EndAddr,
			StartBit:  acs.EndBit + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.RegCount

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
	acs := types.MakeSingleAccess(addr, 0, cfg.Width)
	if acs.EndBit < busWidth-1 {
		gp.Add(gap.Single{
			Addr:      acs.EndAddr,
			StartBit:  acs.EndBit + 1,
			EndBit:    busWidth - 1,
			WriteSafe: false,
		})
	}
	addr += acs.RegCount

	cfg.Access = acs

	return addr
}
