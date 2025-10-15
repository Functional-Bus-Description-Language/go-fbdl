package addrSpace

type AddrSpace interface {
	isAddrSpace() bool
}

func Start(as AddrSpace) int64 {
	switch as := as.(type) {
	case Single:
		return as.Start
	case Array:
		return as.Start
	default:
		panic("should never happen")
	}
}

func End(as AddrSpace) int64 {
	switch as := as.(type) {
	case Single:
		return as.End
	case Array:
		return as.Start + as.Count*as.BlockSize - 1
	default:
		panic("should never happen")
	}
}

func Readdress(as AddrSpace, offset int64) AddrSpace {
	switch as := as.(type) {
	case Single:
		as.Start += offset
		as.End += offset
		return as
	case Array:
		as.Start += offset
		return as
	default:
		panic("should never happen")
	}
}
