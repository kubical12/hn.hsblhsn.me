package services

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/readerviews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/types"
	"github.com/otiai10/opengraph/v2"
	"github.com/pkg/errors"
)

const (
	DefaultMaxHTTPBytes = 1024 * 1024
)

func getSEOData(ctx context.Context, item *types.Item, content []byte) error {
	og := &opengraph.OpenGraph{
		Intent: opengraph.Intent{
			URL: item.URL,
		},
	}
	if err := og.Parse(bytes.NewReader(content)); err != nil {
		return err
	}
	if err := og.ToAbs(); err != nil {
		return err
	}

	// fill data from opengraph if available.
	if og.Title != "" {
		item.Title = og.Title
	}
	if og.Description != "" {
		item.Summary = og.Description
	}
	bestImageUrl := getBestImageURL(og.Image)

	// fill thumbnail data.
	item.ThumbnailUrl = images.ProxiedURL(bestImageUrl, images.ImageSizeThumbnail)
	// fill seo data.
	item.SEO = &types.SEO{
		Title:       og.Title,
		Description: og.Description,
		ImageURL:    images.ProxiedURL(bestImageUrl, images.ImageSizeFull),
	}
	return nil
}

func getBestImageURL(imgList []opengraph.Image) string {
	if len(imgList) == 0 {
		return ""
	}
	best := imgList[0]
	for _, v := range imgList {
		if v.Width > best.Width {
			best = v
		}
	}
	return best.URL
}

func getReadableContent(ctx context.Context, item *types.Item, content []byte) error {
	// check if readability is disabled.
	val := ctx.Value(disableReadabilityCtxKey{})
	if _, ok := val.(struct{}); ok {
		return nil
	}

	readableContent, err := readerviews.Sanitize(ctx, content, item.URL)
	if err != nil {
		return err
	}

	// fetch and prepare readable content.
	item.Content = readableContent
	return nil
}

type disableReadabilityCtxKey struct{}

func disableReadability(ctx context.Context) context.Context {
	return context.WithValue(ctx, disableReadabilityCtxKey{}, struct{}{})
}

func getContentFromURL(ctx context.Context, uri string, maxBytes int64) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "services: could not build request to get http content")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "services: could not send request to get http content")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("service: status code is not 200 while fetching http content")
	}
	defer resp.Body.Close()
	reader := io.LimitReader(resp.Body, maxBytes)

	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.New("service: could not read fetched content")
	}
	return b, nil
}

func getDomainName(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	return u.Host
}
