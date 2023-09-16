package gap

// The gap represents a gap in the occupied registers.
//
// WriteSafe field indicates whether the gap is safe to be written.
// In other words, it indicates whether registers with the particular gap contain only status information.
// Adding writable functionality (for example config or mask) to a gap with WriteSafe set to false implies RMW operation on write.
// Both to the new added functionality, and to the one already placed in the registers.
// This requires the Gap structs to point to the Access structs, doesn't it?
type Gap interface {
	isGap()
	Width() int64
}
