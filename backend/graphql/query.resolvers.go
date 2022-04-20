package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	nodeID := id
	if hackernews.IsUserName(nodeID) {
		user, userErr := r.hackerNews.GetUser(ctx, nodeID)
		if userErr != nil {
			return nil, msgerr.New(userErr, "Could not find a user with the username")
		}
		return &model.User{UserResponse: user}, nil
	}
	idN, err := hackernews.GetIntID(nodeID)
	if err != nil {
		return nil, msgerr.New(err, "Invalid ID")
	}
	result, err := r.hackerNews.GetItem(ctx, idN)
	if err != nil {
		return nil, msgerr.New(err, "Could not get item")
	}
	return ItemToNode(result)
}

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

func (r *queryResolver) JobStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Job], error) {
	list, err := r.hackerNews.GetJobStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get jobs.")
	}
	relayResolver := r.NewRelayJobs(ctx, list)
	jobs, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate jobs.")
	}
	return jobs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
