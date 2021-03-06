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

	"github.com/davecgh/go-spew/spew"
)

var printDebug bool = false

type Logger struct{}

func (l Logger) Write(p []byte) (int, error) {
	print := true

	if len(p) > 4 && string(p)[:5] == "debug" {
		print = printDebug
	}

	if print {
		fmt.Fprintf(os.Stderr, string(p))
	}

	return len(p), nil
}

func main() {
	logger := Logger{}
	log.SetOutput(logger)
	log.SetFlags(0)

	cmdLineArgs := args.Parse()

	if _, ok := cmdLineArgs["-debug"]; ok {
		printDebug = true
	}

	spew.Config.Indent = "\t"
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true

	packages := prs.DiscoverPackages(cmdLineArgs["mainFile"])
	prs.ParsePackages(packages)

	if path, ok := cmdLineArgs["-p"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		spew.Fdump(f, packages)
		err = f.Close()
		if err != nil {
			log.Fatalf("dump parse results: %v", err)
		}
	}

	zeroTimestamp := false
	if _, ok := cmdLineArgs["-zero-timestamp"]; ok {
		zeroTimestamp = true
	}
	bus, pkgsConsts := ins.Instantiate(packages, zeroTimestamp)

	if path, ok := cmdLineArgs["-i"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		spew.Fdump(f, bus)
		err = f.Close()
		if err != nil {
			log.Fatalf("dump instantiation results: %v", err)
		}
	}

	reg.Registerify(bus)

	if path, ok := cmdLineArgs["-r"]; ok {
		f, err := os.Create(path)
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

	if path, ok := cmdLineArgs["-c"]; ok {
		f, err := os.Create(path)
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
