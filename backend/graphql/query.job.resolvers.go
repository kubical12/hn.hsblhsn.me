package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/tasylab/hn.hsblhsn.me/backend/graphql/model"
)

// JobStories is the resolver for the jobStories field.
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
