package prompt

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tx3stn/pkb/internal/dir"
)

// DirectorySelectorFunc is the type def for the selector func used in the
// DirectorySelector struct.
type DirectorySelectorFunc func([]string) (string, error)

// DirectorySelector is a utility struct to enable mocking of calls to the selector
// prompt for easier testability.
type DirectorySelector struct {
	SelectFunc DirectorySelectorFunc
}

// NewDirectorySelector creates a new instance of the DirectorySelector.
func NewDirectorySelector() DirectorySelector {
	return DirectorySelector{
		SelectFunc: SelectDirectory,
	}
}

// Select prompts the user to select a directory's sub directory.
func (d DirectorySelector) Select(parent string) (string, error) {
	subDirectories, err := dir.GetSubDirectories(parent)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGettingSubDirectories, err)
	}

	sort.Strings(subDirectories)

	selected, err := d.SelectFunc(subDirectories)
	if err != nil {
		return "", err
	}

	return selected, nil
}

// SelectDirectory prompts the user to select a sub driectory in the provided
// parent. If the parent directory does not have any subdirectories this will
// error.
func SelectDirectory(subDirectories []string) (string, error) {
	answer := struct {
		Selected string `survey:"directory"`
	}{}

	if err := survey.Ask([]*survey.Question{
		{
			Name: "directory",
			Prompt: &survey.Select{
				Message: "select directory:",
				Options: subDirectories,
			},
		},
	}, &answer); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingDirectory, err)
	}

	return answer.Selected, nil
}
