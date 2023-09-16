package gap

type Pool struct {
	singles []Single
	arrays  []Array
}

func (p *Pool) Add(g Gap) {
	switch g := g.(type) {
	case Single:
		if len(p.singles) == 0 {
			p.singles = append(p.singles, g)
			return
		}
		for i, s := range p.singles {
			if g.Width() < s.Width() {
				p.singles = append(p.singles[:i+1], p.singles[i:]...)
				p.singles[i] = g
				return
			}
			p.singles = append(p.singles, g)
		}
	case Array:
		if len(p.arrays) == 0 {
			p.arrays = append(p.arrays, g)
			return
		}
		for i, a := range p.arrays {
			if g.Width() < a.Width() {
				p.arrays = append(p.arrays[:i+1], p.arrays[i:]...)
				p.arrays[i] = g

				return
			}
		}
	}
}

// GetSingle returns Single gap from the Pool if gap with desired parameters is found in the pool.
// In such a case second return is true, otherwise second return is false.
//
// writeSafe parameter indicates wheter gap has to be write safe.
// If writeSafe = true, then gap must be write safe.
// if writeSafe = false, then gap can be write safe, but does not have to.
func (p *Pool) GetSingle(width int64, writeSafe bool) (Single, bool) {
	for i, s := range p.singles {
		if (s.Width() >= width) && (!writeSafe || (writeSafe && s.WriteSafe)) {
			p.singles = append(p.singles[:i], p.singles[i+1:]...)
			return s, true
		}
	}
	return Single{}, false
}
