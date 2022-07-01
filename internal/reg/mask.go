package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// regMask registerifies a Mask element.
func regMask(mask *elem.Mask, addr int64) int64 {
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
