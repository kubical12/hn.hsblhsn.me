package images

import (
	"fmt"
	"net/url"
	"os"

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
	return fmt.Sprintf(
		"%s/images.jpeg?src=%s&size=%s",
		host,
		url.QueryEscape(src),
		size,
	)
}
