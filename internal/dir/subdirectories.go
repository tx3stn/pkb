// Package dir contains logic related to interacting with directories on the
// filesystem during the the pkb interactions.
package dir

import (
	"os"

	"github.com/pkg/errors"
)

// GetSubDirectories returns a slice of the sub driectories of the provided
// parent path.
func GetSubDirectories(parent string) ([]string, error) {
	allFiles, err := os.ReadDir(parent)
	if err != nil {
		return []string{}, err
	}

	subDirectories := []string{}

	for _, directory := range allFiles {
		if directory.IsDir() {
			subDirectories = append(subDirectories, directory.Name())
		}
	}

	if len(subDirectories) == 0 {
		return []string{}, errors.Wrapf(ErrNoSubDirectories, "no directories found in %s", parent)
	}

	return []string{}, nil
}
