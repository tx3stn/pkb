package prompt

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pkb/internal/dir"
)

// FileSelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type FileSelectorFunc func([]string) (string, error)

// FileSelector is a utility struct to enable mocking of calls to the survey
// prompt for easier testability.
type FileSelector struct {
	SelectFunc FileSelectorFunc
}

// NewFileSelector creates a new instance of the file selector.
func NewFileSelector() FileSelector {
	return FileSelector{
		SelectFunc: selectFile,
	}
}

// SelectFromDir prompts the user to select a file from the provided parent directory.
func (f FileSelector) SelectFromDir(searchDir string) (string, error) {
	allPaths, err := dir.GetAllFilesInDirectory(searchDir)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGettingFilesInDirectory, err)
	}

	selected, err := f.SelectFunc(allPaths)
	if err != nil {
		return "", err
	}

	return selected, nil
}

// selectFile prompts the user to select a file and returns the
// full path of the selected file.
func selectFile(filesInDir []string) (string, error) {
	var selected string

	prompt := huh.NewSelect[string]().
		Options(huh.NewOptions(filesInDir...)...).
		Title("select file:").
		Value(&selected)

	if err := prompt.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingFile, err)
	}

	fmt.Println(prompt.View())

	return selected, nil
}
