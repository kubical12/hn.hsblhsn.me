package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/complexity"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/images"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

type (
	ImageHandler = images.ImageResizeHandler
	GQLHandler   = handler.Server
)

const (
	MaxQueryComplexity      = 500
	complexityNetworkField  = 1
	complexityComputedValue = 10
)

// ComplexityMap is a map of field names to their maximum complexity.
// The map is used to calculate the complexity of a query.
// nolint: gochecknoglobals // initialize once globally.
var ComplexityMap = complexity.Map{
	"StoryConnection":      complexityNetworkField,
	"Story":                complexityNetworkField,
	"Story.html":           complexityComputedValue,
	"CommentConnection":    complexityNetworkField,
	"Comment":              complexityNetworkField,
	"JobConnection":        complexityNetworkField,
	"Job":                  complexityNetworkField,
	"Job.html":             complexityComputedValue,
	"PollConnection":       complexityNetworkField,
	"Poll":                 complexityNetworkField,
	"PollOptionConnection": complexityNetworkField,
	"PollOption":           complexityNetworkField,
	"User":                 complexityNetworkField,
	"OpenGraph":            complexityNetworkField * 2, // nolint:gomnd
}

func NewImageHandler(client *httpclient.CachedClient, logger *zap.Logger) *ImageHandler {
	return images.NewImageResizeHandler(client, logger)
}

func NewGQLHandler(resolver *Resolver, logger *zap.Logger) *GQLHandler {
	config := generated.Config{Resolvers: resolver}
	server := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	// setup extensions.
	server.Use(complexity.NewLimitExtension(MaxQueryComplexity, ComplexityMap))
	// setup error presenters.
	server.SetErrorPresenter(newErrorPresenterFunc(logger))
	server.SetRecoverFunc(newRecoverFunc(logger))
	return server
}

func newErrorPresenterFunc(logger *zap.Logger) graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {
		log := logger.With(
			zap.String("component", "graphql_error_presenter"),
			zap.Error(err),
		)
		gqlErr := graphql.DefaultErrorPresenter(ctx, err)
		var msgErr *msgerr.Error
		if errors.As(err, &msgErr) {
			gqlErr.Message = msgErr.Msg()
		} else {
			gqlErr.Message = "Internal server error"
		}
		log.Error("graphql error", zap.String("response", gqlErr.Message))
		return gqlErr
	}
}

func newRecoverFunc(logger *zap.Logger) graphql.RecoverFunc {
	return func(ctx context.Context, err interface{}) error {
		log := logger.With(
			zap.String("component", "graphql_recover_func"),
			zap.Any("recovered", err),
		)
		log.Error("graphql panic", zap.Any("recovered", err))
		return gqlerror.Errorf("Internal server error!")
	}
}
