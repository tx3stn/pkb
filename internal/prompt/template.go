package prompt

import (
	"fmt"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tx3stn/pkb/internal/config"
	"golang.org/x/exp/maps"
)

// SelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type SelectorFunc func(config.Templates) (config.Template, error)

// TemplateSelector is a utility struct to enable mocking of calls to the
// survey prompt for easier testability.
type TemplateSelector struct {
	SelectFunc SelectorFunc
}

// NewTemplateSelector creates a new instance of the TemplateSelector struct.
func NewTemplateSelector() TemplateSelector {
	return TemplateSelector{
		SelectFunc: SelectTemplate,
	}
}

// SelectTemplateWithSubTemplates is a recursive function to select template
// with nested sub templates and return them in a slice so they can all be
// referenced and the fully nested path to a document can be worked out.
func (t TemplateSelector) SelectTemplateWithSubTemplates(
	templates config.Templates,
	selectedTemplates []config.Template,
) ([]config.Template, error) {
	// If there is only one sub template use that by default, so the user is not
	// given a prompt with only a single value.
	selected := templates.First()

	// More than one, so prompt the user to pick which one they want.
	if len(templates) > 1 {
		var err error
		selected, err = t.SelectFunc(templates)
		if err != nil {
			return []config.Template{}, err
		}
	}

	selectedTemplates = append(selectedTemplates, selected)

	if !selected.HasSubTemplates() {
		return selectedTemplates, nil
	}

	return t.SelectTemplateWithSubTemplates(selected.SubTemplates, selectedTemplates)
}

// SelectTemplate prompts the user to select a template from the ones defined in config.
func SelectTemplate(templates config.Templates) (config.Template, error) {
	answer := struct {
		Selected string `survey:"template"`
	}{}

	templateList := maps.Keys(templates)
	sort.Strings(templateList)

	err := survey.Ask([]*survey.Question{
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "select template:",
				Options: templateList,
			},
		},
	}, &answer)
	if err != nil {
		return config.Template{}, err
	}

	selected, ok := templates[answer.Selected]
	if !ok {
		return config.Template{}, fmt.Errorf("no template named '%s' exists in config file", answer.Selected)
	}

	return selected, nil
}
