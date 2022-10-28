package access

type Access interface {
	RegCount() int64 // Number of occupied registers.
}

type Single interface {
	Access

	Width() int64
}

type SingleSingle interface {
	Single

	Addr() int64 // Returns the same value as StartAddr() and EndAddr()
	Mask() Mask
}

type SingleContinuous interface {
	Single

	StartAddr() int64
	EndAddr() int64

	StartMask() Mask
	EndMask() Mask
}

type Array interface {
	Access

	StartAddr() int64
	EndAddr() int64

	ItemCount() int64
	ItemWidth() int64
}

// ArraySingle describes an access to an array of elements with single element placed within single register.
type ArraySingle interface {
	Array

	Mask() Mask
}

// ArrayContinuous describes an access to an array of elements with single element placed within multiple continuous registers.
type ArrayContinuous interface {
	Array

	StartMask() Mask
	EndMask() Mask
}

// ArrayMultiple describes an access to an array of elements with multiple elements placed within single register.
type ArrayMultiple interface {
	Array

	ItemsPerAccess() int64

	StartMask() Mask
	EndMask() Mask
}
