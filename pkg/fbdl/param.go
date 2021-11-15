package fbdl

// Param represents param element.
type Param struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Default BitStr
	Doc     string
	Range   Range
	Width   int64
}
