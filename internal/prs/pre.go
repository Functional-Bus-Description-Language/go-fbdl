package prs

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func DiscoverPackages(main string) Packages {
	var pathsToLook []string

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathsToLook = append(pathsToLook, cwd)

	cwdfbd := path.Join(cwd, "fbd")
	_, err = os.Stat(cwdfbd)
	if err == nil {
		pathsToLook = append(pathsToLook, cwdfbd)
	}

	fbdpath := os.Getenv("FBDPATH")
	if len(fbdpath) != 0 {
		for _, path := range strings.Split(fbdpath, ":") {
			pathsToLook = append(pathsToLook, path)
		}
	}

	// TODO: Add support for $HOME/local/.lib/fbd

	dbgMsg := fmt.Sprintf("debug: looking for packages in following %d directories:\n", len(pathsToLook))
	for _, path := range pathsToLook {
		dbgMsg += fmt.Sprintf("  %s\n", path)
	}
	log.Print(dbgMsg)

	packages := make(Packages)

	for _, checkPath := range pathsToLook {
		content, err := os.ReadDir(checkPath)
		if err != nil {
			panic(err)
		}

		for _, c := range content {
			pkgPath := path.Join(checkPath, c.Name())

			if !c.IsDir() {
				fileInfo, err := os.Lstat(pkgPath)
				if err != nil {
					panic(err)
				}
				if fileInfo.Mode()&os.ModeSymlink != 0 {
					pointee, err := os.Readlink(path.Join(checkPath, c.Name()))
					if err != nil {
						panic(err)
					}
					pkgPath = path.Join(checkPath, pointee)
				} else {
					continue
				}
			}

			pkgName := c.Name()
			if strings.HasPrefix(pkgName, "fbd-") {
				pkgName = pkgName[4:]
			}
			innerContent, err := os.ReadDir(path.Join(checkPath, c.Name()))
			if err != nil {
				panic(err)
			}
			for _, ic := range innerContent {
				if ic.IsDir() {
					continue
				}
				fileName := ic.Name()
				if strings.HasSuffix(fileName, ".fbd") {
					pkg := Package{Name: pkgName, Path: pkgPath, Symbols: SymbolContainer{}}
					if list, ok := packages[pkgName]; ok {
						list = append(list, &pkg)
					} else {
						packages[pkgName] = []*Package{&pkg}
					}
					break
				}
			}
		}
	}

	// Add main file.
	var tmp []*Package
	tmp = append(tmp, &Package{Name: "main", Path: main, Symbols: SymbolContainer{}})
	packages["main"] = tmp

	pkgsCount := 0
	for _, pkgs := range packages {
		pkgsCount += len(pkgs)
	}

	dbgMsg = fmt.Sprintf("debug: found following %d packages:\n", pkgsCount)
	for _, pkgs := range packages {
		for _, pkg := range pkgs {
			dbgMsg += fmt.Sprintf("  %s: %s\n", pkg.Name, pkg.Path)
		}
	}
	log.Print(dbgMsg)

	return packages
}
