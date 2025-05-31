// Package flags contains interactions with the global vars used as CLI flags
// for the different commands.
package flags

var (
	// ConfigFile is the variable for the `--config` CLI flag.
	ConfigFile string

	// NoEdit is the variable for the `--no-edit` CLI flag used by the `new`
	// command when you don't want to edit the file after creating.
	NoEdit bool

	// Vault is the variable for the `--vault` CLI flag, used to select a config
	// file from the config directory.
	Vault string
)
