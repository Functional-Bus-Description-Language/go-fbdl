package fbdl

// Param represents param element.
type Param struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	// TODO: Should Default be supported for param?
	// Default BitStr
	Doc   string
	Range Range
	Width int64
}
