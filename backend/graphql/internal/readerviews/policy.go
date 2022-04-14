package readerviews

import "github.com/microcosm-cc/bluemonday"

// nolint:gochecknoglobals // This is a stateful operation.
var sanitizationPolicy = bluemonday.UGCPolicy().
	AllowAttrs("target").
	OnElements("a").
	AllowAttrs("bionic-bold").
	OnElements("b").
	AllowAttrs("bionic-span").
	OnElements("span")
