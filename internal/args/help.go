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
  -help     Display help.
  -version  Display version.
  -debug    Print debug messages.
  -add-timestamp  Add bus generation timestamp. The timestamp is not included
                  in the ID calculation. The timestamp is always placed at the
                  end of the bus address space.
  -no-gaps  Do not use gaps during the registerification stage. The flag makes
            the registerification results more predictable in the case of bus
            description changes. On the other hand, the size of the required
            address space is greater.

Parameters:
  -main name  Name of the main bus. Useful for testbenches.
  -c [path]   Dump packages constants to a file (default path is const.json).
`

func printHelp() {
	fmt.Printf(helpMsg, Version)
	os.Exit(0)
}
