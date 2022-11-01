package addrSpace

type Single struct {
	Start int64
	End   int64
}

func (s Single) isAddrSpace() bool { return true }
