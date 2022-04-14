package contentquality

import (
	"strings"
	"sync"
)

// nolint:gochecknoglobals // global for lazy loading
var (
	lowQualityIndicators = []string{}
	initializeOnce       = sync.Once{}
)

func getIndicators() []string {
	initializeOnce.Do(func() {
		lowQualityIndicators = []string{
			// bloomberg
			"Are you a robot",
			// cloudflare
			"You have been blocked",
			// github
			"You can’t perform that action at this time.",

			"switch to a supported browser to continue using twitter.com",
			// youtube
			"https://www.youtube.com/howyoutubeworks",
		}
	})
	return lowQualityIndicators
}

func IsLow(content []byte) bool {
	str := string(content)
	indicators := getIndicators()
	for _, indicator := range indicators {
		if strings.Contains(str, indicator) {
			return true
		}
	}
	return false
}
