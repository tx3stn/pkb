package dir_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/dir"
)

func TestGetSubDirectories(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputDir      string
		expectedError error
		expected      []string
	}{
		"ReturnsErrorWhenNoSubDirectoriesExist": {
			inputDir:      "empty",
			expectedError: dir.ErrNoSubDirectories,
			expected:      []string{},
		},
		"ReturnsErrorDirectoryDoesNotExist": {
			inputDir:      "foo",
			expectedError: dir.ErrReadingDirectory,
			expected:      []string{},
		},
		"ReturnsSliceOfDirectoriesWhenSubDirectoriesExist": {
			inputDir:      "no-ignores",
			expectedError: nil,
			expected:      []string{"sub", "dir"},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := dir.GetSubDirectories(filepath.Join("testdata", tc.inputDir))
			require.ErrorIs(t, err, tc.expectedError)
			reflect.DeepEqual(tc.expected, actual)
		})
	}
}
