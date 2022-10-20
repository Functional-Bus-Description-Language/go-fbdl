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
		case "-add-timestamp":
			args.AddTimestamp = true
		default:
			panic(fmt.Sprintf("unhandled flag '%s', implement me", f))
		}
	}

	handleParam := func(p string) {
		if isValidParam(p) {
			param = p

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
			log.Fatalf("invalid parameter '%s'", p)
		}
	}

	for i, arg := range os.Args[1:] {
		if i == len(os.Args)-2 {
			switch arg {
			case "-help":
				printHelp()
			case "-version":
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
				args.Main = arg
			default:
				panic(fmt.Sprintf("unhandled param '%s', implement me", param))
			}
		} else if maybeVal {
			maybeVal = false

			if isValidFlag(arg) {
				handleFlag(arg)
				continue
			} else if isValidParam(arg) {
				handleParam(arg)
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
			} else {
				handleParam(arg)
			}
		}
	}

	return args
}
