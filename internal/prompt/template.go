package prompt

import (
	"fmt"
	"maps"
	"slices"
	"sort"

	"github.com/charmbracelet/huh"
	"github.com/tx3stn/pkb/internal/config"
)

// TemplateSelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type TemplateSelectorFunc func([]string) (string, error)

// TemplateSelector is a utility struct to enable mocking of calls to the
// survey prompt for easier testability.
type TemplateSelector struct {
	SelectFunc TemplateSelectorFunc
}

// NewTemplateSelector creates a new instance of the TemplateSelector struct.
func NewTemplateSelector() TemplateSelector {
	return TemplateSelector{
		SelectFunc: selectTemplate,
	}
}

// SelectTemplateWithSubTemplates is a recursive function to select template
// with nested sub templates and return them in a slice so they can all be
// referenced and the fully nested path to a document can be worked out.
func (t TemplateSelector) SelectTemplateWithSubTemplates(
	templates config.Templates,
	selectedTemplates []config.Template,
) ([]config.Template, error) {
	var selected config.Template

	// If there is only one sub template use that by default, so the user is not
	// given a prompt with only a single value.
	selected, err := templates.First()
	if err != nil {
		return []config.Template{}, fmt.Errorf("error getting template: %w", err)
	}

	// More than one, so prompt the user to pick which one they want.
	if len(templates) > 1 {
		var err error

		var ok bool

		templateList := slices.AppendSeq(make([]string, 0, len(templates)), maps.Keys(templates))
		sort.Strings(templateList)

		selectedName, err := t.SelectFunc(templateList)
		if err != nil {
			return []config.Template{}, err
		}

		selected, ok = templates[selectedName]
		if !ok {
			return []config.Template{}, fmt.Errorf("%w %s", ErrNoTemplateWithName, selectedName)
		}
	}

	selectedTemplates = append(selectedTemplates, selected)

	if !selected.HasSubTemplates() {
		return selectedTemplates, nil
	}

	return t.SelectTemplateWithSubTemplates(selected.SubTemplates, selectedTemplates)
}

// selectTemplate prompts the user to select a template from the ones defined in config.
func selectTemplate(templates []string) (string, error) {
	var selected string

	if err := huh.NewSelect[string]().
		Options(huh.NewOptions(templates...)...).
		Title("select template:").
		Value(&selected).
		Run(); err != nil {
		return "", fmt.Errorf("%w: %w", ErrSelectingTemplate, err)
	}

	return selected, nil
}
