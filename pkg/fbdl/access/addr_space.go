package access

type AddrSpace interface {
	Start() int64
	End() int64
}

type AddrSpaceSingle interface {
	AddrSpace
}

type AddrSpaceArray interface {
	AddrSpace

	BlockSize() int64
	Count() int64
}
