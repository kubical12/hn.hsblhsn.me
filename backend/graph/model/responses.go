package model

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/opengraphs"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/relays"
)

type Story struct {
	*hackernews.ItemResponse
	ExternalContentLoadable
}

type Comment struct {
	*hackernews.ItemResponse
}

type Job struct {
	*hackernews.ItemResponse
	ExternalContentLoadable
}

type Poll struct {
	*hackernews.ItemResponse
}

type PollOption struct {
	*hackernews.ItemResponse
}

type PageInfo = relays.PageInfo

type (
	CommentConnection = relays.Connection[*Comment]
	CommentEdge       = relays.Edge[*Comment]
)

type (
	PollOptionConnection = relays.Connection[*PollOption]
	PollOptionEdge       = relays.Edge[*PollOption]
)

type (
	StoryConnection = relays.Connection[*Story]
	StoryEdge       = relays.Edge[*Story]
)

type (
	JobConnection = relays.Connection[*Job]
	JobEdge       = relays.Edge[*Job]
)

type (
	OpenGraph = opengraphs.OpenGraph
	Image     = opengraphs.Image
	Favicon   = opengraphs.Favicon
)

func (s *Story) Opengraph(ctx context.Context) (*opengraphs.OpenGraph, error) {
	loader := s.GetLoader(s.URL)
	return loader.Opengraph(ctx)
}

func (s *Story) HTML(ctx context.Context) (*string, error) {
	loader := s.GetLoader(s.URL)
	return loader.HTML(ctx)
}

func (s *Job) Opengraph(ctx context.Context) (*opengraphs.OpenGraph, error) {
	loader := s.GetLoader(s.URL)
	return loader.Opengraph(ctx)
}

func (s *Job) HTML(ctx context.Context) (*string, error) {
	loader := s.GetLoader(s.URL)
	return loader.HTML(ctx)
}
