package fbdl

type Symbol interface {
	Name() string
	LineNumber() uint32
}

type common struct {
	File       *File
	Id         uint32
	lineNumber uint32
	name       string
}

func (c common) Name() string {
	return c.name
}

func (c common) LineNumber() uint32 {
	return c.lineNumber
}

type Constant struct {
	common
	value Expression
}

//type ElementInstantiationType uint8
//
//const (
//	Anonymous ElementInstantiationType = iota
//	Definitive
//)
//
//type ElementType uint8
//
//const (
//	Block ElementType = iota
//	Bus
//	Config
//	Func
//	Status
//)
//
//type Element struct {
//	base
//	Count uint64
//	parent *Symbol
//	Type ElementType
//}
