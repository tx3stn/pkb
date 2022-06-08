package dir_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/dir"
	"github.com/tx3stn/pkb/internal/test"
)

func TestGetSubDirectories(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		inputDir    string
		assertError require.ErrorAssertionFunc
		expected    []string
	}{
		"ReturnsErrorWhenNoSubDirectoriesExist": {
			inputDir:    "empty",
			assertError: test.IsSentinelError(dir.ErrNoSubDirectories),
			expected:    []string{},
		},
		"ReturnsErrorDirectoryDoesNotExist": {
			inputDir:    "foo",
			assertError: require.Error,
			expected:    []string{},
		},
		"ReturnsSliceOfDirectoriesWhenSubDirectoriesExist": {
			inputDir:    "no-ignores",
			assertError: require.NoError,
			expected:    []string{"sub", "dir"},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := dir.GetSubDirectories(filepath.Join("testdata", tc.inputDir))
			tc.assertError(t, err)

			reflect.DeepEqual(tc.expected, actual)
		})
	}
}
