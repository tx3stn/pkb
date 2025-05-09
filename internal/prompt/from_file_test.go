package prompt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/prompt"
)

func TestSelect(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputFile     string
		selectFunc    prompt.FromFileSelectorFunc
		expected      []string
		expectedError error
	}{
		"returns selected options from file": {
			inputFile: "./testdata/example.json",
			selectFunc: func(input []string, title string) ([]string, error) {
				return []string{"foo", "bar"}, nil
			},
			expected:      []string{"foo", "bar"},
			expectedError: nil,
		},
		"returns error when the file doesn't exist": {
			inputFile: "./testdata/foo.json",
			selectFunc: func(input []string, title string) ([]string, error) {
				return []string{}, nil
			},
			expected:      []string{},
			expectedError: prompt.ErrReadingOptionsFile,
		},
		"returns error when the file isn't valid json": {
			inputFile: "./testdata/invalid.json",
			selectFunc: func(input []string, title string) ([]string, error) {
				return []string{}, nil
			},
			expected:      []string{},
			expectedError: prompt.ErrInvalidOptionsFile,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			selector := prompt.OptsFromFileSelector{SelectFunc: tc.selectFunc}

			actual, err := selector.Select(tc.inputFile)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
