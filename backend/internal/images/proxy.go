package images

import (
	"fmt"
	"net/url"
)

func ProxiedURL(src string, size ImageSize) string {
	if src == "" {
		return ""
	}
	return fmt.Sprintf("/images.jpeg?src=%s&size=%s", url.QueryEscape(src), size)
}
