package reg

import (
	//"github.com/Functional-Bus-Description-Language/go-fbdl/internal/val"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	//fbdlVal "github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/val"
)

// regMask registerifies a Mask element.
func regMask(mask *elem.Mask, addr int64) int64 {
	/*
		mask := elem.Mask{
			Name:    insMask.Name,
			Doc:     insMask.Doc,
			IsArray: insMask.IsArray,
			Count:   insMask.Count,
			Atomic:  bool(insMask.Props["atomic"].(val.Bool)),
			Groups:  []string{},
			Width:   int64(insMask.Props["width"].(val.Int)),
		}
	*/

	/*
		if dflt, ok := insMask.Props["default"].(val.BitStr); ok {
			mask.Default = fbdlVal.MakeBitStr(dflt)
		}

		if groups, ok := insMask.Props["groups"].(val.List); ok {
			for _, g := range groups {
				mask.Groups = append(mask.Groups, string(g.(val.Str)))
			}
		}
	*/

	//width := int64(insMask.Props["width"].(val.Int))

	if mask.IsArray() {
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
		mask.SetAccess(access.MakeSingle(addr, 0, mask.Width()))
	}
	addr += mask.Access().RegCount()

	return addr
}