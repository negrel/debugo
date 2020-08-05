package generator

import (
	"fmt"
	"go/printer"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

// Option represent generator option
type Option func(*Generator) error

func optionErr(err error) Option {
	return func(_ *Generator) error {
		return err
	}
}

func isDir(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		err = fmt.Errorf("%v is not a directory", path)
	}

	return err
}

func getAllPackages(dir string) ([]*packages.Package, error) {
	// Getting packages names
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Recursive call to Packages for subdirectory
	subPkgs := make([]*packages.Package, 0, len(files))
	for _, file := range files {
		// we only care about subpackage
		if !file.IsDir() {
			continue
		}

		pkgs, err := getAllPackages(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		subPkgs = append(subPkgs, pkgs...)
	}

	// Loading packages
	config := &packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax |
			packages.NeedImports | packages.NeedCompiledGoFiles |
			packages.NeedFiles | packages.NeedTypes |
			packages.NeedDeps,

		Tests:      false,
		Dir:        dir,
		BuildFlags: []string{},
	}

	pkgs, err := packages.Load(config)
	if err != nil {
		return nil, err
	}

	return append(pkgs, subPkgs...), nil

}

// SrcDir generate a Generator Option that add all
// the packages (and subpackages) at the given path.
func SrcDir(path string) Option {
	path = filepath.Clean(path)

	pkgs, err := getAllPackages(path)
	if err != nil {
		return optionErr(err)
	}

	return func(gen *Generator) error {
		if gen.packages == nil {
			gen.packages = make([]*packages.Package, 0, len(pkgs))
		}

		// Adding packages
		for _, pkg := range pkgs {
			for i, file := range pkg.Syntax {
				_ = printer.Fprint(os.Stdout, pkg.Fset, file)

				_, file.Name.Name = filepath.Split(pkg.GoFiles[i])
			}
		}
		gen.packages = append(gen.packages, pkgs...)

		return err
	}
}

// OutputDir return a Generator Option that set up
// the output directory of the generator.
func OutputDir(path string) Option {
	err := isDir(path)
	if err != nil {
		return optionErr(err)
	}
	path, err = filepath.Abs(path)
	if err != nil {
		return optionErr(err)
	}

	return func(gen *Generator) error {
		gen.outputDir = path
		return nil
	}
}

// CommonDir return a Generator Option that define
// the common directory where common packages are stored.
func CommonDir(path string) Option {
	path = filepath.Clean(path)

	pkgs, err := getAllPackages(path)
	if err != nil {
		return optionErr(err)
	}

	return func(gen *Generator) error {
		if gen.packages == nil {
			gen.commonPkgs = make([]*packages.Package, 0, len(pkgs))
		}

		gen.commonPkgs = append(gen.commonPkgs, pkgs...)

		return nil
	}
}
