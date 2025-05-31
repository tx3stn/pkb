// Package prompt contains logic for prompts and user interactions with the CLI.
package prompt

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
)

// EnterDirectory prompts the user to enter the name of the directory they want
// to save the created template in.
func EnterDirectory(accessible bool) (string, error) {
	return userPrompt("enter directory name:", accessible)
}

// EnterFileName prompts the user to enter the name of the file they are going
// to save a template as, and returns a sanitised.
func EnterFileName(accessible bool) (string, error) {
	return userPrompt("enter file name:", accessible)
}

func userPrompt(promptString string, accessible bool) (string, error) {
	result := ""

	prompt := huh.NewInput().
		Title(promptString).
		Value(&result)

	prompt.WithAccessible(accessible)

	if err := prompt.Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingTemplate, err)
	}

	if !accessible {
		fmt.Println(strings.ReplaceAll(prompt.View(), "\n", ""))
	}

	return result, nil
}
