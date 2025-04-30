package dir

import (
	"fmt"
	"os"
	"path/filepath"
)

const filePermissions = 0o750

// CreateParentDirectories creates the parent directories for the rendered file
// if they don't already exist.
func CreateParentDirectories(outputPath string) error {
	parentDir := filepath.Dir(outputPath)

	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, filePermissions); err != nil {
			return fmt.Errorf(
				"%w parent:%s file:%s",
				ErrCreatingParentDirectories,
				parentDir,
				outputPath,
			)
		}
	}

	return nil
}
