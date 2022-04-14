package featureflags

import (
	"os"
	"strings"
)

func Check(feature string, def bool) bool {
	val := os.Getenv("FEATURE_FLAG_" + strings.ToUpper(feature))
	switch val {
	case "on":
		return true
	case "off":
		return false
	default:
		return def
	}
}
