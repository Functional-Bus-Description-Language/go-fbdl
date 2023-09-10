package access

type Access interface {
	GetRegCount() int64 // Number of occupied registers.
	StartAddr() int64
	EndAddr() int64
	StartBit() int64
	EndBit() int64
	Width() int64         // Total width of single functionality.
	StartRegWidth() int64 // Width occupied in the first register.
	EndRegWidth() int64   // Width occupied in the last register.
}
