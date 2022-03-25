package hackernews

import (
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/otiai10/opengraph/v2"
	"github.com/pkg/errors"
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
func NewFeedImage(ogImg opengraph.Image, proxyEndpoint, feedItemLink string) *FeedItemImage {
	img := &FeedItemImage{
		Alt:    ogImg.Alt,
		Height: ogImg.Height,
		Width:  ogImg.Width,
	}
	uri := ogImg.URL
	if secureURL := ogImg.SecureURL; secureURL != "" {
		uri = secureURL
	}
	img.URL = proxyImgSrc(uri, proxyEndpoint, feedItemLink, ImageSizeThumbnail)
	return img
}

// proxyAllImgSrc replaces all image src with proxied image src.
func proxyAllImgSrc(r io.Reader, proxyEndpoint, feedItemLink string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not parse html")
	}
	var imgIterator = func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		proxied := proxyImgSrc(src, proxyEndpoint, feedItemLink, ImageSizeFull)
		s.SetAttr("src", proxied)
	}
	doc.Find("img").Each(imgIterator)
	out, err := doc.Html()
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not convert html doc to string")
	}
	return out, nil
}

// proxyImgSrc returns a proxied image src.
func proxyImgSrc(imgSrc, proxyEndpoint, feedItemLink string, size ImageSize) string {
	base, _ := url.Parse(feedItemLink)
	escapedURI := url.QueryEscape(joinUrls(base, imgSrc))
	return fmt.Sprintf("%s?imageUrl=%s&size=%s", proxyEndpoint, escapedURI, size)
}

// joinUrls returns a absolute url if relpath is relative.
func joinUrls(base *url.URL, relpath string) string {
	if base == nil {
		return relpath
	}
	src, err := url.Parse(relpath)
	if err == nil && src.IsAbs() {
		return src.String()
	}
	if strings.HasPrefix(relpath, "//") {
		return fmt.Sprintf("%s:%s", base.Scheme, relpath)
	}
	if strings.HasPrefix(relpath, "/") {
		return fmt.Sprintf("%s://%s%s", base.Scheme, base.Host, relpath)
	}
	return fmt.Sprintf("%s://%s%s", base.Scheme, base.Host, path.Join(base.Path, relpath))
}
