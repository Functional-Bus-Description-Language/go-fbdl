package gap

// Gap represents gap in occupied registers.
// writeSafe indicates whether the gap is safe to be written.
// In other words, it indicates whether the registers the gap address to contain only status information.
// Adding writable functionality (for example config or mask) to a gap with writeSafe set to false implies RMW operation on write.
// Both to the new added functionality, and to the one already placed in the registers.
// This requires the gap to point to the Access structs, doesn't it?
type Gap struct {
	StartAddr int64
	EndAddr   int64
	StartBit  int64
	EndBit    int64
	WriteSafe bool
}

func (g Gap) Width() int64 { return g.EndBit - g.StartBit + 1 }

func (g Gap) IsArray() bool {
	return g.EndAddr > g.StartAddr
}

type Pool struct {
	singleGaps []Gap
	arrayGaps  []Gap
}

func (p *Pool) Add(g Gap) {
	if g.IsArray() {
		if len(p.arrayGaps) == 0 {
			p.arrayGaps = append(p.arrayGaps, g)
			return
		}
		for i, ag := range p.arrayGaps {
			if g.Width() < ag.Width() {
				p.arrayGaps = append(p.arrayGaps[:i+1], p.arrayGaps[i:]...)
				p.arrayGaps[i] = g
				return
			}
		}
	} else {
		if len(p.singleGaps) == 0 {
			p.singleGaps = append(p.singleGaps, g)
			return
		}
		for i, sg := range p.singleGaps {
			if g.Width() < sg.Width() {
				p.singleGaps = append(p.singleGaps[:i+1], p.singleGaps[i:]...)
				p.singleGaps[i] = g
				return
			}
			p.singleGaps = append(p.singleGaps, g)
		}
	}
}

// getSingle returns single gap from the Pool if gap with desired parameters is found in the pool.
// In such case second return is true.
// Otherwise second return is false.
// writeSafe parameter indicates wheter gap has to be write safe.
// If writeSafe = true, then gap must be writeSafe.
// if writeSafe = false, then gap can be writeSafe, but does not have to.
func (p *Pool) GetSingle(width int64, writeSafe bool) (Gap, bool) {
	for i, sg := range p.singleGaps {
		if (sg.Width() >= width) && (!writeSafe || writeSafe && sg.WriteSafe) {
			p.singleGaps = append(p.singleGaps[:i], p.singleGaps[i+1:]...)

			return sg, true
		}
	}
	return Gap{}, false
}

/*
func (p *Pool) getArray(width int64, regCount int64) (Gap, bool) {

}
*/
