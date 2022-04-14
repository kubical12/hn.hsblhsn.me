package images

import (
	"fmt"
	"net/url"
	"os"
)

func ProxiedURL(src string, size ImageSize) string {
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
