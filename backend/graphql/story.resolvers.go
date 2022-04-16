package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

func (r *storyResolver) Type(ctx context.Context, obj *model.Story) (string, error) {
	return obj.Type.String(), nil
}

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
