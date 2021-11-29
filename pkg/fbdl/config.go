package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Config represents status element.
type Config struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Doc     string
	Groups  []string
	Range   [2]int64
	Once    bool
	Width   int64
}

func registerifyConfig(insCfg *ins.Element, addr int64) (*Config, int64) {
	cfg := Config{
		Name:    insCfg.Name,
		IsArray: insCfg.IsArray,
		Count:   insCfg.Count,
		Atomic:  bool(insCfg.Properties["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insCfg.Properties["width"].(val.Int)),
	}

	if groups, ok := insCfg.Properties["groups"].(val.List); ok {
		for _, g := range groups {
			cfg.Groups = append(cfg.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insCfg.Properties["width"].(val.Int))

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
