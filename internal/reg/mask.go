package reg

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/access"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/elem"
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
		return regMaskSingle(mask, addr)
	}

	return addr
}

func regMaskSingle(mask *elem.Mask, addr int64) int64 {
	if mask.Width() <= busWidth {
		a := access.MakeSingleSingle(addr, 0, mask.Width())
		mask.SetAccess(a)
	} else {
		a := access.MakeSingleContinuous(addr, 0, mask.Width())
		mask.SetAccess(a)
	}

	addr += mask.Access().RegCount()

	return addr
}
