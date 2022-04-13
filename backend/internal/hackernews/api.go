package hackernews

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/bionify"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"github.com/pkg/errors"
)

type HackerNews struct {
	client *httpclient.CachedClient
}

func NewHackerNews(client *httpclient.CachedClient) *HackerNews {
	return &HackerNews{client: client}
}

func (h *HackerNews) API(ctx context.Context, endpoint string, result any) error {
	// Send the request
	resp, err := h.client.Get(ctx, endpoint)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not send request")
	}
	defer resp.Body.Close()
	// Decode the response body to the target struct
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not decode response")
	}
	return nil
}

func (h *HackerNews) GetTopStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/topstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetNewStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/newstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetAskStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/askstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetShowStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/showstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetJobStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/jobstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// GetTypedItem http request to the endpoint and decode json data to the target struct,.
func (h *HackerNews) GetTypedItem(ctx context.Context, typ ItemType, itemID string) (*ItemResponse, error) {
	idN, idErr := GetIntID(itemID)
	if idErr != nil {
		return nil, idErr
	}
	out, err := h.GetItem(ctx, idN)
	if err != nil {
		return nil, err
	}
	if out.Type != typ {
		msg := fmt.Sprintf(
			"type: %q!=%q on id=%q",
			out.Type, typ, itemID,
		)
		return nil, errors.Wrap(ErrMismatchedType, msg)
	}
	return out, nil
}

func (h *HackerNews) GetItem(ctx context.Context, id int) (*ItemResponse, error) {
	var (
		endpoint = fmt.Sprintf("%s/v0/item/%d.json", host, id)
		out      = new(ItemResponse)
		err      = h.API(ctx, endpoint, out)
	)
	if err != nil {
		return nil, err
	}
	out.Text, err = bionify.HTMLText(out.Text)
	if err != nil {
		return nil, errors.Wrap(err, "hackernews: could not bionify item text")
	}
	return out, nil
}
