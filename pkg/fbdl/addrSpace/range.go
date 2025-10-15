package addrSpace

type Range struct {
	Start int64
	End   int64
}

func (r Range) Shift(offset int64) Range {
	r.Start += offset
	r.End += offset
	return r
}
