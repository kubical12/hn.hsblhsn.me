package hackernews

import (
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/images"
	"github.com/otiai10/opengraph/v2"
)

type ImageSize string

func (f ImageSize) String() string {
	switch f {
	case ImageSizeThumbnail, ImageSizeFull:
		return string(f)
	default:
		return string(ImageSizeFull)
	}
}

func (f ImageSize) Dimension() (height, width uint) {
	switch f {
	case ImageSizeThumbnail:
		return 180, 180
	case ImageSizeFull:
		return 560, 560
	default:
		return 560, 560
	}
}

const (
	// ImageSizeThumbnail is the size of the thumbnail image.
	ImageSizeThumbnail ImageSize = "thumbnail"
	// ImageSizeFull is the size of the full image.
	ImageSizeFull ImageSize = "full"
)

// FeedItemImage is a single image.
type FeedItemImage struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Alt    string `json:"alt"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// NewFeedImage creates a new FeedItemImage from opengraph.Image.
func NewFeedImage(ogImg opengraph.Image, feedItemLink string) *FeedItemImage {
	img := &FeedItemImage{
		Alt:    ogImg.Alt,
		Height: ogImg.Height,
		Width:  ogImg.Width,
	}
	uri := ogImg.URL
	if secureURL := ogImg.SecureURL; secureURL != "" {
		uri = secureURL
	}
	img.URL = images.ProxiedURL(uri, images.ImageSizeFull)
	return img
}
