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

	if err := survey.AskOne(prompt, &name); err != nil {
		return "", fmt.Errorf("error prompting for user input: %w", err)
	}

	return name, nil
}
