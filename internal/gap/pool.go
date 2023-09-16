package gap

type Pool struct {
	singles []Gap
	arrays  []Gap
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
