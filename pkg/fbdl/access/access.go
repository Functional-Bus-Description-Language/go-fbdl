package access

// Access struct represents information required to access a given functionality.
//
// Please note that some information for a given access type is usually redundant.
// However, the redundant information is useful when writing generators, especially dynamic generators.
// The redundant information provides a common interface for different access type, even though
// the access type is a struct, not an interface.
// Having a simple struct with all fields required by different access types is the only
// way to provide a uniform access interface between different programming languages.
//
// The RegWidth is always equal to the bus width.
// However, it is kept as a field of the Access struct to ease writing dynamic generators.
// It prevents passing the bus width as generation functions argument everywhere.
type Access struct {
	Type string

	RegCount int64 // Number of occupied registers.
	RegWidth int64 // Width of bus register, equal to the bus width.

	ItemCount int64 // Number of stored items.
	ItemWidth int64 // Single item width.

	StartAddr int64 // Address of the first register
	EndAddr   int64 // Address of the last register.

	StartBit int64 // Start bit in the first register.
	EndBit   int64 // End bit in the last register.

	StartRegWidth int64 // Width occupied in the first register.
	EndRegWidth   int64 // Width occupied in the last register.
}
