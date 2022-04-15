package images

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/hsblhsn/hn.hsblhsn.me/featureflags"
)

func ProxiedURL(src string, size ImageSize) string {
	if !featureflags.IsOn(featureflags.FeatureImgProxy, false) {
		return src
	}
	if src == "" {
		return ""
	}
	host := os.Getenv("DOMAIN")
	if host != "" {
		host = "https://" + host
	}
	escaped := url.QueryEscape(src)
	return fmt.Sprintf(
		"%s/images.jpeg?src=%s&size=%s",
		host,
		unescape(escaped),
		size,
	)
}

func unescape(str string) string {
	str = strings.ReplaceAll(str, "%2C", `,`)
	str = strings.ReplaceAll(str, "%25", `%`)
	return str
}
