package strutil

import (
	"regexp"
)

var (
	invalidLabelCharRE = regexp.MustCompile(`[^a-zA-Z0-9_]`)
)

// SanitizeLabelName replaces anything that doesn't match
// client_label.LabelNameRE with an underscore.
func SanitizeLabelName(name string) string {
	return invalidLabelCharRE.ReplaceAllString(name, "_")
}
