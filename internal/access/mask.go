package access

type mask struct {
	Start int64
	End   int64
}

type Mask struct {
	mask
}

func (m Mask) Start() int64 { return m.mask.Start }
func (m Mask) End() int64   { return m.mask.End }
func (m Mask) Width() int64 { return m.mask.End - m.mask.Start + 1 }

func makeMask(start, end int64) Mask {
	return Mask{
		mask: mask{
			Start: start, End: end,
		},
	}
}
