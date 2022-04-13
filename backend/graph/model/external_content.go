package model

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/opengraphs"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/readerviews"
	"github.com/pkg/errors"
)

// nolint:gochecknoglobals // global client to call from ExternalContentLoader.
var client *httpclient.CachedClient

func registerClient(c *httpclient.CachedClient) {
	client = c
}

type ExternalContentLoadable struct {
	ext  *ExternalContentLoader
	once sync.Once
}

func (s *ExternalContentLoadable) GetLoader(url string) *ExternalContentLoader {
	s.once.Do(func() {
		s.ext = NewExternalContentLoader(url)
	})
	return s.ext
}

type ExternalContentLoader struct {
	url     string
	content []byte
	once    sync.Once
	mu      sync.Mutex
}

func NewExternalContentLoader(url string) *ExternalContentLoader {
	return &ExternalContentLoader{url: url}
}

func (s *ExternalContentLoader) getContentFromURL(ctx context.Context) error {
	if s.url == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	var err error
	s.once.Do(func() {
		var resp *http.Response
		resp, err = client.Get(ctx, s.url)
		if err != nil {
			return
		}
		defer resp.Body.Close()
		s.content, err = io.ReadAll(resp.Body)
	})
	return errors.Wrap(err, "model: could not get content from url")
}

func (s *ExternalContentLoader) Opengraph(ctx context.Context) (*opengraphs.OpenGraph, error) {
	if err := s.getContentFromURL(ctx); err != nil {
		return nil, err
	}
	// nolint:nilnil // ignore
	if s.url == "" {
		return nil, nil
	}
	out, err := opengraphs.GetOpengraphData(s.url, bytes.NewBuffer(s.content))
	return out, errors.Wrap(err, "model: could not convert opengraph content")
}

func (s *ExternalContentLoader) HTML(ctx context.Context) (*string, error) {
	if err := s.getContentFromURL(ctx); err != nil {
		return nil, err
	}
	// nolint:nilnil // ignore
	if s.url == "" {
		return nil, nil
	}
	out, err := readerviews.Convert(ctx, s.url, bytes.NewBuffer(s.content))
	return &out, errors.Wrap(err, "model: could not convert readerview content")
}
