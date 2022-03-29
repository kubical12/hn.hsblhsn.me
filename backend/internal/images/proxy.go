package images

import (
	"fmt"
	"net/url"
)

func ProxiedURL(src string, size ImageSize) string {
	if src == "" {
		return ""
	}
	return fmt.Sprintf("/api/v1/image.jpeg?src=%s&size=%s", url.QueryEscape(src), size)
}
