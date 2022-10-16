package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/model"
)

// Type is the resolver for the type field.
func (r *pollResolver) Type(ctx context.Context, obj *model.Poll) (string, error) {
	return obj.Type.String(), nil
}

// By is the resolver for the by field.
func (r *pollResolver) By(ctx context.Context, obj *model.Poll) (*model.User, error) {
	resp, err := r.hackerNews.GetUser(ctx, obj.By)
	if err != nil {
		return nil, msgerr.New(err, "Could not get user")
	}
	return &model.User{UserResponse: resp}, nil
}

// Comments is the resolver for the comments field.
func (r *pollResolver) Comments(ctx context.Context, obj *model.Poll, after *string, first *int) (*relays.Connection[*model.Comment], error) {
	relayResolver := r.NewRelayComments(ctx, obj.Kids)
	comments, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not get comments on the poll")
	}
	return comments, nil
}

// PollOptions is the resolver for the pollOptions field.
func (r *pollResolver) PollOptions(ctx context.Context, obj *model.Poll, after *string, first *int) (*relays.Connection[*model.PollOption], error) {
	relayResolver := r.NewRelayPollOptions(ctx, obj.Parts)
	options, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not get poll options")
	}
	return options, nil
}

// Type is the resolver for the type field.
func (r *pollOptionResolver) Type(ctx context.Context, obj *model.PollOption) (string, error) {
	return obj.Type.String(), nil
}

// By is the resolver for the by field.
func (r *pollOptionResolver) By(ctx context.Context, obj *model.PollOption) (*model.User, error) {
	resp, err := r.hackerNews.GetUser(ctx, obj.By)
	if err != nil {
		return nil, msgerr.New(err, "Could not get user")
	}
	return &model.User{UserResponse: resp}, nil
}

// Poll is the resolver for the poll field.
func (r *pollOptionResolver) Poll(ctx context.Context, obj *model.PollOption) (string, error) {
	return strconv.Itoa(obj.Poll), nil
}

// Poll returns generated.PollResolver implementation.
func (r *Resolver) Poll() generated.PollResolver { return &pollResolver{r} }

// PollOption returns generated.PollOptionResolver implementation.
func (r *Resolver) PollOption() generated.PollOptionResolver { return &pollOptionResolver{r} }

type (
	pollResolver       struct{ *Resolver }
	pollOptionResolver struct{ *Resolver }
)
