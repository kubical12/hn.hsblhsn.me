package readerviews

import "github.com/microcosm-cc/bluemonday"

var sanitizationPolicy = bluemonday.UGCPolicy()

func init() {
	sanitizationPolicy.
		AllowAttrs("target").
		OnElements("a").
		AllowAttrs("bionic-bold").
		OnElements("b").
		AllowAttrs("bionic-span").
		OnElements("span")
}
