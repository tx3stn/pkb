package config

import (
	"golang.org/x/exp/maps"
)

// Templates represents the structure of a collection of templates.
type Templates map[string]Template

// Template represents the config options for a custom template file used
// for template and sub templates.
type Template struct {
	CustomDateFormat string    `json:"custom_date_format,omitempty"`
	File             string    `json:"file"`
	NameFormat       string    `json:"name_format,omitempty"`
	OutputDir        string    `json:"output_dir"`
	SubTemplates     Templates `json:"sub_templates,omitempty"`
}

// First is a convenience method to return the first template from the
// Templates map.
func (ts Templates) First() Template {
	return maps.Values(ts)[0]
}

// GetNumberOfSubTemplates returns the number of sub templates defined for a
// template.
func (t Template) GetNumberOfSubTemplates() int {
	return len(maps.Keys(t.SubTemplates))
}

// HasSubTemplates checks if the Template struct has sub templates.
// SubTemplates are offered at the pointing of using a template to give a
// specific version of a file.
func (t Template) HasSubTemplates() bool {
	switch t.GetNumberOfSubTemplates() {
	case 0:
		return false
	// handle an empty template inside sub templates
	case 1:
		return !maps.Values(t.SubTemplates)[0].isEmpty()
	default:
		return true
	}
}

// isEmpty checks if the config for the template object is empty.
func (t Template) isEmpty() bool {
	return t.File == "" && len(t.SubTemplates) == 0
}
