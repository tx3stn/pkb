package prompt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/prompt"
)

func TestSelectFromDir(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputDir      string
		selectFunc    prompt.FileSelectorFunc
		expected      string
		expectedError error
	}{
		"returns selected file from directory": {
			inputDir: "./testdata",
			selectFunc: func(input []string) (string, error) {
				return "example.json", nil
			},
			expected:      "example.json",
			expectedError: nil,
		},
		"returns error when the directory doesn't exist": {
			inputDir: "./testdata/foo",
			selectFunc: func(input []string) (string, error) {
				return "", nil
			},
			expected:      "",
			expectedError: prompt.ErrGettingFilesInDirectory,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			selector := prompt.FileSelector{SelectFunc: tc.selectFunc}

			actual, err := selector.SelectFromDir(tc.inputDir)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
