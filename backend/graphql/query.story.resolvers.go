package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

// TopStories is the resolver for the topStories field.
func (r *queryResolver) TopStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetTopStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get top stories.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate top stories.")
	}
	return stories, nil
}

// NewStories is the resolver for the newStories field.
func (r *queryResolver) NewStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetNewStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get new stories.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate new stories.")
	}
	return stories, nil
}

// AskStories is the resolver for the askStories field.
func (r *queryResolver) AskStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetAskStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get ask HN.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate ask HN.")
	}
	return stories, nil
}

// ShowStories is the resolver for the showStories field.
func (r *queryResolver) ShowStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetShowStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get show HN.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate show HN.")
	}
	return stories, nil
}
