package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/algolia"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, query string, after *string, first *int) (*relays.Connection[model.Node], error) {
	var (
		page    = 1
		perPage = 10
	)
	if after != nil {
		afterN, err := strconv.Atoi(*after)
		if err != nil {
			return nil, msgerr.New(err, "Invalid after")
		}
		page = afterN + 1
	}
	if first != nil {
		perPage = *first
	}
	result, err := r.algolia.Search(ctx, "story", query, &algolia.PaginationInput{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, msgerr.New(err, "Could not search")
	}
	conn, err := result.ToConnection()
	if err != nil {
		return nil, msgerr.New(err, "Could not build node connection")
	}
	return conn, nil
}
