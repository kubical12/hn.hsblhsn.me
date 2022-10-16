package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/model"
)

// Type is the resolver for the type field.
func (r *storyResolver) Type(ctx context.Context, obj *model.Story) (string, error) {
	return obj.Type.String(), nil
}

// By is the resolver for the by field.
func (r *storyResolver) By(ctx context.Context, obj *model.Story) (*model.User, error) {
	resp, err := r.hackerNews.GetUser(ctx, obj.By)
	if err != nil {
		return nil, msgerr.New(err, "Could not get user")
	}
	return &model.User{UserResponse: resp}, nil
}

// Comments is the resolver for the comments field.
func (r *storyResolver) Comments(ctx context.Context, obj *model.Story, after *string, first *int) (*relays.Connection[*model.Comment], error) {
	relayResolver := r.NewRelayComments(ctx, obj.Kids)
	comments, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate comments on the story")
	}
	return comments, nil
}

// Story returns generated.StoryResolver implementation.
func (r *Resolver) Story() generated.StoryResolver { return &storyResolver{r} }

type storyResolver struct{ *Resolver }
