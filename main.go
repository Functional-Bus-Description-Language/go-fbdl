package main

import (
	"fmt"
	"github.com/Functional-Bus-Description-Language/go-fbdl/fbdl"
	"github.com/davecgh/go-spew/spew"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

const VERSION string = "0.0.0"

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

type Options struct {
	Version           bool   `short:"v" long:"version" description:"Display version."`
	DumpPackages      string `short:"p" description:"Dump packages to a file." optional:"true" optional-value:"pkgs.txt"`
	DumpInstantiation string `short:"i" description:"Dump instantiation to a file." optional:"true" optional-value:"inst.txt"`
}

//func foo() {
//	fmt.Println(
//		`Functional Bus Description Language compiler front-end.
//Version`, VERSION)
//
//	flag.PrintDefaults()
//}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	spew.Config.Indent = "  "
	spew.Config.DisablePointerAddresses = true
	spew.Config.DisableCapacities = true

	var opts Options
	args, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	//flag.Usage = foo
	//versionFlag := flag.Bool("v", false, "Display version.")
	//dumpPackages := flag.String("p", "pkgs.txt", "Dump packages to a file.")
	//flag.Parse()

	if opts.Version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	packages := fbdl.DiscoverPackages(args[0])
	fbdl.ParsePackages(packages)

	if opts.DumpPackages != "" {
		f, err := os.Create(opts.DumpPackages)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		spew.Fdump(f, packages)
	}
}
