// Package prompt contains logic for prompts and user interactions with the CLI.
package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// EnterDirectory prompts the user to enter the name of the directory they want
// to save the created template in.
func EnterDirectory() (string, error) {
	return userPrompt("directory name")
}

// EnterFileName prompts the user to enter the name of the file they are going
// to save a template as, and returns a sanitised.
func EnterFileName() (string, error) {
	return userPrompt("file name")
}

func userPrompt(inputType string) (string, error) {
	name := ""
	prompt := &survey.Input{
		Message: fmt.Sprintf("enter %s:", inputType),
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		return "", err
	}

	return name, nil
}
