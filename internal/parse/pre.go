package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func DiscoverPackages(main string) Packages {
	var paths_to_look []string

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cwdfbd := path.Join(cwd, "fbd")
	_, err = os.Stat(cwdfbd)
	if err == nil {
		paths_to_look = append(paths_to_look, cwdfbd)
	}

	fbdpath := os.Getenv("FBDPATH")
	if len(fbdpath) != 0 {
		for _, path := range strings.Split(fbdpath, ":") {
			paths_to_look = append(paths_to_look, path)
		}
	}

	// TODO: Add support for $HOME/local/.lib/fbd

	fmt.Println(paths_to_look)

	packages := make(Packages)

	for _, check_path := range paths_to_look {
		content, err := ioutil.ReadDir(check_path)
		if err != nil {
			panic(err)
		}

		for _, c := range content {
			if c.IsDir() == false {
				continue
			}

			pkg_name := c.Name()
			if strings.HasPrefix(pkg_name, "fbd-") {
				pkg_name = pkg_name[4:]
			}
			inner_content, err := ioutil.ReadDir(path.Join(check_path, pkg_name))
			if err != nil {
				panic(err)
			}
			for _, ic := range inner_content {
				if ic.IsDir() {
					continue
				}
				file_name := ic.Name()
				if strings.HasSuffix(file_name, ".fbd") {
					pkg := Package{Name: pkg_name, Path: path.Join(check_path, c.Name()), Symbols: make(map[string]Symbol)}
					if list, ok := packages[pkg_name]; ok {
						list = append(list, &pkg)
					} else {
						packages[pkg_name] = []*Package{&pkg}
					}
					break
				}
			}
		}
	}

	// Add main file.
	var tmp []*Package
	tmp = append(tmp, &Package{Name: "main", Path: main, Symbols: make(map[string]Symbol)})
	packages["main"] = tmp

	return packages
}
