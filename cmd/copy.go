package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aymanbagabas/go-osc52"
	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/prompt"
)

// CreateCopy creates the new command "copy" used to select a note to copy
// to your system clipboard.
func CreateCopy() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			dir, err := config.GetDirectory()
			if err != nil {
				return err
			}

			selected, err := prompt.SelectExistingFile(dir)
			if err != nil {
				return err
			}

			content, err := os.ReadFile(filepath.Clean(selected))
			if err != nil {
				return err
			}

			osc52.Copy(string(content))
			fmt.Printf("copied %s contents to clipboard", selected)
			return nil
		},
		Short: "select a note and copy it's content to your system clipboard",
		Use:   "copy",
	}

	return cmd
}
