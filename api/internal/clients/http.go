package clients

import (
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/pkg/errors"
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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "clients: error while creating items request")
	}
	resp, err := HTTP().Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "clients: error while fetching items")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("clients: error while fetching items: %d", resp.StatusCode)
	}
	reader := http.MaxBytesReader(nil, resp.Body, HTTPMaxBodyLimit)
	return reader, nil
}
