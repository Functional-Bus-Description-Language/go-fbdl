package args

import (
	"fmt"
	"log"
	"os"
)

func Parse() Args {
	param := ""
	val := false
	maybeVal := false

	args := Args{}

	handleFlag := func(f string) {
		switch f {
		case "-debug":
			args.Debug = true
		case "-help":
			printHelp()
		case "-version":
			printVersion()
		case "-zero-timestamp":
			args.ZeroTimestamp = true
		default:
			panic(fmt.Sprintf("unhandled flag '%s', implement me", f))
		}
	}

	for i, arg := range os.Args[1:] {
		if i == len(os.Args)-2 {
			if arg == "-help" {
				printHelp()
			}
			if arg == "-version" {
				printVersion()
			}
			if val {
				log.Fatalf("missing path to main file")
			}
			args.MainFile = arg
			continue
		}

		if val {
			val = false

			switch param {
			case "-main":
				args.MainFile = arg
			default:
				panic(fmt.Sprintf("unhandled param '%s', implement me", param))
			}
		} else if maybeVal {
			maybeVal = false

			if isValidFlag(arg) {
				handleFlag(arg)
			}

			if isValidParam(arg) {
				param = arg
				continue
			}

			switch param {
			case "-p":
				args.DumpPrs = arg
			case "-i":
				args.DumpIns = arg
			case "-r":
				args.DumpReg = arg
			case "-c":
				args.DumpConsts = arg
			}
		} else {
			if isValidFlag(arg) {
				handleFlag(arg)
			} else if isValidParam(arg) {
				param = arg

				maybeVal = true
				// Parameters default values.
				switch param {
				case "-p":
					args.DumpPrs = "prs.txt"
				case "-i":
					args.DumpIns = "ins.txt"
				case "-r":
					args.DumpReg = "reg.json"
				case "-c":
					args.DumpConsts = "const.json"
				default:
					maybeVal = false
					val = true
				}
			} else {
				log.Fatalf("invalid parameter '%s'", arg)
			}
		}
	}

	return args
}
