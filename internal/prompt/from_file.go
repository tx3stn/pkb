package prompt

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
)

// FromFileSelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type FromFileSelectorFunc func([]string, string) ([]string, error)

// OptsFromFileSelector is a utility struct to enable mocking of calls to the survey
// prompt for easier testability.
type OptsFromFileSelector struct {
	SelectFunc FromFileSelectorFunc
}

// NewOptsFromFileSelector creates a new instance of the file selector.
func NewOptsFromFileSelector() OptsFromFileSelector {
	return OptsFromFileSelector{
		SelectFunc: selectFromOptions,
	}
}

// Select reads the provided JSON file, and provides the options
// as a list for the user to select.
func (o OptsFromFileSelector) Select(jsonPath string) ([]string, error) {
	listFile, err := os.ReadFile(filepath.Clean(jsonPath))
	if err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrReadingOptionsFile, err)
	}

	var values []string
	if err := json.Unmarshal(listFile, &values); err != nil {
		return []string{}, fmt.Errorf("%w: %w", ErrInvalidOptionsFile, err)
	}

	selected, err := o.SelectFunc(
		values,
		strings.TrimSuffix(filepath.Base(jsonPath), filepath.Ext(jsonPath)),
	)
	if err != nil {
		return []string{}, err
	}

	return selected, nil
}

// selectFromOptions prompts the user to select multiple values from the template provided.
func selectFromOptions(opts []string, fileName string) ([]string, error) {
	var selected []string

	prompt := huh.NewMultiSelect[string]().
		Options(huh.NewOptions(opts...)...).
		Title(fmt.Sprintf("select from %s:", fileName)).
		Value(&selected)

	if err := prompt.Run(); err != nil {
		return []string{}, fmt.Errorf("error selecting options from template: %w", err)
	}

	fmt.Println(prompt.View())

	return selected, nil
}
