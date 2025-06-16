package args

var (
	MainBus  string
	MainFile string

	Debug        bool
	AddTimestamp bool

	DumpReg    string
	DumpConsts string
)

func isValidFlag(f string) bool {
	flags := map[string]bool{
		"-help": true, "-version": true, "-debug": true, "-add-timestamp": true,
	}
	if _, ok := flags[f]; ok {
		return true
	}
	return false
}

func isValidParam(p string) bool {
	params := map[string]bool{
		"-main": true, "-r": true, "-c": true,
	}
	if _, ok := params[p]; ok {
		return true
	}
	return false
}
