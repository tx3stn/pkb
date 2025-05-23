// Package template contains logic related to creating files from templates.
package template

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	goTemplate "text/template"
	"time"

	"github.com/tx3stn/pkb/internal/config"
	"github.com/tx3stn/pkb/internal/date"
	"github.com/tx3stn/pkb/internal/dir"
	"github.com/tx3stn/pkb/internal/prompt"
)

// Renderer holds the config required to render and save the template.
type Renderer struct {
	Config           config.Config
	CreatedFilePath  string
	DirectoryPrompt  func() (string, error)
	DirectorySelect  func(string) (string, error)
	Name             string
	NamePrompt       func() (string, error)
	SelectedTemplate config.Template
	Time             time.Time
	Templates        []config.Template
}

// NewRenderer creates a new instance of the TemplateRenderer.
func NewRenderer(conf config.Config, templates []config.Template) Renderer {
	dir := prompt.NewDirectorySelector()

	return Renderer{
		Config:          conf,
		DirectoryPrompt: prompt.EnterDirectory,
		DirectorySelect: dir.Select,
		NamePrompt:      prompt.EnterFileName,
		Time:            time.Now(),
		Templates:       templates,
	}
}

// CreateAndSaveFile creates the required file from the provided template
// and saves it in the correct output directory.
func (t *Renderer) CreateAndSaveFile() (string, error) {
	if err := t.Config.ValidatePaths(); err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}

	t.SelectedTemplate = t.Templates[len(t.Templates)-1]

	outputPath, err := t.OutputPath()
	if err != nil {
		return "", err
	}

	if err := dir.CreateParentDirectories(outputPath); err != nil {
		return "", fmt.Errorf("error creating parent directories: %w", err)
	}

	templateFile := filepath.Clean(
		filepath.Join(t.Config.Directory, t.Config.TemplateDir, t.SelectedTemplate.File),
	)

	contents, err := os.ReadFile(templateFile)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	file, err := os.Create(filepath.Clean(outputPath))
	if err != nil {
		return "", fmt.Errorf("error creating file: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("error closing created file: %s", err)
			os.Exit(1)
		}
	}()

	t.CreatedFilePath = outputPath

	if err := t.Render(string(contents), file, templateFile); err != nil {
		return "", err
	}

	fmt.Printf("file created: %s\n", outputPath)

	return outputPath, nil
}

// GetFileName either prompts the user for input or uses one of the supported
// name specifiers to automatically set the name.
func (t *Renderer) GetFileName() (string, error) {
	if t.SelectedTemplate.NameFormat == "" {
		return t.NamePrompt()
	}

	outputString := t.SelectedTemplate.NameFormat

	if strings.Contains(outputString, "{{.Date}}") {
		outputString = strings.ReplaceAll(outputString, "{{.Date}}", t.Time.Format("2006-01-02"))
	}

	if strings.Contains(outputString, "{{.Prompt}") {
		promptString, err := t.NamePrompt()
		if err != nil {
			return "", err
		}

		outputString = strings.ReplaceAll(outputString, "{{.Prompt}}", promptString)
	}

	year, week := t.Time.ISOWeek()
	if strings.Contains(outputString, "{{.Week}}") {
		outputString = strings.ReplaceAll(outputString, "{{.Week}}", strconv.Itoa(week))
	}

	if strings.Contains(outputString, "{{.Year}}") {
		outputString = strings.ReplaceAll(outputString, "{{.Year}}", strconv.Itoa(year))
	}

	return outputString, nil
}

// Render reads the template content and expands any variables.
func (t *Renderer) Render(content string, writer io.Writer, templatePath string) error {
	now := t.Time
	year, week := now.ISOWeek()

	config := Variables{
		Name:        t.Name,
		Date:        now.Format("2006-01-02"),
		Directory:   filepath.Base(filepath.Dir(t.CreatedFilePath)),
		TemplateDir: filepath.Dir(templatePath),
		Time:        now.Format("15:04"),
		Week:        week,
		Year:        year,
	}

	// If a custom date format is specified on the template config run it through
	// the date utils to better support human friendly output.
	if t.SelectedTemplate.CustomDateFormat != "" {
		config.CustomDateFormat = now.Format(t.SelectedTemplate.CustomDateFormat)

		if date.IncludesSuffixFormat(config.CustomDateFormat) {
			fixedDate, err := date.ReplaceSuffixFormatter(config.CustomDateFormat)
			if err != nil {
				return fmt.Errorf("error adding date suffix: %w", err)
			}

			config.CustomDateFormat = fixedDate
		}
	}

	tpl, err := goTemplate.New("template").
		Funcs(goTemplate.FuncMap{
			"date":  func() string { return config.Date },
			"time":  func() string { return config.Time },
			"title": func() string { return config.Name },
		}).
		Parse(content)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}

	if err := tpl.Execute(writer, config); err != nil {
		return fmt.Errorf("error writing template: %w", err)
	}

	return nil
}

// OutputPath walks the sub template config to get build the full output path
// handling any nested sub templates and prompts or selections for output
// directories.
func (t *Renderer) OutputPath() (string, error) {
	output := []string{t.Config.Directory}

	for _, config := range t.Templates {
		outputDir := config.OutputDir

		var err error
		if config.OutputDir == "{{.Prompt}}" {
			outputDir, err = t.DirectoryPrompt()
			if err != nil {
				return "", err
			}
		}

		if config.OutputDir == "{{.Select}}" {
			outputDir, err = t.DirectorySelect(filepath.Join(output...))
			if err != nil {
				return "", err
			}
		}

		if strings.Contains(config.OutputDir, "{{.Year}}") {
			year, _ := t.Time.ISOWeek()
			outputDir = strings.ReplaceAll(config.OutputDir, "{{.Year}}", strconv.Itoa(year))
		}

		output = append(output, SanitiseDirPath(outputDir))
	}

	if t.Name == "" {
		fileName, err := t.GetFileName()
		if err != nil {
			return "", err
		}

		t.Name = fileName
	}

	output = append(output, SanitiseFileName(t.Name))

	return filepath.Join(output...), nil
}
