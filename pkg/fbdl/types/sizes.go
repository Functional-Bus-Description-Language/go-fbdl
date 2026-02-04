package types

// Block address space size.
//
// The actual address space size of a given block equals the value of the Aligned field.
//
// The OwnAligned might be useful, for example, for device tree source generation.
type Sizes struct {
	Own        int64 // Size required by block own data.
	OwnAligned int64 // Aligned own size.
	Cumulated  int64 // Cumulated size including own sizes of all subblocks.
	Aligned    int64 // Aligned size including aligned size of all subblocks.
}
