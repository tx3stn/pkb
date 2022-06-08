// Package date contains logic for date based requirements/formatting not
// handled by the standard lib.
package date

import (
	"fmt"
	"regexp"
	"strconv"
)

// DaySuffix returns the string for the current date when you need a more
// human friendly date suffix.
func DaySuffix(x int) string {
	suffix := "th"
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			suffix = "st"
		}
	case 2:
		if x%100 != 12 {
			suffix = "nd"
		}
	case 3:
		if x%100 != 13 {
			suffix = "rd"
		}
	}
	return suffix
}

// IncludesSuffixFormat checks if the supplied date format string includes a
// day suffix identifier so you can check if it should be replaced with the
// correct value or not.
func IncludesSuffixFormat(input string) bool {
	re := regexp.MustCompile(`(?P<start>.*\d+)(?P<suffix>th|st|nd|rd){1}(?P<end>.*)`)
	return re.MatchString(input)
}

// ReplaceSuffixFormatter updates the date string to use the correct suffux value.
func ReplaceSuffixFormatter(formatString string) (string, error) {
	re := regexp.MustCompile(`(?P<start>.*)(?P<day>\d+)(?P<suffix>th|st|nd|rd){1}(?P<end>.*)`)
	dayVal := getCaptureGroupValueFromString(re, formatString, "day")
	if dayVal == "" {
		return "", fmt.Errorf("could not find group day in string %s", formatString)
	}

	dayInt, err := strconv.ParseInt(dayVal, 10, 0)
	if err != nil {
		return "", err
	}

	suffix := DaySuffix(int(dayInt))
	s := re.ReplaceAllString(formatString, fmt.Sprintf("${start}${day}%s${end}", suffix))
	return s, nil
}

// getCaptureGroupValueFromString searches the regex match for the capture
// group with the specified name.
func getCaptureGroupValueFromString(re *regexp.Regexp, search string, groupName string) string {
	groupNames := re.SubexpNames()
	for _, match := range re.FindAllStringSubmatch(search, -1) {
		for groupID, groupValue := range match {
			if groupNames[groupID] == groupName {
				return groupValue
			}
		}
	}

	return ""
}
