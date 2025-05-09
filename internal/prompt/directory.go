package prompt

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/huh"
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
		SelectFunc: selectDirectory,
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

// selectDirectory prompts the user to select a sub driectory in the provided
// parent. If the parent directory does not have any subdirectories this will
// error.
func selectDirectory(subDirectories []string) (string, error) {
	huhOpts := make([]huh.Option[string], len(subDirectories))

	for i, v := range subDirectories {
		huhOpts[i] = huh.NewOption(v, v)
	}

	var selected string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Options(huhOpts...).
				Title("select directory:").
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingDirectory, err)
	}

	return selected, nil
}
