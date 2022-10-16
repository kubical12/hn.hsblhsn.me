package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/model"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	var (
		nodeID    = id
		errOnUser error
	)
	if hackernews.IsUserName(nodeID) {
		user, err := r.hackerNews.GetUser(ctx, nodeID)
		if err != nil {
			errOnUser = err
			goto itemResolver
		}
		return &model.User{UserResponse: user}, nil
	}
itemResolver:
	idN, err := hackernews.GetIntID(nodeID)
	if err != nil {
		if errOnUser != nil {
			return nil, msgerr.New(errOnUser, "Could not find a user with the username")
		}
		return nil, msgerr.New(err, "Invalid ID")
	}
	result, err := r.hackerNews.GetItem(ctx, idN)
	if err != nil {
		return nil, msgerr.New(err, "Could not get item")
	}
	return ItemToNode(result)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
