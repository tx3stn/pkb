package create_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/create"
)

func TestGetFileName(t *testing.T) {
	t.Parallel()

	testTime, _ := time.Parse(time.RFC3339, "2022-09-19T16:20:00Z")

	testCases := map[string]struct {
		renderer      create.TemplateRenderer
		expected      string
		expectedError error
	}{
		"uses prompt when no value in config": {
			renderer: create.TemplateRenderer{
				NamePrompt: func() (string, error) {
					return "prompted for this string", nil
				},
				Templates: []config.Template{{}},
			},
			expected:      "prompted for this string",
			expectedError: nil,
		},
		"combines values when mutiple provided": {
			renderer: create.TemplateRenderer{
				NamePrompt: func() (string, error) {
					return "wow this is great", nil
				},
				SelectedTemplate: config.Template{
					NameFormat: "{{.Date}}-{{.Prompt}}-{{.Week}}-{{.Year}}-foo",
				},
				Time: testTime,
			},
			expected:      "2022-09-19-wow this is great-38-2022-foo",
			expectedError: nil,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.renderer.GetFileName()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestOutputPath(t *testing.T) {
	t.Parallel()

	testTime, _ := time.Parse(time.RFC3339, "2022-09-19T16:20:00Z")

	testCases := map[string]struct {
		templateRenderer create.TemplateRenderer
		expectedError    error
		expected         string
	}{
		"returns path for single template": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "simple.md",
				Templates: []config.Template{
					{
						File:      "magic.tpl.md",
						OutputDir: "magic",
					},
				},
			},
			expectedError: nil,
			expected:      "/home/username/notes/magic/simple.md",
		},
		"creates full nested dir path when there are subtemplates": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "nested-example.md",
				Templates: []config.Template{
					{
						File:      "foo.tpl.md",
						OutputDir: "foo",
					},
					{
						File:      "bar.tpl.md",
						OutputDir: "bar",
					},
					{
						File:      "wow.tpl.md",
						OutputDir: "wow",
					},
				},
			},
			expectedError: nil,
			expected:      "/home/username/notes/foo/bar/wow/nested-example.md",
		},
		"prompts user for directory input when specified in template config": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "simple.md",
				DirectoryPrompt: func() (string, error) {
					return "foo/dir", nil
				},
				Templates: []config.Template{
					{
						File:      "magic.tpl.md",
						OutputDir: "{{.Prompt}}",
					},
				},
			},
			expectedError: nil,
			expected:      "/home/username/notes/foo/dir/simple.md",
		},
		"prompts user to select directory when specified in config": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "works.md",
				DirectorySelect: func(_ string) (string, error) {
					return "foo-dir", nil
				},
				Templates: []config.Template{
					{
						File:      "works.md",
						OutputDir: "{{.Select}}",
					},
				},
			},
			expectedError: nil,
			expected:      "/home/username/notes/foo-dir/works.md",
		},
		"adds year to path": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "simple.md",
				Templates: []config.Template{
					{
						File:      "magic.tpl.md",
						OutputDir: "magic/{{.Year}}",
					},
				},
				Time: testTime,
			},
			expectedError: nil,
			expected:      "/home/username/notes/magic/2022/simple.md",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.templateRenderer.OutputPath()
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRender(t *testing.T) {
	t.Parallel()

	testTime, _ := time.Parse(time.RFC3339, "2022-09-19T16:20:00Z")

	testCases := map[string]struct {
		renderer        create.TemplateRenderer
		templateContent string
		expected        string
		expectedError   error
	}{
		"expands expected variables": {
			renderer: create.TemplateRenderer{
				Config: config.Config{
					Templates: map[string]config.Template{},
				},
				Name: "example doc",
				SelectedTemplate: config.Template{
					CustomDateFormat: "Monday 2nd January",
				},
				Time: testTime,
			},
			templateContent: "{{.Date}}\n{{.Name}}\n{{.Time}}\n{{.CustomDateFormat}}\n{{.Week}}\n{{.Year}}",
			expected:        "2022-09-19\nexample doc\n16:20\nMonday 19th September\n38\n2022",
			expectedError:   nil,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var actual bytes.Buffer
			err := tc.renderer.Render(tc.templateContent, &actual)
			require.ErrorIs(t, err, tc.expectedError)
			assert.Equal(t, tc.expected, actual.String())
		})
	}
}
