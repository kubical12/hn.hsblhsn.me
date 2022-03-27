package hackernews

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/clients"
	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/grpc/readabilityclient"
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
	ID            int              `json:"id"`
	Title         string           `json:"title"`
	Body          string           `json:"body"`
	HTML          string           `json:"__html"`
	Link          string           `json:"link"`
	Images        []*FeedItemImage `json:"images"`
	TotalPoints   int              `json:"totalPoints"`
	TotalComments int              `json:"totalComments"`
	mu            sync.Mutex
}

// NewFeedItemFromHN converts a hackernews post to a FeedItem.
func NewFeedItemFromHN(resp *hnFeedItemResponse) *FeedItem {
	return &FeedItem{
		ID:            resp.Id,
		Title:         resp.Title,
		Body:          resp.Text,
		Link:          resp.Url,
		TotalPoints:   resp.Score,
		TotalComments: len(resp.Kids),
	}
}

// UseOpengraph fetches the opengraph data from the linked website.
// It overwrites the original title and body with the opengraph data.
func (f *FeedItem) UseOpengraph(ctx context.Context) error {
	og, err := opengraph.Fetch(f.Link, opengraph.Intent{
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
	// add images.
	// overwrite the original title and body with the opengraph data
	// if the opengraph data is available.
	f.mu.Lock()
	defer f.mu.Unlock()
	f.Images = make([]*FeedItemImage, len(og.Image))
	for i, img := range og.Image {
		f.Images[i] = NewFeedImage(img, og.URL)
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
	reader, err := clients.SendHTTPRequest(ctx, f.Link)
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
		Identifier: f.Link,
	})
	if err != nil {
		return errors.Wrap(err, "hackernews: error while calling readability")
	}
	// sanitize the html document
	htmlContent, err := readerviews.Sanitize(resp.GetBody(), f.Link)
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
