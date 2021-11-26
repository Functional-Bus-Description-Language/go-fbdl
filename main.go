package main

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/args"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/pkg/fbdl"

	"github.com/davecgh/go-spew/spew"

	"encoding/json"
	"fmt"
	"log"
	"os"
)

var printDebug bool = false

type fbdlLogger struct{}

func (l fbdlLogger) Write(p []byte) (int, error) {
	print := true

	if string(p)[:5] == "debug" {
		print = printDebug
	}

	if print {
		fmt.Fprintf(os.Stderr, string(p))
	}

	return len(p), nil
}

func main() {
	logger := fbdlLogger{}
	log.SetOutput(logger)
	log.SetFlags(0)

	cmdLineArgs := args.Parse()

	if _, ok := cmdLineArgs["--debug"]; ok {
		printDebug = true
	}
	log.Printf("debug: Dupa")

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

	insBus := ins.Instantiate(packages)

	if path, ok := cmdLineArgs["-i"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		spew.Fdump(f, insBus)
		err = f.Close()
		if err != nil {
			log.Fatalf("dump instantiation results: %v", err)
		}
	}

	regBus := fbdl.Registerify(insBus)

	if path, ok := cmdLineArgs["-r"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}

		byteArray, err := json.MarshalIndent(regBus, "", "\t")
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
}
