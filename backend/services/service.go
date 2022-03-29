package services

import (
	"bytes"
	"context"
	"image"
	"log"
	"net/http"
	"net/url"

	"github.com/hsblhsn/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/types"
	"github.com/hsblhsn/queues"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	GetListByType(ctx context.Context, typ types.ListType, page uint8) (*types.List, error)
	GetItemByID(ctx context.Context, id uint32) (*types.Item, error)
	GetResizedImage(ctx context.Context, uri string, size images.ImageSize) (image.Image, error)
}

type service struct {
	hn             *hackernews.APIClient
	client         *http.Client
	maxConcurrency uint
}

// NewService returns a highly concurrent implementation of the Service interface.
func NewService(client *hackernews.APIClient) Service {
	return &service{
		hn:             client,
		client:         http.DefaultClient,
		maxConcurrency: 10,
	}
}

// GetListByType fetches all the item ids of the list type.
// Paginates the list of ids.
// Fetches and returns the the items.
// This implementation is highly concurrent.
func (s *service) GetListByType(ctx context.Context, typ types.ListType, page uint8) (*types.List, error) {
	result, _, err := s.hn.LiveDataApi.GETTopstoriesJson(ctx, nil)
	if err != nil {
		return nil, err
	}
	var (
		idsToFetch = types.NewIDList(result).Paginate(int(page))
		items      = make([]*types.Item, len(idsToFetch))
	)

	// closure function to fetch items concurrently.
	fetch := func(ctx context.Context, q *queues.Q, index int, id uint32) {
		defer q.Done()
		item, err := s.GetItemByID(disableReadability(ctx), id)
		if err != nil {
			log.Println(err)
		}
		items[index] = item
	}
	// fire up all the items to be fetched in a concurrent queue.
	q := queues.New(s.maxConcurrency)
	for index, id := range idsToFetch {
		q.Add(1)
		go fetch(ctx, q, index, uint32(id))
	}

	// wait and return the items when fetched.
	q.Wait()
	return &types.List{Type: typ, Page: page, Items: items}, nil
}

// GetItemByID fetches a single item with the given id.
func (s *service) GetItemByID(ctx context.Context, id uint32) (*types.Item, error) {
	result, _, err := s.hn.ItemsApi.GETItemItemIdJson(ctx, int32(id), nil)
	if err != nil {
		return nil, err
	}
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("unexpected type %T", result)
	}
	item := new(types.Item)
	if err := mapstructure.WeakDecode(resultMap, item); err != nil {
		return nil, err
	}

	// parse the url to get hostname.
	{
		uri, err := url.Parse(item.URL)
		if err != nil {
			return nil, err
		}
		if uri.Scheme != "http" && uri.Scheme != "https" {
			return nil, nil
		}
		item.Domain = uri.Host
	}
	content, err := getContentFromURL(ctx, item.URL, DefaultMaxHTTPBytes)
	if err != nil {
		return nil, err
	}
	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return getReadableContent(gCtx, item, content)
	})
	g.Go(func() error {
		return getSEOData(gCtx, item, content)
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return item, nil
}

// GetResizedImage gets an image from the url and resizes it based on the given size.
func (s *service) GetResizedImage(ctx context.Context, uri string, size images.ImageSize) (image.Image, error) {
	content, err := getContentFromURL(ctx, uri, DefaultMaxHTTPBytes)
	if err != nil {
		return nil, err
	}
	img, err := images.Resize(bytes.NewReader(content), size)
	if err != nil {
		return nil, err
	}
	return img, nil
}
