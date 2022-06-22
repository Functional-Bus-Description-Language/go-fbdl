package fbdl

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl/access"
)

// gap represents gap in occupied registers.
// writeSafe indicates whether the gap is safe to be written.
// In other words, it indicates whether the registers the gap address to contain only status information.
// Adding writable functionality (for example config or mask) to a gap with writeSafe set to false implies RMW operation on write.
// Both to the new added functionality, and to the one already placed in the registers.
// This requires the gap to point to the Access structs, doesn't it?
type gap struct {
	startAddr int64
	endAddr   int64
	mask      access.Mask
	writeSafe bool
}

func (g gap) Width() int64    { return g.mask.Width() }
func (g gap) StartBit() int64 { return g.mask.Lower }

func (g gap) isArray() bool {
	if g.endAddr > g.startAddr {
		return true
	}
	return false
}

type gapPool struct {
	singleGaps []gap
	arrayGaps  []gap
}

func (gp *gapPool) Add(g gap) {
	if g.isArray() {
		if len(gp.arrayGaps) == 0 {
			gp.arrayGaps = append(gp.arrayGaps, g)
			return
		}
		for i, ag := range gp.arrayGaps {
			if g.Width() < ag.Width() {
				gp.arrayGaps = append(gp.arrayGaps[:i+1], gp.arrayGaps[i:]...)
				gp.arrayGaps[i] = g
				return
			}
		}
	} else {
		if len(gp.singleGaps) == 0 {
			gp.singleGaps = append(gp.singleGaps, g)
			return
		}
		for i, sg := range gp.singleGaps {
			if g.Width() < sg.Width() {
				gp.singleGaps = append(gp.singleGaps[:i+1], gp.singleGaps[i:]...)
				gp.singleGaps[i] = g
				return
			}
			gp.singleGaps = append(gp.singleGaps, g)
		}
	}
}

// getSingle returns single gap from the gapPool if gap with desired parameters is found in the pool.
// In such case second return is true.
// Otherwise second return is false.
// writeSafe parameter indicates wheter gap has to be write safe.
// If writeSafe = true, then gap must be writeSafe.
// if writeSafe = false, then gap can be writeSafe, but does not have to.
func (gp *gapPool) getSingle(width int64, writeSafe bool) (gap, bool) {
	for i, sg := range gp.singleGaps {
		if (sg.Width() >= width) && (!writeSafe || writeSafe && sg.writeSafe) {
			gp.singleGaps = append(gp.singleGaps[:i], gp.singleGaps[i+1:]...)

			return sg, true
		}
	}
	return gap{}, false
}

/*
func (gp *gapPool) getArray(width int64, regCount int64) (gap, bool) {

}
*/
