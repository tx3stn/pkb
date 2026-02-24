// Package editor contains logic for sending commands to or interacting with
// the editor the user defined in config.
package editor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// Open opens the provided editor in the specified directory.
func Open(ctx context.Context, editorCmd string, directory string) error {
	return OpenFile(ctx, editorCmd, directory, ".")
}

// OpenFile opens the provided file.
func OpenFile(ctx context.Context, editorCmd string, directory string, fileName string) error {
	// #nosec G204 -- args are intentional git CLI flags/subcommands
	cmd := exec.CommandContext(ctx, editorCmd, fileName)
	cmd.Dir = directory
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error opening %s in %s: %w", fileName, editorCmd, err)
	}

	return nil
}
