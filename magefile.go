// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Create a full release with goreleaser. Builds binaries and pushes them to GitHub
func Release() error {
	return sh.Run("goreleaser", "--rm-dist")
}

// Create a pre-release with goreleaser. Builds the binaries but does not push them.
func PreRelease() error {
	return sh.Run("goreleaser", "--rm-dist", "--snapshot")
}

// Build a binary named "microcorn-dev" and place it in /usr/local/bin
func BuildDev() error {
	return sh.Run("go", "build", "-ldflags=-X github.com/kevingimbel/srvc/cmd/cmd.version=local-dev", "-i", "-o", "/usr/local/bin/srvc-dev", "./cmd/")
}
