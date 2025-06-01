package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/editor"
	"github.com/tx3stn/pkb/internal/flags"
)

// CreateOpen creates the new command "open" used to open your editor to edit existing notes.
func CreateOpen() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get(flags.ConfigFile, flags.Vault)
			if err != nil {
				return fmt.Errorf("error getting config: %w", err)
			}

			if err := editor.Open(conf.Editor, conf.Directory); err != nil {
				return fmt.Errorf("error opening file in editor: %w", err)
			}

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "open",
	}

	return cmd
}
