// Package config contains logic related to user config files.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
		AccessibleMode bool      `json:"accessible_mode"`
		Directory      string    `json:"directory"`
		Editor         string    `json:"editor"`
		IgnoreDirs     []string  `json:"ignore_dirs"`
		IgnoreFiles    []string  `json:"ignore_files"`
		TemplateDir    string    `json:"template_dir"`
		Templates      Templates `json:"templates"`
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

// ValidatePaths checks the paths defined in the config file exist, to give
// helpful error messages when they don't.
func (c Config) ValidatePaths() error {
	if _, err := os.Stat(c.Directory); os.IsNotExist(err) {
		return fmt.Errorf("%w '%s'", ErrDirectoryDoesNotExist, c.Directory)
	}

	templatePath := filepath.Join(c.Directory, c.TemplateDir)
	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
		return fmt.Errorf("%w '%s'", ErrTemplateDirectoryDoesNotExist, templatePath)
	}

	return nil
}
