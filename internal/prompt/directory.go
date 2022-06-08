package prompt

import (
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tx3stn/pkb/internal/dir"
)

// SelectDirectory prompts the user to select a sub driectory in the provided
// parent. If the parent directory does not have any subdirectories this will
// error.
func SelectDirectory(parent string) (string, error) {
	subDirectories, err := dir.GetSubDirectories(parent)
	if err != nil {
		return "", err
	}

	sort.Strings(subDirectories)

	answer := struct {
		Selected string `survey:"directory"`
	}{}

	// TODO: look at making the 'select directory' text configurable.
	err = survey.Ask([]*survey.Question{
		{
			Name: "directory",
			Prompt: &survey.Select{
				Message: "select directory:",
				Options: subDirectories,
			},
		},
	}, &answer)
	if err != nil {
		return "", err
	}

	return answer.Selected, nil
}
