// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build mage
// +build mage

package main

import (
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"github.com/elastic/package-registry/packages"
)

var (
	buildDir     = "./build"
	testingDir   = "./testing"
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

	fsBuilder := func(p *packages.Package) (packages.PackageFileSystem, error) {
		return packages.NewExtractedPackageFileSystem(p)
	}

	for _, packagePath := range packagePaths {
		srcDir := packagePath + "/"
		p, err := packages.NewPackage(srcDir, fsBuilder)
		if err != nil {
			return errors.Wrapf(err, "new package from %s", srcDir)
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

	// Change to the testing directory to run package-registry from there as it contains the config for it.
	err = os.Chdir(testingDir)
	if err != nil {
		return errors.Wrapf(err, "can't change directory to %s", testingDir)
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

	err = ModTidy()
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

func ModTidy() error {
	return sh.RunV("go", "mod", "tidy")
}

func TestIntegration() error {
	err := Build()
	if err != nil {
		return err
	}
	return sh.RunV("go", "test", "testing/main_integration_test.go", "-v", "-tags=integration", "-test.timeout", "30m", "2>&1", "|", "go-junit-report", ">", "junit-report.xml")
}
