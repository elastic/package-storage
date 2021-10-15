// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration

package testing

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/magefile/mage/sh"
	"github.com/stretchr/testify/require"
)

// TestSetup tests if Kibana can be run against the current registry
// and the setup command works as expected.
func TestSetup(t *testing.T) {
	err := os.Chdir("environments")
	require.NoError(t, err)

	// Make sure services are shut down again at the end of the test
	defer func() {
		archiveContainerLogs(t)

		err = sh.Run("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "down", "-v")
		require.NoError(t, err)
	}()
	// Spin up services
	go func() {
		err = sh.Run("docker-compose", "-f", "snapshot.yml", "pull")
		require.NoError(t, err)

		err = sh.Run("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "up", "--force-recreate", "--remove-orphans", "--build")
		require.NoError(t, err)
	}()

	// Check for 5 minutes if service is available
	for i := 0; i < 5*60; i++ {
		output, _ := sh.Output("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "ps")
		if err != nil {
			// Log errors but do not act on it as at first it might not be ready yet
			log.Println(err)
		}
		// 3 services must report healthy
		c := strings.Count(output, "healthy")
		if c == 3 {
			break
		}

		// Wait 1 second between each iteration
		time.Sleep(1 * time.Second)
	}

	// Run setup in fleet against registry to see if no errors are returned
	req, err := http.NewRequest("POST", "http://elastic:changeme@localhost:5601/api/fleet/setup", nil)
	require.NoError(t, err)

	req.Header.Add("kbn-xsrf", "ingest_manager")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, string(body))

	require.Equal(t, 200, resp.StatusCode, string(body))

	packageStrings, err := getPackages(t)
	require.NoError(t, err)

	t.Run("install-packages", func(t *testing.T) {
		// Go through all packages and check if they can be installed
		for _, p := range packageStrings {
			// Get a local copy
			p := p
			t.Run(p, func(t *testing.T) {
				installPackage(t, p)
			})
		}
	})
}

func installPackage(t *testing.T, p string) {
	req, err := http.NewRequest("POST", "http://elastic:changeme@localhost:5601/api/fleet/epm/packages/"+p, nil)
	require.NoError(t, err)

	req.Header.Add("kbn-xsrf", "ingest_manager")
	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err, string(body))

	require.Equal(t, 200, resp.StatusCode, string(body))
	log.Println(p)
}

type Package struct {
	Version string
	Name    string
}

func getPackages(t *testing.T) ([]string, error) {
	// The kibana.version must be in sync with the stack version used in snapshot.yml
	resp, err := http.Get("http://localhost:8080/search?experimental=true&kibana.version=7.13.0")
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	return getPackageStrings(body)
}

func getPackageStrings(body []byte) ([]string, error) {
	var packages = []Package{}
	err := json.Unmarshal(body, &packages)
	if err != nil {
		return nil, err
	}

	var packageStrings []string
	for _, p := range packages {
		packageStrings = append(packageStrings, p.Name+"-"+p.Version)
	}

	return packageStrings, nil
}

func archiveContainerLogs(t require.TestingT) {
	const stackLogsPath = "../../build/elastic-stack-logs"

	err := os.RemoveAll(stackLogsPath) // clean the directory from the previous reports.
	require.NoError(t, err)

	err = os.MkdirAll(stackLogsPath, 0755)
	require.NoError(t, err)

	packageRegistryLogs, _ := sh.Output("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "logs", "--no-color", "--timestamps", "package-registry")
	ioutil.WriteFile(filepath.Join(stackLogsPath, "package-registry.log"), []byte(packageRegistryLogs), 0644)

	elasticsearchLogs, _ := sh.Output("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "logs", "--no-color", "--timestamps", "elasticsearch")
	ioutil.WriteFile(filepath.Join(stackLogsPath, "elasticsearch.log"), []byte(elasticsearchLogs), 0644)

	kibanaLogs, _ := sh.Output("docker-compose", "-f", "snapshot.yml", "-f", "local.yml", "logs", "--no-color", "--timestamps", "kibana")
	ioutil.WriteFile(filepath.Join(stackLogsPath, "kibana.log"), []byte(kibanaLogs), 0644)
}
