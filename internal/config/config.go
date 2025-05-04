// Package config contains logic related to user config files.
package config

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

// CtxKey is the type for the config that gets bound to the cobra context
// so config values can be accessed by cobra commands.
type CtxKey string

// ContextKey is the key value required to access the cobra command context.
const ContextKey CtxKey = "config"

type (
	// Config represents the options defined in the config file.
	Config struct {
		Directory   string    `json:"directory"`
		Editor      string    `json:"editor"`
		TemplateDir string    `json:"template_dir"`
		Templates   Templates `json:"templates"`
	}
)

// Get fetches the config via viper and converts it to a config struct so it
// can be used properly.
func Get() (Config, error) {
	conf := viper.AllSettings()

	jsonContent, err := json.Marshal(conf)
	if err != nil {
		return Config{}, fmt.Errorf("error marshalling config: %w", err)
	}

	parsedConfig := Config{}
	if err := json.Unmarshal(jsonContent, &parsedConfig); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrUnmashallingJSON, err)
	}

	return parsedConfig, nil
}

// GetDirectory returns the directory value defined in config.
func GetDirectory() (string, error) {
	dir := viper.GetString("directory")
	if dir == "" {
		return "", ErrNoDirectory
	}

	return dir, nil
}
