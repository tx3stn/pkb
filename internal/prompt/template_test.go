package prompt_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/prompt"
)

func TestSelectTemplateWithSubTemplates(t *testing.T) {
	t.Parallel()

	errorPickingTemplate := errors.New("error picking template")

	testCases := map[string]struct {
		selectorFunc  prompt.SelectorFunc
		input         config.Templates
		expected      []config.Template
		expectedError error
	}{
		"returns single template with no sub templates": {
			selectorFunc: func(templates config.Templates) (config.Template, error) {
				return templates["foo"], nil
			},
			input:         config.Templates{"foo": {File: "foo.tpl.md"}},
			expected:      []config.Template{{File: "foo.tpl.md"}},
			expectedError: nil,
		},
		"returns multiple templates when you have nested sub templates": {
			selectorFunc: func(templates config.Templates) (config.Template, error) {
				return templates["bar"], nil
			},
			input: config.Templates{
				"bar": {
					File: "one.tpl.md",
					SubTemplates: map[string]config.Template{
						"bar": {
							File: "two.tpl.md",
							SubTemplates: map[string]config.Template{
								"bar": {
									File: "three.tpl.md",
								},
							},
						},
					},
				},
			},
			expected: []config.Template{
				{
					File:      "one.tpl.md",
					OutputDir: "",
					SubTemplates: config.Templates{
						"bar": {
							File:      "two.tpl.md",
							OutputDir: "",
							SubTemplates: map[string]config.Template{
								"bar": {
									File:      "three.tpl.md",
									OutputDir: "",
								},
							},
						},
					},
				},
				{
					File:      "two.tpl.md",
					OutputDir: "",
					SubTemplates: config.Templates{
						"bar": {
							File:      "three.tpl.md",
							OutputDir: "",
						},
					},
				},
				{
					File:      "three.tpl.md",
					OutputDir: "",
				},
			},
			expectedError: nil,
		},
		"returns error when select errors": {
			selectorFunc: func(templates config.Templates) (config.Template, error) {
				return config.Template{}, errorPickingTemplate
			},
			input: config.Templates{
				"foo": {File: "foo.tpl.md"},
				"bar": {File: "bar.tpl.md"},
			},
			expected:      []config.Template{},
			expectedError: errorPickingTemplate,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			selector := prompt.TemplateSelector{
				SelectFunc: tc.selectorFunc,
			}

			selectedTemplates := []config.Template{}

			actual, err := selector.SelectTemplateWithSubTemplates(tc.input, selectedTemplates)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
