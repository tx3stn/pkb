package prompt

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
)

// SelectFromJSONFile reads the provided JSON file, and provides the options
// as a list for the user to select.
func SelectFromJSONFile(jsonPath string) ([]string, error) {
	listFile, err := os.ReadFile(filepath.Clean(jsonPath))
	if err != nil {
		return []string{}, fmt.Errorf("error reading file: %w", err)
	}

	var values []string
	if err := json.Unmarshal(listFile, &values); err != nil {
		return []string{}, fmt.Errorf("error unmarshalling options file to json: %w", err)
	}

	answer := struct {
		Selected []string `survey:"opts"`
	}{}

	if err = survey.Ask([]*survey.Question{
		{
			Name: "opts",
			Prompt: &survey.MultiSelect{
				Message: "selecting from options defined in file:",
				Options: values,
			},
		},
	}, &answer); err != nil {
		return []string{}, fmt.Errorf("error selecting options from template: %w", err)
	}

	return answer.Selected, nil
}
