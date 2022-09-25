package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	opengraph "github.com/otiai10/opengraph/v2"
)

// ID is the resolver for the id field.
func (r *openGraphResolver) ID(ctx context.Context, obj *opengraph.OpenGraph) (string, error) {
	//nolint:gosec // not a cryptographic operation
	hashBytes := md5.Sum([]byte(obj.URL))
	return hex.EncodeToString(hashBytes[:]), nil
}

// OpenGraph returns generated.OpenGraphResolver implementation.
func (r *Resolver) OpenGraph() generated.OpenGraphResolver { return &openGraphResolver{r} }

type openGraphResolver struct{ *Resolver }
