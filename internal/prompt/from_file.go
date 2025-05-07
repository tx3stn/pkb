package prompt

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

// FromFileSelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type FromFileSelectorFunc func([]string) ([]string, error)

// OptsFromFileSelector is a utility struct to enable mocking of calls to the survey
// prompt for easier testability.
type OptsFromFileSelector struct {
	SelectFunc FromFileSelectorFunc
}

// NewOptsFromFileSelector creates a new instance of the file selector.
func NewOptsFromFileSelector() OptsFromFileSelector {
	return OptsFromFileSelector{
		SelectFunc: SelectFromOpts,
	}
}

// SelectFromOpts prompts the user to pick multiple options from the values provided
// by the template.
func SelectFromOpts(opts []string) ([]string, error) {
	answer := struct {
		Selected []string `survey:"opts"`
	}{}

	if err := survey.Ask([]*survey.Question{
		{
			Name: "opts",
			Prompt: &survey.MultiSelect{
				Message: "selecting from options defined in file:",
				Options: opts,
			},
		},
	}, &answer); err != nil {
		return []string{}, fmt.Errorf("error selecting options from template: %w", err)
	}

	return answer.Selected, nil
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

	selected, err := o.SelectFunc(values)
	if err != nil {
		return []string{}, err
	}

	return selected, nil
}
