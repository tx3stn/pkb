package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/editor"
	"github.com/tx3stn/pkb/internal/flags"
	"github.com/tx3stn/pkb/internal/prompt"
)

// CreateEdit creates the new command "edit" used to open your editor to edit existing notes.
func CreateEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return fmt.Errorf("error getting config: %w", err)
			}

			if flags.Pick {
				selector := prompt.NewFileSelector(conf.IgnoreDirs, conf.IgnoreFiles)
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
			}

			if err := editor.Open(conf.Editor, conf.Directory); err != nil {
				return fmt.Errorf("error opening editor: %w", err)
			}

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "edit",
	}

	cmd.Flags().
		BoolVar(&flags.Pick, "pick", false, "select the file you want to open before opening your editor")

	return cmd
}