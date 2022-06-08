package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tx3stn/pkb/internal/config"
)

func TestHashSubTemplates(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		template config.Template
		expected bool
	}{
		"returns false when config has empty sub templates": {
			template: config.Template{
				SubTemplates: config.Templates{},
			},
			expected: false,
		},
		"returns false when config has empty template in sub templates": {
			template: config.Template{
				SubTemplates: config.Templates{"foo": {}},
			},
			expected: false,
		},
		"returns false when config has no sub templates": {
			template: config.Template{},
			expected: false,
		},
		"returns true when config has one valid template": {
			template: config.Template{
				SubTemplates: config.Templates{
					"foo": {
						File: "foo.tpl.md",
					},
				},
			},
			expected: true,
		},
		"returns true when config has many valid templates": {
			template: config.Template{
				SubTemplates: config.Templates{
					"foo": {
						File: "foo.tpl.md",
					},
					"bar": {
						File: "bar.tpl.md",
					},
				},
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := tc.template.HasSubTemplates()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
