// Package cmd contains the different CLI commands for interactions in pkb.
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tx3stn/pkb/internal/flags"
)

// Version is the CLI version set via linker flags at build time.
//
//nolint:gochecknoglobals
var Version string

//nolint:gochecknoglobals
var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		return nil
	},
	Short:   "manage notes in markdown files",
	Use:     "pkb",
	Version: Version,
}

// Execute executes the root command.
func Execute() error {
	//nolint:wrapcheck
	return rootCmd.Execute()
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(CreateNew())
	rootCmd.AddCommand(CreateEdit())
	rootCmd.AddCommand(CreateCopy())
	rootCmd.AddCommand(CreateOpen())
	rootCmd.PersistentFlags().
		StringVar(&flags.ConfigFile, "config", "", "config file if not held at default location")
}
