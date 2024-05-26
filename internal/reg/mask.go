package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/fn"
)

// regMask registerifies a Mask functionality.
func regMask(mask *fn.Mask, addr *address) {
	var a access.Access

	if mask.IsArray {
		panic("unimplemented")
		/* Should it be implemented the same way as for Status?
		if width == busWidth {

		} else if busWidth%width == 0 || insMask.Count < busWidth/width {
			mask.Access = access.MakeArrayMultiple(mask.Count, addr, width)
			// TODO: This is a place for adding a potential Gap.
		} else {
			panic("unimplemented")
		}
		*/
	} else {
		a = access.MakeSingle(addr.value, 0, mask.Width)
	}
	addr.inc(a.GetRegCount())

	mask.Access = a
}
