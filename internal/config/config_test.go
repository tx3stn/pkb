package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/config"
)

func TestFindConfigFile(t *testing.T) {
	testCases := map[string]struct {
		xdgEnvValue   string
		homeEnvValue  string
		vaultFlag     string
		expected      string
		expectedError error
	}{
		"ReturnsXdgFileWhenExists": {
			xdgEnvValue:   "testdata/xdg/valid",
			homeEnvValue:  "testdata/home/",
			vaultFlag:     "",
			expected:      "testdata/xdg/valid/pkb/pkb.json",
			expectedError: nil,
		},
		"ReturnsHomeFileWhenExists": {
			xdgEnvValue:   "",
			homeEnvValue:  "testdata/home/",
			vaultFlag:     "",
			expected:      "testdata/home/.config/pkb/pkb.json",
			expectedError: nil,
		},
		"ReturnsEmptyStringWhenNoEnvVarsAreSet": {
			xdgEnvValue:   "",
			homeEnvValue:  "",
			vaultFlag:     "",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)
			t.Setenv("HOME", tc.homeEnvValue)

			file, err := config.FindConfigFile(tc.vaultFlag)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, file)
		})

		t.Cleanup(func() {
			err := os.Unsetenv("XDG_CONFIG_DIR")
			require.NoError(t, err)

			err = os.Unsetenv("HOME")
			require.NoError(t, err)
		})
	}
}

func TestGet(t *testing.T) {
	testCases := map[string]struct {
		fileFlag      string
		vaultFlag     string
		xdgEnvValue   string
		expectedError error
		expected      config.Config
	}{
		"ReturnsErrorWhenFileIsInvalid": {
			fileFlag:      "",
			vaultFlag:     "",
			xdgEnvValue:   "testdata/xdg/invalid",
			expectedError: config.ErrUnmashallingJSON,
			expected:      config.Config{},
		},
		"ReturnsFileSpecifiedByVaultFlag": {
			fileFlag:      "",
			vaultFlag:     "work",
			xdgEnvValue:   "testdata/xdg/valid",
			expectedError: nil,
			expected: config.Config{
				Directory: "/home/username/notes",
				Editor:    "nvim",
				Templates: config.Templates{
					"foo": config.Template{
						File:      "bar.tpl.md",
						OutputDir: "bar",
					},
				},
			},
		},
		"ReturnsErrorIfVaultFlagFileIsNotFound": {
			fileFlag:      "",
			vaultFlag:     "foo",
			xdgEnvValue:   "testdata/xdg/valid",
			expectedError: config.ErrConfigNotFound,
			expected:      config.Config{},
		},
		"ReturnsFileSpecifiedByFileFlagIfValid": {
			fileFlag:      "testdata/xdg/valid/pkb/work.json",
			vaultFlag:     "",
			xdgEnvValue:   "testdata/xdg/valid",
			expectedError: nil,
			expected: config.Config{
				Directory: "/home/username/notes",
				Editor:    "nvim",
				Templates: config.Templates{
					"foo": config.Template{
						File:      "bar.tpl.md",
						OutputDir: "bar",
					},
				},
			},
		},
		"ReturnsErrorIfFileFlagFileIsNotFound": {
			fileFlag:      "testdata/xdg/valid/pkb/foo.json",
			vaultFlag:     "",
			xdgEnvValue:   "testdata/xdg/valid",
			expectedError: config.ErrConfigNotFound,
			expected:      config.Config{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)

			actual, err := config.Get(tc.fileFlag, tc.vaultFlag)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})

		t.Cleanup(func() {
			err := os.Unsetenv("XDG_CONFIG_DIR")
			require.NoError(t, err)
		})
	}
}
