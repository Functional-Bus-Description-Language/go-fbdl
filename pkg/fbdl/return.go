package fbdl

// Param represents param element.
type Return struct {
	Name    string
	Doc     string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Width int64
}
