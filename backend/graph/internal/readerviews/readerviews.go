package readerviews

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/bionify"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/grpc/readabilityclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/images"
	"github.com/pkg/errors"
)

const (
	ClientReadinessTimeout = time.Second * 3
)

func Convert(ctx context.Context, link string, content *bytes.Buffer) (string, error) {
	// convert the webpage to a readable html document
	ready := isReadabilityClientReady(ctx, ClientReadinessTimeout)
	if !ready {
		return "", errors.New("hackernews: readability client not ready")
	}
	rc := readabilityClient()
	resp, err := rc.GetReadableDocument(ctx, &readabilityclient.GetReadableDocumentRequest{
		Html:       content.String(),
		Identifier: link,
	})
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not call readability server")
	}
	return TransformHTML(link, strings.NewReader(resp.GetBody()))
}

func TransformHTML(link string, content io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		return "", errors.Wrap(err, "readerviews: could not parse html")
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, ok := selection.Attr("src")
		if !ok || src == "" {
			selection.Remove()
			return
		}
		absSrc := toAbs(link, src)
		proxied := images.ProxiedURL(absSrc, images.ImageSizeFull)
		selection.SetAttr("src", proxied)
	})
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if !ok || href == "" {
			selection.Remove()
			return
		}
		absLink := toAbs(link, href)
		selection.SetAttr("href", absLink)
		selection.SetAttr("target", "_blank")
	})
	doc.Find("p").Each(func(i int, selection *goquery.Selection) {
		selection.SetHtml(bionify.Text(selection.Text()))
	})
	htmlContent, err := doc.Html()
	if err != nil {
		return "", errors.Wrap(err, "readerviews: error while rendering html")
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
