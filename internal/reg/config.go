package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/gap"
	_ "github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/elem"
	//fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// regConfig registerifies a Config element.
func regConfig(cfg *elem.Config, addr int64, gp *gap.Pool) int64 {
	/*
		cfg := elem.Config{
			Name:    insCfg.Name,
			Doc:     insCfg.Doc,
			IsArray: insCfg.IsArray,
			Count:   insCfg.Count,
			Atomic:  bool(insCfg.Props["atomic"].(val.Bool)),
			Groups:  []string{},
			Width:   int64(insCfg.Props["width"].(val.Int)),
		}
	*/

	/*
		if dflt, ok := insCfg.Props["default"].(val.BitStr); ok {
			cfg.Default = fbdlVal.MakeBitStr(dflt)
		}

		if groups, ok := insCfg.Props["groups"].(val.List); ok {
			for _, g := range groups {
				cfg.Groups = append(cfg.Groups, string(g.(val.Str)))
			}
		}
	*/

	//width := int64(insCfg.Props["width"].(val.Int))

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
				Mask:      access.Mask{Upper: busWidth - 1, Lower: cfg.Access.EndBit() + 1},
				WriteSafe: false,
			})
		}
	}
	addr += cfg.Access.RegCount()

	return addr
}
