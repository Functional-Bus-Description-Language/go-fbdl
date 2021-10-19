package reg

type AddrSpace interface {
	Start() uint
	End() uint
	IsArray() bool
	Count() uint
}

type AddrSpaceSingle struct {
	start, end uint
}

func (s AddrSpaceSingle) Start() uint { return s.start }

func (s AddrSpaceSingle) End() uint { return s.end }

func (s AddrSpaceSingle) IsArray() bool { return false }

func (s AddrSpaceSingle) Count() uint { return 1 }

type AddrSpaceArray struct {
	start     uint
	count     uint
	BlockSize uint
}

func (a AddrSpaceArray) GetAddress(i uint) (start uint, end uint) {
	start = a.start + i*a.BlockSize
	end = start + a.BlockSize - 1

	return
}

func (a AddrSpaceArray) Start() uint { return a.start }

func (a AddrSpaceArray) End() uint {
	return a.start + a.count*a.BlockSize - 1
}

func (a AddrSpaceArray) IsArray() bool { return true }

func (a AddrSpaceArray) Count() uint { return a.count }
