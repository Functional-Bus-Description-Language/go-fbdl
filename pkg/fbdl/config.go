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
	Range   Range
	Once    bool
	Width   int64
}

func registerifyConfig(insCfg *ins.Element, addr int64) (*Config, int64) {
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
	}
	addr += cfg.Access.RegCount()

	return &cfg, addr
}
