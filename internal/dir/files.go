package dir

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
)

// GetAllFilesInDirectory returns a slice of all of the files in a given directory.
func GetAllFilesInDirectory(dir string) ([]string, error) {
	filePaths := []string{}

	if err := filepath.WalkDir(dir, func(path string, f fs.DirEntry, err error) error {
		if f == nil {
			return fmt.Errorf("%w: %s", ErrInvalidDirectoryPath, dir)
		}

		if f.IsDir() && slices.Contains(ignoreDirectories(), f.Name()) {
			return filepath.SkipDir
		}

		if !f.IsDir() && !slices.Contains(ignoreFiles(), f.Name()) {
			filePaths = append(filePaths, path)
		}

		return nil
	}); err != nil {
		return []string{}, fmt.Errorf("error walking directory: %w", err)
	}

	return filePaths, nil
}

// TODO: support defining these in config file.
func ignoreDirectories() []string {
	return []string{".git", ".obsidian"}
}

func ignoreFiles() []string {
	return []string{".gitignore"}
}
