package dir

import (
	"io/fs"
	"path/filepath"

	"golang.org/x/exp/slices"
)

// GetAllFilesInDirectory returns a slice of all of the files in a given directory.
func GetAllFilesInDirectory(dir string) ([]string, error) {
	filePaths := []string{}
	err := filepath.WalkDir(dir, func(path string, f fs.DirEntry, err error) error {
		if f.IsDir() && slices.Contains(ignoreDirectories(), f.Name()) {
			return filepath.SkipDir
		}

		if !f.IsDir() && !slices.Contains(ignoreFiles(), f.Name()) {
			filePaths = append(filePaths, path)
		}

		return nil
	})
	if err != nil {
		return []string{}, err
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
