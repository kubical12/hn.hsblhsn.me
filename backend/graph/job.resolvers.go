package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/model"
)

func (r *jobResolver) Type(ctx context.Context, obj *model.Job) (string, error) {
	return obj.Type.String(), nil
}

// Job returns generated.JobResolver implementation.
func (r *Resolver) Job() generated.JobResolver { return &jobResolver{r} }

type jobResolver struct{ *Resolver }
