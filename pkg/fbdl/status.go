package fbdl

// Status represents status element.
type Status struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Doc     string
	Groups  []string
	Once    bool
	Width   int64
}
