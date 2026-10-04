package parser

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Functional-Bus-Description-Language/go-fbdl/internal/util"
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
	visitedDirs := make(map[util.DirID]struct{})

	for _, path := range pathsToLook {
		findPkgsInDir(path, packages, visitedDirs)
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

func findPkgsInDir(dirPath string, pkgs Packages, visitedDirs map[util.DirID]struct{}) {
	dirID, err := util.GetDirID(dirPath)
	if err != nil {
		return
	}

	if _, ok := visitedDirs[dirID]; ok {
		return
	}

	visitedDirs[dirID] = struct{}{}

	dirEntries, err := os.ReadDir(dirPath)
	// The directory may disappear after GetDirID.
	// For example, when multiple tests are run in parallel, various
	// tools might create and delete directories dynamically.
	if os.IsNotExist(err) {
		return
	}
	if err != nil {
		panic(err)
	}

	findPkgsInEntries(dirPath, dirEntries, pkgs, visitedDirs)
}

// NOTE: Entries may disappear after ReadDir.
// In such a case skip them without losing their siblings.
func findPkgsInEntries(
	dirPath string,
	dirEntries []os.DirEntry,
	pkgs Packages,
	visitedDirs map[util.DirID]struct{},
) {
	base := filepath.Base(dirPath)
	pkgPath := dirPath
	hasPkgPrefix := strings.HasPrefix(base, "fbd-")
	isPkgDir := false

	for _, de := range dirEntries {
		dePath := filepath.Join(dirPath, de.Name())
		fileInfo, err := os.Lstat(dePath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			panic(err)
		}
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			dePath, err = filepath.EvalSymlinks(dePath)
			// Ignore unresolvable links, but still inspect the remaining entries.
			if err != nil {
				continue
			}
		}

		fileInfo, err = os.Stat(dePath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			panic(err)
		}

		if fileInfo.IsDir() {
			findPkgsInDir(dePath, pkgs, visitedDirs)
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
		pkgs[pkgName] = append(pkgs[pkgName], &pkg)
	}
}
