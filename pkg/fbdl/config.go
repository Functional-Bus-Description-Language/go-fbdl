package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Config represents config element.
type Config struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Groups  []string
	//Range   Range
	Once  bool
	Width int64
}

// HasDecreasingAccessOrder returns true if config must be accessed
// from the end register to the start register order.
// It is useful only in case of some atomic configs.
// If the end register is narrower, then starting writing from the end register
// saves some flip flops, becase the atomic shadow regsiter can be narrower.
func (c *Config) HasDecreasingAccessOrder() bool {
	if !c.Atomic {
		return false
	}

	if asc, ok := c.Access.(AccessSingleContinuous); ok {
		if !asc.IsEndMaskWider() {
			return true
		}
	}

	return false
}

// regConfig registerifies a Config element.
func regConfig(insCfg *ins.Element, addr int64, gp *gapPool) (*Config, int64) {
	cfg := Config{
		Name:    insCfg.Name,
		Doc:     insCfg.Doc,
		IsArray: insCfg.IsArray,
		Count:   insCfg.Count,
		Atomic:  bool(insCfg.Props["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insCfg.Props["width"].(val.Int)),
	}

	if dflt, ok := insCfg.Props["default"].(val.BitStr); ok {
		cfg.Default = MakeBitStr(dflt)
	}

	if groups, ok := insCfg.Props["groups"].(val.List); ok {
		for _, g := range groups {
			cfg.Groups = append(cfg.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insCfg.Props["width"].(val.Int))

	if insCfg.IsArray {
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
		cfg.Access = makeAccessSingle(addr, 0, width)
		if cfg.Access.EndBit() < busWidth-1 {
			gp.Add(gap{
				isArray:   false,
				startAddr: cfg.Access.EndAddr(),
				endAddr:   cfg.Access.EndAddr(),
				mask:      AccessMask{Upper: busWidth - 1, Lower: cfg.Access.EndBit() + 1},
				writeSafe: false,
			})
		}
	}
	addr += cfg.Access.RegCount()

	return &cfg, addr
}
