// Package flags contains interactions with the global vars used as CLI flags
// for the different commands.
package flags

var (
	// ConfigFile is the variable for the `--config` CLI flag.
	ConfigFile string

	// NoEdit is the variable for the `--no-edit` CLI flag used by the `new`
	// command when you don't want to edit the file after creating.
	NoEdit bool

	// Pick is the variable for the `--pick` CLI flag used by the `edit` command
	// when you want to explicitly pick the file to edit through `pkb`.
	Pick bool
)
