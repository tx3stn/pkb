package prompt

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pkb/internal/dir"
)

// DirectorySelectorFunc is the type def for the selector func used in the
// DirectorySelector struct.
type DirectorySelectorFunc func([]string, bool) (string, error)

// DirectorySelector is a utility struct to enable mocking of calls to the selector
// prompt for easier testability.
type DirectorySelector struct {
	SelectFunc     DirectorySelectorFunc
	accessibleMode bool
}

// NewDirectorySelector creates a new instance of the DirectorySelector.
func NewDirectorySelector(accessible bool) DirectorySelector {
	return DirectorySelector{
		SelectFunc:     selectDirectory,
		accessibleMode: accessible,
	}
}

// Select prompts the user to select a directory's sub directory.
func (d DirectorySelector) Select(parent string) (string, error) {
	subDirectories, err := dir.GetSubDirectories(parent)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrGettingSubDirectories, err)
	}

	sort.Strings(subDirectories)

	selected, err := d.SelectFunc(subDirectories, d.accessibleMode)
	if err != nil {
		return "", err
	}

	return selected, nil
}

// selectDirectory prompts the user to select a sub driectory in the provided
// parent. If the parent directory does not have any subdirectories this will
// error.
func selectDirectory(subDirectories []string, accessible bool) (string, error) {
	var selected string

	prompt := huh.NewSelect[string]().
		Options(huh.NewOptions(subDirectories...)...).
		Title("select directory:").
		Value(&selected)

	if accessible {
		if err := prompt.RunAccessible(os.Stdout, os.Stdin); err != nil {
			return "", fmt.Errorf("%w: %w", ErrSelectingDirectory, err)
		}

		return selected, nil
	}

	if err := prompt.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingDirectory, err)
	}

	fmt.Println(strings.ReplaceAll(prompt.View(), "\n", ""))

	return selected, nil
}
