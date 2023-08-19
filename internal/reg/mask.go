package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regMask registerifies a Mask element.
func regMask(mask *fn.Mask, addr int64) int64 {
	var a access.Access

	if mask.IsArray {
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
		a = access.MakeSingle(addr, 0, mask.Width)
	}
	addr += a.RegCount()

	mask.Access = a

	return addr
}
