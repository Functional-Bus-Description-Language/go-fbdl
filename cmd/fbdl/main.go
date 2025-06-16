package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/args"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"
)

var printDebug bool = false

type Logger struct{}

func (l Logger) Write(p []byte) (int, error) {
	print := true

	if len(p) > 4 && string(p)[:5] == "debug" {
		print = printDebug
	}

	if print {
		fmt.Fprint(os.Stderr, string(p))
	}

	return len(p), nil
}

func main() {
	logger := Logger{}
	log.SetOutput(logger)
	log.SetFlags(0)

	args.Parse()

	printDebug = args.Debug

	packages := prs.DiscoverPackages(args.MainFile)
	prs.ParsePackages(packages)

	bus, pkgsConsts, err := ins.Instantiate(packages, args.MainBus)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if bus != nil {
		reg.Registerify(bus, args.AddTimestamp)
	}

	if args.DumpReg != "" {
		f, err := os.Create(args.DumpReg)
		if err != nil {
			panic(err)
		}

		byteArray, err := json.MarshalIndent(bus, "", "\t")
		if err != nil {
			log.Fatalf("marshal registerification results: %v", err)
		}

		_, err = f.Write(byteArray)
		if err != nil {
			log.Fatalf("dump registerification results: %v", err)
		}

		err = f.Close()
		if err != nil {
			log.Fatalf("dump registerification results: %v", err)
		}
	}

	if args.DumpConsts != "" {
		f, err := os.Create(args.DumpConsts)
		if err != nil {
			panic(err)
		}

		byteArray, err := json.MarshalIndent(pkgsConsts, "", "\t")
		if err != nil {
			log.Fatalf("marshal packages constants: %v", err)
		}

		_, err = f.Write(byteArray)
		if err != nil {
			log.Fatalf("dump packages constants: %v", err)
		}

		err = f.Close()
		if err != nil {
			log.Fatalf("dump packages constants: %v", err)
		}
	}
}
