package fbdl

// Config represents status element.
type Config struct {
	Name    string
	IsArray bool
	Count   int64
	Access  Access

	// Properties
	Atomic  bool
	Default BitStr
	Doc     string
	Groups  []string
	Range   [2]int64
	Once    bool
	Width   int64
}
