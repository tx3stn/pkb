package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateParentDirectories creates the parent directories for the rendered file
// if they don't already exist.
func CreateParentDirectories(outputPath string) error {
	parentDir := filepath.Dir(outputPath)

	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, 0o750); err != nil {
			return fmt.Errorf("error creating parent directory %s for file %s", parentDir, outputPath)
		}
	}

	return nil
}
