package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
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
		SelectFunc: SelectFileFromDirectory,
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

// SelectFileFromDirectory prompts the user to select a file and returns the
// full path of the selected file.
func SelectFileFromDirectory(filesInDir []string) (string, error) {
	answer := struct {
		Selected string `survey:"file"`
	}{}

	if err := survey.Ask([]*survey.Question{
		{
			Name: "file",
			Prompt: &survey.Select{
				Message: "select existing file:",
				Options: filesInDir,
			},
		},
	}, &answer); err != nil {
		return "", fmt.Errorf("error selecting existing file: %w", err)
	}

	return answer.Selected, nil
}
