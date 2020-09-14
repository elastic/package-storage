// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"github.com/elastic/package-registry/util"
)

var (
	buildDir     = "./build"
	publicDir    = filepath.Join(buildDir, "public")
	packagePaths = []string{"./packages"}
)

func Build() error {
	err := os.RemoveAll(publicDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(publicDir, 0755)
	if err != nil {
		return err
	}

	err = buildPackages()
	if err != nil {
		return err
	}

	err = dryRunPackageRegistry()
	if err != nil {
		return err
	}
	return nil
}

func buildPackages() error {
	packagePaths, err := findPackages()
	if err != nil {
		return err
	}

	for _, packagePath := range packagePaths {
		srcDir := packagePath + "/"
		p, err := util.NewPackage(srcDir)
		if err != nil {
			return err
		}
		dstDir := filepath.Join(publicDir, "package", p.Name, p.Version)

		err = copyPackageFromSource(srcDir, dstDir)
		if err != nil {
			return err
		}
	}
	return nil
}

func findPackages() ([]string, error) {
	var matches []string
	for _, sourceDir := range packagePaths {
		err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			f, err := os.Stat(path)
			if err != nil {
				return err
			}

			if !f.IsDir() {
				return nil // skip as the path is not a directory
			}

			manifestPath := filepath.Join(path, "manifest.yml")
			_, err = os.Stat(manifestPath)
			if os.IsNotExist(err) {
				return nil
			}
			matches = append(matches, path)
			return filepath.SkipDir
		})
		if err != nil {
			return nil, err
		}
	}
	return matches, nil
}

func copyPackageFromSource(src, dst string) error {
	err := os.MkdirAll(dst, 0755)
	if err != nil {
		return err
	}
	err = sh.RunV("rsync", "-a", src, dst)
	if err != nil {
		return err
	}

	return nil
}

func dryRunPackageRegistry() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "reading current directory failed")
	}
	defer os.Chdir(currentDir)

	err = os.Chdir(buildDir)
	if err != nil {
		return errors.Wrapf(err, "can't change directory to %s", buildDir)
	}

	// Creates a basic package-registry config which points to the packages.
	config := `
package_paths:
- "../packages"
`

	err = ioutil.WriteFile("config.yml", []byte(config), 0644)
	if err != nil {
		return err
	}

	err = sh.Run("go", "run", "github.com/elastic/package-registry", "-dry-run=true")
	if err != nil {
		return errors.Wrap(err, "package-registry dry-run failed")
	}
	return nil
}

func Check() error {
	err := Build()
	if err != nil {
		return err
	}

	err = Vendor()
	if err != nil {
		return err
	}

	// Check if no changes are shown
	err = sh.RunV("git", "update-index", "--refresh")
	if err != nil {
		return err
	}
	return sh.RunV("git", "diff-index", "--exit-code", "HEAD", "--")
}

func Clean() error {
	return os.RemoveAll(buildDir)
}

func Vendor() error {
	fmt.Println(">> mod - updating vendor directory")

	err := sh.RunV("go", "mod", "vendor")
	if err != nil {
		return err
	}

	err = sh.RunV("go", "mod", "verify")
	if err != nil {
		return err
	}
	return nil
}

func TestIntegration() error {
	err := Build()
	if err != nil {
		return err
	}
	return sh.RunV("go", "test", "testing/main_integration_test.go", "-v", "-tags=integration", "2>&1", "|", "go-junit-report", ">", "junit-report.xml")
}
