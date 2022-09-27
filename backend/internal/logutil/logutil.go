package logutil

import "strings"

//nolint:gochecknoglobals // this is a global variable for performance.
var sanitizationReplacer = strings.NewReplacer(
	"\r", "",
	"\n", "",
)

func Sanitize(s string) string {
	return sanitizationReplacer.Replace(s)
}
