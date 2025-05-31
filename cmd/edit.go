package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/editor"
	"github.com/tx3stn/pkb/internal/prompt"
)

// CreateEdit creates the new command "edit" used to edit an existing note in your editor.
func CreateEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return fmt.Errorf("error getting config: %w", err)
			}

			selector := prompt.NewFileSelector(
				conf.IgnoreDirs,
				conf.IgnoreFiles,
				conf.AccessibleMode,
			)
			file, err := selector.SelectFromDir(conf.Directory)
			if err != nil {
				return fmt.Errorf("error selecting file: %w", err)
			}

			absPath, err := filepath.Abs(file)
			if err != nil {
				return fmt.Errorf("error creating absolute path for file: %w", err)
			}

			if err := editor.OpenFile(conf.Editor, conf.Directory, absPath); err != nil {
				return fmt.Errorf("error opening file in editor: %w", err)
			}

			return nil
		},
		Short: "select and edit an existing note",
		Use:   "edit",
	}

	return cmd
}
