package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
)

// Status represents status element.
type Status struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Doc     string
	Groups  []string
	Once    bool
	Width   int64
}

func registerifyStatus(blk *Block, insSt *ins.Element, addr int64) int64 {
	s := Status{
		Name:    insSt.Name,
		IsArray: insSt.IsArray,
		Count:   insSt.Count,
		Atomic:  bool(insSt.Properties["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insSt.Properties["width"].(val.Int)),
	}

	if groups, ok := insSt.Properties["groups"].(val.List); ok {
		for _, g := range groups {
			s.Groups = append(s.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insSt.Properties["width"].(val.Int))

	if insSt.IsArray {
		if width == busWidth {

		} else if busWidth%width == 0 || insSt.Count < busWidth/width {
			s.Access = makeAccessArrayMultiple(s.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("not yet implemented")
		}
	} else {
		s.Access = makeAccessSingle(addr, 0, width)
	}
	addr += s.Access.Count()

	blk.addStatus(&s)

	return addr
}
