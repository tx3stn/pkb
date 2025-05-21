package dir_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/dir"
)

func TestGetAllFilesInDirectory(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputDir      string
		ignoreDirs    []string
		ignoreFiles   []string
		expectedError error
		expected      []string
	}{
		"ReturnsAllFilesInDirectory": {
			inputDir:      "no-ignores",
			ignoreDirs:    []string{},
			ignoreFiles:   []string{},
			expectedError: nil,
			expected: []string{
				"testdata/no-ignores/sub/dir/one",
				"testdata/no-ignores/two",
				"testdata/no-ignores/three",
			},
		},
		"DoesNotReturnFilesInIgnoredDirectory": {
			inputDir:      "ignores",
			ignoreDirs:    []string{".obsidian"},
			ignoreFiles:   []string{},
			expectedError: nil,
			expected: []string{
				"testdata/ignores/foo",
				"testdata/ignores/bar",
			},
		},
		"DoesNotReturnIgnoredFiles": {
			inputDir:      "ignores",
			ignoreDirs:    []string{".obsidian"},
			ignoreFiles:   []string{"bar"},
			expectedError: nil,
			expected: []string{
				"testdata/ignores/foo",
			},
		},
		"ReturnsEmptySliceForEmptyDirectory": {
			inputDir:      "empty",
			ignoreDirs:    []string{},
			ignoreFiles:   []string{},
			expectedError: nil,
			expected:      []string{},
		},
		"ReturnsErrorForInvalidInputDirectory": {
			inputDir:      "foo",
			ignoreDirs:    []string{},
			ignoreFiles:   []string{},
			expectedError: dir.ErrInvalidDirectoryPath,
			expected:      []string{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := dir.GetAllFilesInDirectory(
				filepath.Join("testdata", tc.inputDir),
				tc.ignoreDirs,
				tc.ignoreFiles,
			)
			require.ErrorIs(t, err, tc.expectedError)
			reflect.DeepEqual(tc.expected, actual)
		})
	}
}