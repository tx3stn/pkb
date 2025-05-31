//nolint:dupl
package prompt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/prompt"
)

func TestDirectorySelect(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		parentDir     string
		selectFunc    prompt.DirectorySelectorFunc
		expected      string
		expectedError error
	}{
		"returns selected options": {
			parentDir: "./testdata",
			selectFunc: func(input []string, accessible bool) (string, error) {
				return "bar", nil
			},
			expected:      "bar",
			expectedError: nil,
		},
		"returns error when the directory doesn't exist": {
			parentDir: "./error",
			selectFunc: func(input []string, accessible bool) (string, error) {
				return "", nil
			},
			expected:      "",
			expectedError: prompt.ErrGettingSubDirectories,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			selector := prompt.DirectorySelector{SelectFunc: tc.selectFunc}

			actual, err := selector.Select(tc.parentDir)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
