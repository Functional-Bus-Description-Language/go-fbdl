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
		fmt.Fprint(os.Stderr, string(p))
	}

	return len(p), nil
}

func main() {
	logger := Logger{}
	log.SetOutput(logger)
	log.SetFlags(0)

	args := args.Parse()

	printDebug = args.Debug

	spew.Config.Indent = "\t"
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true

	packages := prs.DiscoverPackages(args.MainFile)
	prs.ParsePackages(packages)

	if args.DumpPrs != "" {
		f, err := os.Create(args.DumpPrs)
		if err != nil {
			panic(err)
		}
		spew.Fdump(f, packages)
		err = f.Close()
		if err != nil {
			log.Fatalf("dump parse results: %v", err)
		}
	}

	mainName := "Main"
	if args.Main != "" {
		mainName = args.Main
	}
	bus, pkgsConsts, err := ins.Instantiate(packages, mainName)
	if err != nil {
		log.Fatalf("instantiation: %v", err)
	}

	if args.DumpIns != "" {
		f, err := os.Create(args.DumpIns)
		if err != nil {
			panic(err)
		}
		spew.Fdump(f, bus)
		err = f.Close()
		if err != nil {
			log.Fatalf("dump instantiation results: %v", err)
		}
	}

	if bus != nil {
		reg.Registerify(bus, args.NoTimestamp)
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
