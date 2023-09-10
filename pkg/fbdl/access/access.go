package access

type Access interface {
	GetRegCount() int64 // Number of occupied registers.
	GetStartAddr() int64
	GetEndAddr() int64
	GetStartBit() int64
	GetEndBit() int64
	GetWidth() int64      // Total width of single functionality.
	StartRegWidth() int64 // Width occupied in the first register.
	EndRegWidth() int64   // Width occupied in the last register.
}
