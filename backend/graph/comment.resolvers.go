package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/model"
)

func (r *commentResolver) Type(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.Type.String(), nil
}

func (r *commentResolver) Comments(ctx context.Context, obj *model.Comment, after *string, first *int) (*relays.Connection[*model.Comment], error) {
	relayResolver := r.NewRelayComments(ctx, obj.Kids)
	comments, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate comments on the comment")
	}
	return comments, nil
}

func (r *commentResolver) Parent(ctx context.Context, obj *model.Comment) (string, error) {
	return strconv.Itoa(obj.Parent), nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
