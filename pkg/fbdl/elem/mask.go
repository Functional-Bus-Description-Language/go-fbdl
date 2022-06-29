package elem

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// Mask represents mask element.
type Mask struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  access.Access

	// Properties
	Atomic  bool
	Default BitStr
	Groups  []string
	Once    bool
	Width   int64
}

// regMask registerifies a Mask element.
func regMask(insMask *ins.Element, addr int64) (*Mask, int64) {
	mask := Mask{
		Name:    insMask.Name,
		Doc:     insMask.Doc,
		IsArray: insMask.IsArray,
		Count:   insMask.Count,
		Atomic:  bool(insMask.Props["atomic"].(val.Bool)),
		Groups:  []string{},
		Width:   int64(insMask.Props["width"].(val.Int)),
	}

	if dflt, ok := insMask.Props["default"].(val.BitStr); ok {
		mask.Default = MakeBitStr(dflt)
	}

	if groups, ok := insMask.Props["groups"].(val.List); ok {
		for _, g := range groups {
			mask.Groups = append(mask.Groups, string(g.(val.Str)))
		}
	}

	width := int64(insMask.Props["width"].(val.Int))

	if insMask.IsArray {
		panic("not yet implemented")
		/* Should it be implemented the same way as for Status?
		if width == busWidth {

		} else if busWidth%width == 0 || insMask.Count < busWidth/width {
			mask.Access = access.MakeArrayMultiple(mask.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("not yet implemented")
		}
		*/
	} else {
		mask.Access = access.MakeSingle(addr, 0, width)
	}
	addr += mask.Access.RegCount()

	return &mask, addr
}
