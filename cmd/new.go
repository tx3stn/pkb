package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/create"
	"github.com/tx3stn/pkb/internal/editor"
	"github.com/tx3stn/pkb/internal/flags"
	"github.com/tx3stn/pkb/internal/prompt"
)

// CreateNew creates the new command "new" used to create new notes.
func CreateNew() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return fmt.Errorf("error getting config file: %w", err)
			}

			selected := []config.Template{}
			selector := prompt.NewTemplateSelector()

			selected, err = selector.SelectTemplateWithSubTemplates(conf.Templates, selected)
			if err != nil {
				return fmt.Errorf("error selecting template: %w", err)
			}

			renderer := create.NewTemplateRenderer(conf, selected)
			createdFile, err := renderer.CreateAndSaveFile()
			if err != nil {
				return fmt.Errorf("error creating file: %w", err)
			}

			if !flags.NoEdit {
				if err := editor.OpenFile(conf.Editor, conf.Directory, createdFile); err != nil {
					return fmt.Errorf("error opening file in editor: %w", err)
				}
			}

			return nil
		},
		Short: "Create a new note. Select from templates defined in your config file.",
		Use:   "new",
	}

	cmd.Flags().
		BoolVar(&flags.NoEdit, "no-edit", false, "don't open the file in your editor after creating")

	return cmd
}
