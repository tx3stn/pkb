// Package cmd contains the different CLI commands for interactions in pkb.
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(CreateNew())
	rootCmd.AddCommand(CreateEdit())
	rootCmd.AddCommand(CreateCopy())
	rootCmd.PersistentFlags().
		StringVar(&flags.ConfigFile, "config", "", "config file if not held at default location")

	if err := viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		log.Printf("error binding --config flag: %s", err)
		os.Exit(1)
	}
}

func initConfig() {
	if flags.ConfigFile == "" {
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.AddConfigPath("$XDG_CONFIG_DIR/pkb/")
		viper.AddConfigPath("$HOME/.config/pkb/")
	} else {
		viper.SetConfigFile(flags.ConfigFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error trying to read config file: %s", err)
		os.Exit(1)
	}
}
