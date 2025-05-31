package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/config"
)

func TestFindConfigFile(t *testing.T) {
	// TODO: look at mocking current directory
	testCases := map[string]struct {
		xdgEnvValue   string
		homeEnvValue  string
		expected      string
		expectedError error
	}{
		"ReturnsXdgFileWhenExists": {
			xdgEnvValue:   "testdata/xdg/valid",
			homeEnvValue:  "testdata/home/",
			expected:      "testdata/xdg/valid/pkb/pkb.json",
			expectedError: nil,
		},
		"ReturnsHomeFileWhenExists": {
			xdgEnvValue:   "",
			homeEnvValue:  "testdata/home/",
			expected:      "testdata/home/.config/pkb/pkb.json",
			expectedError: nil,
		},
		"ReturnsEmptyStringWhenNoEnvVarsAreSet": {
			xdgEnvValue:   "",
			homeEnvValue:  "",
			expected:      "",
			expectedError: nil,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)
			t.Setenv("HOME", tc.homeEnvValue)

			file, err := config.FindConfigFile()
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
		xdgEnvValue   string
		expectedError error
		expected      config.Config
	}{
		"ReturnsErrorWhenFileIsInvalid": {
			fileFlag:      "",
			xdgEnvValue:   "testdata/xdg/invalid",
			expectedError: config.ErrUnmashallingJSON,
			expected:      config.Config{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Setenv("XDG_CONFIG_DIR", tc.xdgEnvValue)

			actual, err := config.Get(tc.fileFlag)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})

		t.Cleanup(func() {
			err := os.Unsetenv("XDG_CONFIG_DIR")
			require.NoError(t, err)
		})
	}
}
