package main

import (
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/args"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"

	"github.com/davecgh/go-spew/spew"

	"log"
	"os"
)

type logWriter struct{}

type Options struct {
	Version               bool   `short:"v" long:"version" description:"Display version."`
	DumpPackages          string `short:"p" description:"Dump packages to a file." optional:"true" optional-value:"pkgs.txt"`
	DumpInstantiation     string `short:"i" description:"Dump instantiation to a file." optional:"true" optional-value:"ins.txt"`
	DumpRegisterification string `short:"r" description:"Dump registerification to a file." optional:"true" optional-value:"reg.txt"`
}

func main() {
	log.SetFlags(0)

	cmdLineArgs := args.Parse()

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
		defer f.Close()
		spew.Fdump(f, packages)
	}

	insBus := ins.Instantiate(packages)

	if path, ok := cmdLineArgs["-i"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, insBus)
	}

	regBus := reg.Registerify(insBus)

	if path, ok := cmdLineArgs["-r"]; ok {
		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, regBus)
	}
}
