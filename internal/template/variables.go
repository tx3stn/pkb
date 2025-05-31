package template

import (
	"path/filepath"
	"strings"

	"github.com/tx3stn/pkb/internal/prompt"
)

// Variables are the variables that are expanded when rendering the template.
type Variables struct {
	AccessibleMode   bool
	CustomDateFormat string
	Date             string
	Directory        string
	Name             string
	TemplateDir      string
	Time             string
	Week             int
	Year             int
}

// SelectFromList prompts a user to select the options to add from the specified
// json file.
func (v Variables) SelectFromList(listFileName string) string {
	opts := prompt.NewOptsFromFileSelector(v.AccessibleMode)

	selected, err := opts.Select(filepath.Join(v.TemplateDir, listFileName))
	if err != nil {
		return err.Error()
	}

	return strings.Join(selected, ", ")
}
