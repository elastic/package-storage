// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build mage

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

var (
	buildDir     = "./build"
	publicDir    = filepath.Join(buildDir, "public")
	packagePaths = []string{"./packages"}
)

func Build() error {
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
