package readerviews

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/images"
	"github.com/microcosm-cc/bluemonday"
)

// toAbs returns a absolute url if relpath is relative.
func toAbs(base, relpath string) string {
	baseURL, err := url.Parse(base)
	if err != nil || baseURL.Host == "" {
		return relpath
	}
	if baseURL.Scheme == "" {
		baseURL.Scheme = "https"
	}

	relURL, err := url.Parse(relpath)
	if err == nil && relURL.IsAbs() {
		return relURL.String()
	}
	if strings.HasPrefix(relpath, "//") {
		return fmt.Sprintf("%s:%s", baseURL.Scheme, relpath)
	}
	if strings.HasPrefix(relpath, "/") {
		return fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, relpath)
	}
	return fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, path.Join(baseURL.Path, relpath))
}

var sanitizationPolicy = bluemonday.UGCPolicy()

func Sanitize(html string, contentLink string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			s.Remove()
			return
		}
		absSrc := toAbs(contentLink, src)
		proxied := images.ProxiedURL(absSrc, images.ImageSizeFull)
		s.SetAttr("src", proxied)
	})
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			s.Remove()
			return
		}
		if strings.HasPrefix(href, "#") {
			// this is a fragment link.
			return
		}
		absLink := toAbs(contentLink, href)
		s.SetAttr("href", absLink)
		s.SetAttr("target", "_blank")
	})
	content, err := doc.Html()
	if err != nil {
		return "", err
	}
	content = sanitizationPolicy.Sanitize(content)
	return content, nil
}
