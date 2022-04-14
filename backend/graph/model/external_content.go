package model

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/opengraphs"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/readerviews"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// nolint:gochecknoglobals // global client to call from ExternalContentLoader.
var (
	client *httpclient.CachedClient
	logger *zap.Logger
)

func registerDependencies(c *httpclient.CachedClient, l *zap.Logger) {
	client = c
	logger = l.With(zap.String("component", "model_external_content"))
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

func (s *ExternalContentLoader) Opengraph(ctx context.Context) *opengraphs.OpenGraph {
	if err := s.getContentFromURL(ctx); err != nil {
		logger.Error("model: could not get content from url", zap.Error(err))
		return nil
	}

	if s.url == "" {
		return nil
	}
	out, err := opengraphs.GetOpengraphData(s.url, bytes.NewBuffer(s.content))
	logger.Error("model: could not convert opengraph content", zap.Error(err))
	return out
}

func (s *ExternalContentLoader) HTML(ctx context.Context) *string {
	if err := s.getContentFromURL(ctx); err != nil {
		logger.Error("model: could not get content from url", zap.Error(err))
		return nil
	}

	if s.url == "" {
		return nil
	}
	out, err := readerviews.Convert(ctx, s.url, bytes.NewBuffer(s.content))
	logger.Error("model: could not convert readerview content", zap.Error(err))
	return &out
}
