package readerviews

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/grpc/readabilityclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
	"github.com/pkg/errors"
)

func Sanitize(ctx context.Context, content []byte, link string) (string, error) {
	// convert the webpage to a readable html document
	ready := isReadabilityClientReady(ctx, time.Second*3)
	if !ready {
		return "", errors.New("hackernews: readability client not ready")
	}
	rc := readabilityClient()
	resp, err := rc.GetReadableDocument(ctx, &readabilityclient.GetReadableDocumentRequest{
		Html:       string(content),
		Identifier: link,
	})
	if err != nil {
		return "", errors.Wrap(err, "hackernews: error while calling readability")
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.GetBody()))
	if err != nil {
		return "", err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok || src == "" {
			s.Remove()
			return
		}
		absSrc := toAbs(link, src)
		proxied := images.ProxiedURL(absSrc, images.ImageSizeFull)
		s.SetAttr("src", proxied)
	})
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok || href == "" {
			s.Remove()
			return
		}
		if strings.HasPrefix(href, "#") {
			// this is a fragment link.
			return
		}
		absLink := toAbs(link, href)
		s.SetAttr("href", absLink)
		s.SetAttr("target", "_blank")
	})
	htmlContent, err := doc.Html()
	if err != nil {
		return "", err
	}
	return sanitizationPolicy.Sanitize(htmlContent), nil
}

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
