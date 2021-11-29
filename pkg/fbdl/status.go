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

func registerifyStatus(insSt *ins.Element, addr int64) (*Status, int64) {
	st := Status{
		Name:    insSt.Name,
		IsArray: insSt.IsArray,
		Count:   insSt.Count,
		Atomic:  bool(insSt.Properties["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insSt.Properties["width"].(val.Int)),
	}

	if groups, ok := insSt.Properties["groups"].(val.List); ok {
		for _, g := range groups {
			st.Groups = append(st.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insSt.Properties["width"].(val.Int))

	if insSt.IsArray {
		if width == busWidth {

		} else if busWidth%width == 0 || insSt.Count <= busWidth/width || width < busWidth/2 {
			st.Access = makeAccessArrayMultiplePacked(st.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("not yet implemented")
		}
	} else {
		st.Access = makeAccessSingle(addr, 0, width)
	}
	addr += st.Access.RegCount()

	return &st, addr
}

func registerifyStatusArraySingle(insSt *ins.Element, addr, startBit int64) (*Status, int64) {
	st := Status{
		Name:    insSt.Name,
		IsArray: insSt.IsArray,
		Count:   insSt.Count,
		Atomic:  bool(insSt.Properties["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insSt.Properties["width"].(val.Int)),
	}

	if groups, ok := insSt.Properties["groups"].(val.List); ok {
		for _, g := range groups {
			st.Groups = append(st.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insSt.Properties["width"].(val.Int))

	st.Access = makeAccessArraySingle(insSt.Count, addr, startBit, width)

	addr += st.Access.RegCount()

	return &st, addr
}
