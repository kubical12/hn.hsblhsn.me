package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/clients"
	"github.com/hsblhsn/queues"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// hnFeedItemResponse is the response from HackerNews API.
type hnFeedItemResponse struct {
	Id     int    `json:"id"`
	Type   string `json:"type"`
	Time   int    `json:"time"`
	Text   string `json:"text"`
	Parent int    `json:"parent"`
	Kids   []int  `json:"kids"`
	Url    string `json:"url"`
	Score  int    `json:"score"`
	Title  string `json:"title"`
}

// HackerNews is a HackerNews client.
type HackerNews struct {
	imageProxyURL string
	cache         map[FeedKind]IDList
	mu            sync.RWMutex
}

// NewHackerNews returns a new HackerNews client.
func NewHackerNews(imgProxyURL string) *HackerNews {
	return &HackerNews{
		imageProxyURL: imgProxyURL,
		cache:         make(map[FeedKind]IDList),
		mu:            sync.RWMutex{},
	}
}

// GetFeed returns a list of FeedItem.
// If the context is canceled, it returns an error.
func (h *HackerNews) GetFeed(ctx context.Context, kind FeedKind, page int) (Feed, error) {
	idList, err := h.getFeedItemIds(ctx, kind)
	if err != nil {
		return Feed{FeedItems: []*FeedItem{}}, err
	}
	if idList.Len() == 0 {
		return Feed{FeedItems: []*FeedItem{}}, nil
	}
	ids := idList.Paginate(page)
	items := h.getFeedItems(ctx, ids)
	return Feed{FeedItems: items}, nil
}

// GetFeedItem returns a single FeedItem.
func (h *HackerNews) GetFeedItem(ctx context.Context, id int, readability bool) (*FeedItem, error) {
	endpoint := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", id)
	reader, err := clients.SendHTTPRequest(ctx, endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "hackernews: error while fetching item")
	}
	defer reader.Close()
	var itemResp hnFeedItemResponse
	if err := json.NewDecoder(reader).Decode(&itemResp); err != nil {
		return nil, errors.Wrap(err, "hackernews: error while decoding item")
	}
	item := NewFeedItemFromHN(&itemResp, h.imageProxyURL)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return item.UseOpengraph(ctx)
	})
	if readability {
		g.Go(func() error {
			return item.UseReadability(ctx)
		})
	}
	if err := g.Wait(); err != nil {
		return nil, errors.Wrap(err, "hackernews: error while fetching item contents")
	}

	return item, nil
}

// getFeedItemIds fetches the list of IDs of feed kind.
func (h *HackerNews) getFeedItemIds(ctx context.Context, kind FeedKind) (IDList, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.cache == nil {
		h.cache = make(map[FeedKind]IDList)
	}
	if ids, ok := h.cache[kind]; ok {
		return ids, nil
	}
	endpoint := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/%s.json", kind)
	reader, err := clients.SendHTTPRequest(ctx, endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "hackernews: error while fetching items")
	}
	defer reader.Close()
	var itemList []int
	if err := json.NewDecoder(reader).Decode(&itemList); err != nil {
		return nil, errors.Wrap(err, "hackernews: error while decoding items")
	}
	h.cache[kind] = itemList
	return h.cache[kind], nil
}

// getFeedItems fetches the list of FeedItem from a list of IDs.
func (h *HackerNews) getFeedItems(ctx context.Context, ids []int) []*FeedItem {
	var fetchItem = func(q *queues.Q, items []*FeedItem, index, id int) {
		defer q.Done()
		item, err := h.GetFeedItem(ctx, id, false)
		if err != nil {
			log.Println(err)
		}
		items[index] = item
	}
	q := queues.New(10)
	items := make([]*FeedItem, len(ids))
	for index, id := range ids {
		q.Add(1)
		go fetchItem(q, items, index, id)
	}
	q.Wait()
	return items
}
