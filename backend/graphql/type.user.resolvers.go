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

// DatabaseID is the resolver for the databaseId field.
func (r *userResolver) DatabaseID(ctx context.Context, obj *model.User) (string, error) {
	return obj.UserResponse.ID, nil
}

// Submitted is the resolver for the submitted field.
func (r *userResolver) Submitted(ctx context.Context, obj *model.User, after *string, first *int) (*relays.Connection[model.Node], error) {
	relayResolver := r.NewRelaySubmitted(ctx, obj.Submitted)
	nodes, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate submitted nodes")
	}
	return nodes, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
