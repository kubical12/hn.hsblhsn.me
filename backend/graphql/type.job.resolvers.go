package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

// Type is the resolver for the type field.
func (r *jobResolver) Type(ctx context.Context, obj *model.Job) (string, error) {
	return obj.Type.String(), nil
}

// By is the resolver for the by field.
func (r *jobResolver) By(ctx context.Context, obj *model.Job) (*model.User, error) {
	resp, err := r.hackerNews.GetUser(ctx, obj.By)
	if err != nil {
		return nil, msgerr.New(err, "Could not get user")
	}
	return &model.User{UserResponse: resp}, nil
}

// Job returns generated.JobResolver implementation.
func (r *Resolver) Job() generated.JobResolver { return &jobResolver{r} }

type jobResolver struct{ *Resolver }
