// Custom package for command line arguments parsing.
package args

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const Version string = "0.0.0"

func printVersion() {
	fmt.Println(Version)
	os.Exit(0)
}

func Parse() map[string]string {
	args := map[string]string{}

	argsLen := len(os.Args)

	if argsLen == 1 {
		printHelp()
	}

	skipNext := false

	// If there are help flags anywhere, just print help.
	for _, s := range os.Args[1:] {
		if s == "-help" {
			printHelp()
		}
	}

	// If there are version flags anywhere, just print version.
	for _, s := range os.Args[1:] {
		if s == "-version" {
			printVersion()
		}
	}

	osArgs := os.Args[1 : argsLen-1]
	for i, s := range osArgs {
		if skipNext {
			skipNext = false
			continue
		}

		if !strings.HasPrefix(s, "-") {
			log.Fatalf("unexpected argument %s", s)
		}

		switch s {
		case "-p", "-i", "-r", "-c":
			if i < argsLen-3 && !strings.HasPrefix(osArgs[i+1], "-") {
				args[s] = osArgs[i+1]
				skipNext = true
			} else {
				switch s {
				case "-p":
					args[s] = "prs.txt"
				case "-i":
					args[s] = "ins.txt"
				case "-r":
					args[s] = "reg.json"
				case "-c":
					args[s] = "const.json"
				}
			}
		case "-debug":
			args["-debug"] = ""
		case "-zero-timestamp":
			args["-zero-timestamp"] = ""
		default:
			log.Fatalf("invalid option %s", s)
		}
	}

	args["mainFile"] = os.Args[argsLen-1]

	if args["mainFile"][0] == '-' {
		printHelp()
	}

	return args
}

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
  -p [path]  Dump parse results to a file (default path is prs.txt).
  -i [path]  Dump instantiation results to a file (default path is ins.txt).
  -r [path]  Dump registerification results to a file (default path is reg.json).
  -c [path]  Dump packages constants to a file (default path is const.json).
`

func printHelp() {
	fmt.Printf(helpMsg, Version)
	os.Exit(0)
}
