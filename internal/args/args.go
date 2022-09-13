package args

type Args struct {
	Debug         bool
	ZeroTimestamp bool

	Main string

	DumpPrs    string
	DumpIns    string
	DumpReg    string
	DumpConsts string

	MainFile string
}

func isValidFlag(f string) bool {
	flags := map[string]bool{
		"-help": true, "-version": true, "-debug": true, "-zero-timestamp": true,
	}
	if _, ok := flags[f]; ok {
		return true
	}
	return false
}

func isValidParam(p string) bool {
	params := map[string]bool{
		"-main": true, "-p": true, "-i": true, "-r": true, "-c": true,
	}
	if _, ok := params[p]; ok {
		return true
	}
	return false
}
