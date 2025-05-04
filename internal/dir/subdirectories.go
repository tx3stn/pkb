// Package dir contains logic related to interacting with directories on the
// filesystem during the pkb interactions.
package dir

import (
	"fmt"
	"os"
)

// GetSubDirectories returns a slice of the sub driectories of the provided
// parent path.
func GetSubDirectories(parent string) ([]string, error) {
	allFiles, err := os.ReadDir(parent)
	if err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrReadingDirectory, err)
	}

	subDirectories := []string{}

	for _, directory := range allFiles {
		if directory.IsDir() {
			subDirectories = append(subDirectories, directory.Name())
		}
	}

	if len(subDirectories) == 0 {
		return []string{}, fmt.Errorf("%w: parent '%s'", ErrNoSubDirectories, parent)
	}

	return subDirectories, nil
}
