package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/mod/modfile"
)

// modFilePath returns the path to the go.mod file for the given source file.
//
// pkg is the file's package path, relative to the module root.
func modFilePath(file string) (mf string, pkg string, err error) {
	dir := file

	for dir != "/" {
		dir = path.Dir(dir)
		mf := path.Join(dir, "go.mod")

		info, err := os.Stat(mf)

		if err != nil {
			if !os.IsNotExist(err) {
				return "", "", err
			}
		} else if !info.IsDir() {
			return mf, pkg, nil
		}

		pkg = path.Join(path.Base(dir), pkg)
	}

	return "", "", fmt.Errorf("could not find go.mod for %s", file)
}

// packagePath returns the absolute package path for the given file.
func packagePath(file string) (string, error) {
	mf, pkg, err := modFilePath(file)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(mf)
	if err != nil {
		return "", err
	}

	mod := modfile.ModulePath(data)
	if mod == "" {
		return "", fmt.Errorf("missing or malformed module path in %s", mod)
	}

	return path.Join(mod, pkg), nil
}
