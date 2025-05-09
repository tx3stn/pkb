package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aymanbagabas/go-osc52/v2"
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
				return fmt.Errorf("error getting directory from config: %w", err)
			}

			selector := prompt.NewFileSelector()
			selected, err := selector.SelectFromDir(dir)
			if err != nil {
				return fmt.Errorf("error selecting file: %w", err)
			}

			content, err := os.ReadFile(filepath.Clean(selected))
			if err != nil {
				return fmt.Errorf("error reading file: %w", err)
			}

			if _, err := osc52.New(string(content)).WriteTo(os.Stderr); err != nil {
				return fmt.Errorf("error copying to clibboard: %w", err)
			}

			log.Printf("copied %s contents to clipboard", selected)

			return nil
		},
		Short: "select a note and copy it's content to your system clipboard",
		Use:   "copy",
	}

	return cmd
}
