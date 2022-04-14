package contentquality

import "strings"

var lowQualityIndicators = []string{
	// bloomberg
	"Are you a robot",
	// cloudflare
	"You have been blocked",
	// github
	"You can’t perform that action at this time.",
	//twitter
	"switch to a supported browser to continue using twitter.com",
	// youtube
	"https://www.youtube.com/howyoutubeworks",
}

func IsLow(content []byte) bool {
	str := string(content)
	for _, indicator := range lowQualityIndicators {
		if strings.Contains(str, indicator) {
			return true
		}
	}
	return false
}
