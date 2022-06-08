package create

import (
	"fmt"
	"regexp"
	"strings"
)

// SanitiseDirPath sanitises the file path of a directory, swapping out any
// spaces or special characters for hyphens.
func SanitiseDirPath(path string) string {
	separators := regexp.MustCompile(`[ &=+:*]`)
	path = separators.ReplaceAllString(path, "-")

	// Swap any occurrences of doubled up separators from instances where multiple
	// invalid characters have been repalced.
	doubleSeparators := regexp.MustCompile(`--`)
	path = doubleSeparators.ReplaceAllString(path, "-")

	return strings.Trim(path, "-")
}

// SanitiseFileName removes any spaces or special characters so the format
// is valid to use as a file name.
func SanitiseFileName(name string) string {
	baseName := strings.Trim(name, " ")

	hasExtention := false
	if strings.HasSuffix(baseName, ".md") {
		hasExtention = true
	}

	baseName = SanitiseDirPath(baseName)

	separators := regexp.MustCompile(`[/]`)
	baseName = separators.ReplaceAllString(baseName, "-")

	if !hasExtention {
		return fmt.Sprintf("%s.md", baseName)
	}

	return baseName
}
