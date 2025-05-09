// Package prompt contains logic for prompts and user interactions with the CLI.
package prompt

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

// EnterDirectory prompts the user to enter the name of the directory they want
// to save the created template in.
func EnterDirectory() (string, error) {
	return userPrompt("enter directory name:")
}

// EnterFileName prompts the user to enter the name of the file they are going
// to save a template as, and returns a sanitised.
func EnterFileName() (string, error) {
	return userPrompt("enter file name:")
}

func userPrompt(promptString string) (string, error) {
	result := ""

	prompt := huh.NewInput().
		Title(promptString).
		Value(&result)

	if err := prompt.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingTemplate, err)
	}

	fmt.Println(prompt.View())

	return result, nil
}
