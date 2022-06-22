package access

type Access interface {
	RegCount() int64 // RegCount returns the number of occupied registers.
	IsArray() bool
	StartAddr() int64
	EndAddr() int64
	EndBit() int64
	Width() int64 // Width returns total width of single element.
}
