package contentquality

import "strings"

var lowQualityIndicators = []string{
	"Are you a robot",       // bloomberg
	"You have been blocked", // cloudflare
	"You signed in with another tab or window. to refresh your session", // github
	"switch to a supported browser to continue using twitter.com",       //twitter                                    // twitter
	"https://www.youtube.com/howyoutubeworks",                           // youtube
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
