package fbdl

// Range represents possible value range.
// IsRepresentable indicates whether range can be represented.
// As bounds are of type int64 too wide range cannot be represented.
type Range struct {
	IsRepresentable bool
	Upper, Lower    int64
}
