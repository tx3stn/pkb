package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pkb/internal/dir"
)

// FileSelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type FileSelectorFunc func([]string, bool) (string, error)

// FileSelector is a utility struct to enable mocking of calls to the survey
// prompt for easier testability.
type FileSelector struct {
	SelectFunc     FileSelectorFunc
	IgnoreDirs     []string
	IgnoreFiles    []string
	accessibleMode bool
}

// NewFileSelector creates a new instance of the file selector.
func NewFileSelector(ignoreDirs []string, ignoreFiles []string, accessible bool) FileSelector {
	return FileSelector{
		SelectFunc:     selectFile,
		IgnoreDirs:     ignoreDirs,
		IgnoreFiles:    ignoreFiles,
		accessibleMode: accessible,
	}
}

// SelectFromDir prompts the user to select a file from the provided parent directory.
func (f FileSelector) SelectFromDir(searchDir string) (string, error) {
	allPaths, err := dir.GetAllFilesInDirectory(searchDir, f.IgnoreDirs, f.IgnoreFiles)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGettingFilesInDirectory, err)
	}

	selected, err := f.SelectFunc(allPaths, f.accessibleMode)
	if err != nil {
		return "", err
	}

	return selected, nil
}

// selectFile prompts the user to select a file and returns the
// full path of the selected file.
func selectFile(filesInDir []string, accessible bool) (string, error) {
	var selected string

	prompt := huh.NewSelect[string]().
		Options(huh.NewOptions(filesInDir...)...).
		Title("select file:").
		Value(&selected)

	prompt.WithAccessible(accessible)

	if err := prompt.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingFile, err)
	}

	if !accessible {
		fmt.Println(strings.ReplaceAll(prompt.View(), "\n", ""))
	}

	return selected, nil
}
