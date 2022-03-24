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

// FeedItemImage is a single image.
type FeedItemImage struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Alt    string `json:"alt"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// NewFeedImage creates a new FeedItemImage from opengraph.Image.
func NewFeedImage(ogImg opengraph.Image, fallBackHost string) *FeedItemImage {
	img := &FeedItemImage{
		Alt:    ogImg.Alt,
		Height: ogImg.Height,
		Width:  ogImg.Width,
	}
	uri := ogImg.URL
	if secureURL := ogImg.SecureURL; secureURL != "" {
		uri = secureURL
	}
	img.URL = proxyImgSrc(uri, fallBackHost)
	return img
}

// proxyAllImgSrc replaces all image src with proxied image src.
func proxyAllImgSrc(r io.Reader, imgFallbackHost string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not parse html")
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		proxied := proxyImgSrc(src, imgFallbackHost)
		s.SetAttr("src", proxied)
	})
	out, err := doc.Html()
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not convert html doc to string")
	}
	return out, nil
}

// proxyImgSrc returns a proxied image src.
func proxyImgSrc(uri string, fallBackHost string) string {
	base, _ := url.Parse(fallBackHost)
	escapedURI := url.QueryEscape(joinToAbsolute(base, uri))
	return fmt.Sprintf("http://localhost:8080/api/v1/images?imageUrl=%s", escapedURI)
}

// joinToAbsolute returns a absolute url if relpath is relative.
func joinToAbsolute(base *url.URL, relpath string) string {
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
