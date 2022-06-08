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
		inputDir    string
		assertError require.ErrorAssertionFunc
		expected    []string
	}{
		"ReturnsAllFilesInDirectory": {
			inputDir:    "no-ignores",
			assertError: require.NoError,
			expected: []string{
				"testdata/no-ignores/sub/dir/one",
				"testdata/no-ignores/two",
				"testdata/no-ignores/three",
			},
		},
		"DoesNotReturnFilesIngnoredDirectory": {
			inputDir:    "ignores",
			assertError: require.NoError,
			expected: []string{
				"testdata/ignores/foo",
				"testdata/ignores/bar",
			},
		},
		"ReturnsEmptySliceForEmptyDirectory": {
			inputDir:    "empty",
			assertError: require.NoError,
			expected:    []string{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := dir.GetAllFilesInDirectory(filepath.Join("testdata", tc.inputDir))
			tc.assertError(t, err)

			reflect.DeepEqual(tc.expected, actual)
		})
	}
}
