package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/tx3stn/pkb/internal/dir"
)

// SelectExistingFile prompt the user to select a file and returns the
// full path of the selected file.
func SelectExistingFile(searchDir string) (string, error) {
	allPaths, err := dir.GetAllFilesInDirectory(searchDir)
	if err != nil {
		return "", err
	}

	answer := struct {
		Selected string `survey:"file"`
	}{}

	err = survey.Ask([]*survey.Question{
		{
			Name: "file",
			Prompt: &survey.Select{
				Message: "select existing note:",
				Options: allPaths,
			},
		},
	}, &answer)
	if err != nil {
		return "", err
	}

	return answer.Selected, nil
}
