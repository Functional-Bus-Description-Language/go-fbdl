package args

import (
	"fmt"
	"log"
	"os"
)

func Parse() {
	param := ""
	val := false
	maybeVal := false

	handleFlag := func(f string) {
		switch f {
		case "-debug":
			Debug = true
		case "-help":
			printHelp()
		case "-version":
			printVersion()
		case "-add-timestamp":
			AddTimestamp = true
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
			case "-r":
				DumpReg = "reg.json"
			case "-c":
				DumpConsts = "const.json"
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

			MainFile = arg
			continue
		}

		if val {
			val = false

			switch param {
			case "-main":
				MainBus = arg
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
			case "-r":
				DumpReg = arg
			case "-c":
				DumpConsts = arg
			}
		} else {
			if isValidFlag(arg) {
				handleFlag(arg)
			} else {
				handleParam(arg)
			}
		}
	}

	// Arguments post processing

	if MainBus == "" {
		MainBus = "Main"
	}

	if MainFile == "" {
		log.Fatalf("missing path to main file")
	}
}
