package date_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/date"
)

func TestDaySuffix(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		day      int
		expected string
	}{
		"1st": {
			day:      1,
			expected: "st",
		},
		"3rd": {
			day:      3,
			expected: "rd",
		},
		"11th": {
			day:      11,
			expected: "th",
		},
		"12th": {
			day:      12,
			expected: "th",
		},
		"13th": {
			day:      13,
			expected: "th",
		},
		"21st": {
			day:      21,
			expected: "st",
		},
		"23rd": {
			day:      23,
			expected: "rd",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := date.DaySuffix(tc.day)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIncludesSuffixFormat(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected bool
	}{
		"returns true for valid th": {
			input:    "friday 10th may",
			expected: true,
		},
		"returns true for valid st": {
			input:    "friday 1st may",
			expected: true,
		},
		"returns true for valid nd": {
			input:    "friday 2nd may",
			expected: true,
		},
		"returns true for valid rd": {
			input:    "friday 3rd may",
			expected: true,
		},
		"returns false when no match found": {
			input:    "friday 3 may",
			expected: false,
		},
		"returns false when no match found for similar strings": {
			input:    "3foo th st nd rd",
			expected: false,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := date.IncludesSuffixFormat(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestReplaceSuffixFormatter(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input       string
		expected    string
		assertError require.ErrorAssertionFunc
	}{
		"converts format string with incorrect suffix": {
			input:       "Monday 1th Nov",
			expected:    "Monday 1st Nov",
			assertError: require.NoError,
		},
		"leaves correctly formatted string as is": {
			input:       "Monday 2nd Nov",
			expected:    "Monday 2nd Nov",
			assertError: require.NoError,
		},
		"returns error when day can't be extracted": {
			input:       "Monday Nov",
			expected:    "",
			assertError: require.Error,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := date.ReplaceSuffixFormatter(tc.input)
			tc.assertError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
