package addrSpace

type Array struct {
	Start     int64
	Count     int64
	BlockSize int64
}

func (a Array) isAddrSpace() bool { return true }
