// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// +build mage

package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/magefile/mage/sh"
)

var (
	buildDir     = "./build"
	publicDir    = filepath.Join(buildDir, "public")
	packagePaths = []string{"./packages"}
	tarGz        = true
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

	err = sh.Run("go", "get", "-u", "github.com/elastic/package-registry/dev/generator")
	if err != nil {
		return err
	}

	for _, p := range packagePaths {
		err := sh.Run("generator", "-sourceDir="+p, "-publicDir="+publicDir, "-tarGz="+strconv.FormatBool(tarGz))
		if err != nil {
			return err
		}
	}
	return nil
}

func Clean() error {
	return os.RemoveAll(buildDir)
}
