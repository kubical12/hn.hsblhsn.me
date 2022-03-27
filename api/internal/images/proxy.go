package images

import (
	"fmt"
	"net/url"
)

func ProxiedURL(src string, size ImageSize) string {
	return fmt.Sprintf("/api/v1/feed_images?imageUrl=%s&size=%s", url.QueryEscape(src), size)
}
