// Package editor contains logic for sending commands to or interacting with
// the editor the user defined in config.
package editor

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// Open opens the provided editor in the specified directory.
func Open(editorCmd string, directory string) error {
	return OpenFile(editorCmd, directory, ".")
}

// OpenFile opens the provided file.
func OpenFile(editorCmd string, directory string, fileName string) error {
	cmd := exec.Command(editorCmd, fileName)
	cmd.Dir = directory
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "error opening %s in %s", fileName, editorCmd)
	}

	return nil
}
