package args

import (
	"fmt"
	"os"
)

var helpMsg string = `Functional Bus Description Language compiler front-end written in Go.
Version: %s

Usage:
  fbdl [flags] [parameters] /path/to/main/fbd/file

Flags:
  -help            Display help.
  -version         Display version.
  -debug           Print debug messages.
  -zero-timestamp  Zero bus timestamp. Useful for regression tests.

Parameters:
  -main name  Name of the main bus. Useful for testbenches.
  -p [path]   Dump parse results to a file (default path is prs.txt).
  -i [path]   Dump instantiation results to a file (default path is ins.txt).
  -r [path]   Dump registerification results to a file (default path is reg.json).
  -c [path]   Dump packages constants to a file (default path is const.json).
`

func printHelp() {
	fmt.Printf(helpMsg, Version)
	os.Exit(0)
}
