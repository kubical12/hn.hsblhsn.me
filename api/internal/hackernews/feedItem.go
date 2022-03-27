package hackernews

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"sync"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/clients"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityclient"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/images"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/readerviews"
	"github.com/microcosm-cc/bluemonday"
	"github.com/otiai10/opengraph/v2"
	"github.com/pkg/errors"
)

// HTMLViewerPolicy is the policy for the html sanitizer.
var HTMLViewerPolicy = bluemonday.UGCPolicy()

// FeedItem is a single post.
// It uses opengraph data to fill in missing fields.
type FeedItem struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	HTML          string `json:"__html"`
	Domain        string `json:"domain"`
	URL           string `json:"url"`
	HackerNewsUrl string `json:"hackerNewsUrl"`
	ThumbnailUrl  string `json:"thumbnailUrl"`
	TotalPoints   int    `json:"totalPoints"`
	TotalComments int    `json:"totalComments"`
	mu            sync.Mutex
}

// NewFeedItemFromHN converts a hackernews post to a FeedItem.
func NewFeedItemFromHN(resp *hnFeedItemResponse) *FeedItem {
	hnLink := fmt.Sprintf("https://news.ycombinator.com/item?id=%d", resp.ID)
	var domain string
	u, err := url.Parse(resp.URL)
	if err == nil {
		domain = u.Host
	}

	return &FeedItem{
		ID:            resp.ID,
		Title:         resp.Title,
		Body:          resp.Text,
		URL:           resp.URL,
		Domain:        domain,
		HackerNewsUrl: hnLink,
		TotalPoints:   resp.Score,
		TotalComments: len(resp.Kids),
	}
}

// UseOpengraph fetches the opengraph data from the linked website.
// It overwrites the original title and body with the opengraph data.
func (f *FeedItem) UseOpengraph(ctx context.Context) error {
	og, err := opengraph.Fetch(f.URL, opengraph.Intent{
		Context:    ctx,
		HTTPClient: clients.HTTP(),
		Strict:     false,
	})
	if err != nil {
		return errors.Wrap(err, "hackernews: could not fetch opengraph url")
	}
	if err := og.ToAbs(); err != nil {
		return errors.Wrap(err, "hackernews: could not convert relative url to absolute url")
	}
	var bestImage opengraph.Image
	if len(og.Image) > 0 {
		bestImage = og.Image[0]
	}
	for _, img := range og.Image {
		if img.Width > bestImage.Width {
			bestImage = img
		}
	}

	// add images.
	// overwrite the original title and body with the opengraph data
	// if the opengraph data is available.
	f.mu.Lock()
	defer f.mu.Unlock()
	if bestImage.URL != "" {
		f.ThumbnailUrl = images.ProxiedURL(bestImage.URL, images.ImageSizeThumbnail)
	}
	if og.Title != "" {
		f.Title = og.Title
	}
	if og.Description != "" {
		f.Body = og.Description
	}
	return nil
}

// UseReadability returns a readable version of the post.
// It fetches the original content from linked website.
// Then proxies all the images on the page.
// Then converts the webpage to a readable html document.
// At last it sanitizes the html document.
func (f *FeedItem) UseReadability(ctx context.Context) error {
	// get the original content from the linked website
	reader, err := clients.SendHTTPRequest(ctx, f.URL)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not get readability data")
	}
	contentBytes, err := io.ReadAll(reader)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not read readability data")
	}

	// convert the webpage to a readable html document
	ready := clients.IsReadabilityClientReady(ctx, time.Second*3)
	if !ready {
		return errors.New("hackernews: readability client not ready")
	}
	rc := clients.ReadabilityClient()
	resp, err := rc.GetReadableDocument(ctx, &readabilityclient.GetReadableDocumentRequest{
		Html:       string(contentBytes),
		Identifier: f.URL,
	})
	if err != nil {
		return errors.Wrap(err, "hackernews: error while calling readability")
	}
	// sanitize the html document
	htmlContent, err := readerviews.Sanitize(resp.GetBody(), f.URL)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not sanitize html")
	}
	f.HTML = htmlContent

	// prepare the title.
	if title := resp.GetTitle(); title != "" {
		f.Title = title
	}

	return nil
}
