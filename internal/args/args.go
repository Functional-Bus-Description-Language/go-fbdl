// Custom package for command line arguments parsing.
package args

var (
	MainBus  string
	MainFile string

	AddTimestamp bool
	Debug        bool
	NoGaps       bool

	DumpConsts string
)

func isValidFlag(f string) bool {
	flags := map[string]bool{
		"-add-timestamp": true,
		"-debug":         true,
		"-help":          true,
		"-no-gaps":       true,
		"-version":       true,
	}
	if _, ok := flags[f]; ok {
		return true
	}
	return false
}

func isValidParam(p string) bool {
	params := map[string]bool{
		"-c":    true,
		"-main": true,
	}
	if _, ok := params[p]; ok {
		return true
	}
	return false
}
