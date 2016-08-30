package main

import (
	"os"
	"path/filepath"
)

// initServer will initialize a server that uses 'dir' as the parent directory
// for storing server related files.
func initServer(dir string) error {
	// Create the parent dir if it does not exist.
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return err
	}
	// Create the data dir if it does not exist.
	err = os.MkdirAll(filepath.Join(dir, dataDir), 0700)
	if err != nil {
		return err
	}
}
