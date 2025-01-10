package prs

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DiscoverPackages(main string) Packages {
	var pathsToLook []string

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	pathsToLook = append(pathsToLook, cwd)

	fbdpath := os.Getenv("FBDPATH")
	if len(fbdpath) != 0 {
		pathsToLook = append(pathsToLook, strings.Split(fbdpath, string(os.PathListSeparator))...)
	}

	dbgMsg := fmt.Sprintf("debug: looking for packages in following %d directories:\n", len(pathsToLook))
	for _, path := range pathsToLook {
		dbgMsg += fmt.Sprintf("  %s\n", path)
	}
	log.Print(dbgMsg)

	packages := make(Packages)

	for _, path := range pathsToLook {
		pkgs := checkDir(path)
		for name, ps := range pkgs {
			packages[name] = append(packages[name], ps...)
		}
	}

	// Add main file.
	var tmp []*Package
	tmp = append(tmp, &Package{Name: "main", Path: main})
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

func checkDir(dirPath string) Packages {
	packages := make(Packages)

	dirEntires, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	base := filepath.Base(dirPath)
	pkgPath := dirPath
	hasPkgPrefix := strings.HasPrefix(base, "fbd-")
	isPkgDir := false

	for _, de := range dirEntires {
		dePath := filepath.Join(dirPath, de.Name())
		fileInfo, err := os.Lstat(dePath)
		if err != nil {
			panic(err)
		}
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			dePath, err = filepath.EvalSymlinks(dePath)
			if err != nil {
				panic(err)
			}
		}

		fileInfo, err = os.Stat(dePath)
		if err != nil {
			panic(err)
		}

		if fileInfo.IsDir() {
			pkgs := checkDir(dePath)
			for name, ps := range pkgs {
				packages[name] = append(packages[name], ps...)
			}
			continue
		}

		if !hasPkgPrefix || isPkgDir {
			continue
		}

		fileName := de.Name()
		if strings.HasSuffix(fileName, ".fbd") {
			isPkgDir = true
		}
	}

	if isPkgDir {
		pkgName := strings.TrimPrefix(base, "fbd-")
		pkg := Package{Name: pkgName, Path: pkgPath}
		packages[pkgName] = append(packages[pkgName], &pkg)
	}

	return packages
}
