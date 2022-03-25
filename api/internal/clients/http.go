package clients

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/hsblhsn/hn.hsblhsn.me/api/internal/caches"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	HTTPMaxBodyLimit int64 = 1024 * 1024
)

var (
	httpSingletonClient *http.Client
	httpOnce            = sync.Once{}
)

// HTTP returns a singleton http client.
func HTTP() *http.Client {
	httpOnce.Do(func() {
		httpSingletonClient = &http.Client{}
	})
	return httpSingletonClient
}

// SendHTTPRequest sends an HTTP request.
// If the context is canceled, it returns an error.
func SendHTTPRequest(ctx context.Context, url string) (io.ReadCloser, error) {
	cache := caches.Cache()
	val, err := cache.Get([]byte(url))
	if err == nil {
		return ioutil.NopCloser(bytes.NewReader(val)), nil
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "clients: error while creating items request")
	}
	req.Header.Set("User-Agent", "Hackernews Reader By @hsblhsn")
	resp, err := HTTP().Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "clients: error while fetching items")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("clients: error while fetching items: %d", resp.StatusCode)
	}
	reader := http.MaxBytesReader(nil, resp.Body, HTTPMaxBodyLimit)

	val, err = ioutil.ReadAll(reader)
	if err != nil {
		return nil, errors.Wrap(err, "clients: error while reading response body")
	}
	if err := cache.Set([]byte(url), val, 600); err != nil {
		// just log the error
		zap.L().Error("clients: error while caching response", zap.Error(err))
	}
	return ioutil.NopCloser(bytes.NewReader(val)), nil
}
