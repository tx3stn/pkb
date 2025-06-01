// Package config contains logic related to user config files.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

// Get returns the config read from the file.
func Get(fileFlag string, vaultFlag string) (Config, error) {
	var file string

	var err error

	if fileFlag == "" {
		// TODO: add support for multiple files in same dir
		file, err = FindConfigFile(vaultFlag)
		if err != nil {
			return Config{}, err
		}
	} else {
		_, err := os.Stat(fileFlag)
		if os.IsNotExist(err) {
			return Config{}, fmt.Errorf("%w: %s", ErrConfigNotFound, fileFlag)
		}

		if err != nil {
			return Config{}, fmt.Errorf("error checking for existence of config file: %w", err)
		}

		file = fileFlag
	}

	if file == "" {
		return Config{}, ErrConfigNotFound
	}

	content, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrReadingConfigFile, err)
	}

	var conf Config
	if err = json.Unmarshal(content, &conf); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrUnmashallingJSON, err)
	}

	return conf, nil
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

// FindConfigFile checks the expected paths for a pkb config file and returns the
// path to it if found.
// The paths are checked in the order of precedence:
//   - XDG_CONFIG_DIR
//   - HOME/.config
func FindConfigFile(vaultFlag string) (string, error) {
	paths := []string{}
	dirName := "pkb"
	configFileName := "pkb.json"

	if vaultFlag != "" {
		configFileName = vaultFlag + ".json"
	}

	if xdg, ok := os.LookupEnv("XDG_CONFIG_DIR"); ok {
		paths = append(paths, filepath.Join(xdg, dirName))
	}

	if home, ok := os.LookupEnv("HOME"); ok {
		paths = append(paths, filepath.Join(home, ".config", dirName))
	}

	if len(paths) == 0 {
		return "", nil
	}

	for _, path := range paths {
		file := filepath.Join(path, configFileName)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			// no config file at location, continue looking.
			continue
		}

		return file, nil
	}

	return "", nil
}
