// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build tools

package tools

import (
	// Ensure that all its dependencies are tracked by go mod, so it can be run with go run.
	_ "github.com/elastic/elastic-package"
	_ "github.com/elastic/package-registry"
)
