package main

import (
	"fmt"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/ins"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/prs"
	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/reg"

	"github.com/davecgh/go-spew/spew"
	"github.com/jessevdk/go-flags"

	"log"
	"os"
)

const VERSION string = "0.0.0"

type logWriter struct{}

type Options struct {
	Version               bool   `short:"v" long:"version" description:"Display version."`
	DumpPackages          string `short:"p" description:"Dump packages to a file." optional:"true" optional-value:"pkgs.txt"`
	DumpInstantiation     string `short:"i" description:"Dump instantiation to a file." optional:"true" optional-value:"ins.txt"`
	DumpRegisterification string `short:"r" description:"Dump registerification to a file." optional:"true" optional-value:"reg.txt"`
}

func main() {
	log.SetFlags(0)

	spew.Config.Indent = "\t"
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true
	spew.Config.SortKeys = true

	var opts Options
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	if opts.Version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	packages := prs.DiscoverPackages(args[0])
	prs.ParsePackages(packages)

	if opts.DumpPackages != "" {
		f, err := os.Create(opts.DumpPackages)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, packages)
	}

	insBus := ins.Instantiate(packages)

	if opts.DumpInstantiation != "" {
		f, err := os.Create(opts.DumpInstantiation)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, insBus)
	}

	regBus := reg.Registerify(insBus)

	if opts.DumpRegisterification != "" {
		f, err := os.Create(opts.DumpRegisterification)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, regBus)
	}
}
