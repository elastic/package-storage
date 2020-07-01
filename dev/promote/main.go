package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/magefile/mage/sh"
)

const (
	workingDir = "../../build/promote"
)

func main() {
	fmt.Println("Welcome to the promote script")

	err := promote()
	if err != nil {
		log.Fatal(err)
	}
}

// promote moves a package from one branch to an other
// After moving over the package from one branch to an other, it also
// validates if the branches are still valid by running a check and building the container
// If all of this is successful, the new versions are pushed
func promote() error {

	//promote mysql/1.3.2 snapshot staging

	// These are params but only the to can be passed in, the from is directly
	// the previous branch in the release workflow
	p := "snapshot/0.0.1"
	to := "staging"
	from := "snapshot"

	err := os.RemoveAll(workingDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(workingDir, 0755)
	if err != nil {
		return err
	}

	err = os.Chdir(workingDir)
	if err != nil {
		return err
	}

	err = sh.Run("git", "clone", "--single-branch", "--branch", from, "https://github.com/elastic/package-storage.git", from)
	if err != nil {
		return err
	}
	err = sh.Run("git", "clone", "--single-branch", "--branch", to, "https://github.com/elastic/package-storage.git", to)
	if err != nil {
		return err
	}

	oldPath := filepath.Join(from, "packages", p)
	newPath := filepath.Join(to, "packages", p)
	err = os.MkdirAll(newPath, 0755)
	if err != nil {
		return err
	}

	err = sh.Run("mv", oldPath, filepath.Dir(newPath))
	if err != nil {
		return err
	}

	// Check to package
	err = os.Chdir(to)
	if err != nil {
		return err
	}

	err = sh.Run("git", "add", "packages/"+p)
	if err != nil {
		return err
	}
	err = sh.Run("git", "commit", "-a", "-m", "Add package "+p+"")
	if err != nil {
		return err
	}

	err = sh.Run("mage", "build")
	if err != nil {
		return err
	}

	err = sh.Run("docker", "build", ".")
	if err != nil {
		return err
	}

	// Check from package
	err = os.Chdir("../" + from)
	if err != nil {
		return err
	}

	err = sh.Run("git", "commit", "-a", "-m", "Remove package "+p+"")
	if err != nil {
		return err
	}

	err = sh.Run("mage", "build")
	if err != nil {
		return err
	}

	err = sh.Run("docker", "build", ".")
	if err != nil {
		return err
	}

	// Push changes
	// Push to changes
	// Push from changes

	return nil
}
