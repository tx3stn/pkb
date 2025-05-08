package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tx3stn/pkb/internal/template"
)

func TestSanitiseFileName(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected string
	}{
		"returns expected string for single word": {
			input:    "foo",
			expected: "foo.md",
		},
		"converts spaces in input to hyphens": {
			input:    "foo bar baz ",
			expected: "foo-bar-baz.md",
		},
		"removes special characters that aren't file path safe": {
			input:    "new&&test_who*dis+",
			expected: "new-test_who-dis.md",
		},
		"doesn't add file extension if provided in name": {
			input:    "already done.md",
			expected: "already-done.md",
		},
		"handles slashes in name": {
			input:    "do/this/one",
			expected: "do-this-one.md",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := template.SanitiseFileName(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
