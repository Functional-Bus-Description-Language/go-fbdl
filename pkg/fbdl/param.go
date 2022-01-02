package fbdl

// Param represents param element.
type Param struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	// TODO: Should Default be supported for param?
	// Default BitStr
	Range Range
	Width int64
}
