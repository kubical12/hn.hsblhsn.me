package graphql

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/extensions/complexity"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/extensions/timeout"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/featureflags"
	"github.com/pkg/errors"
	"github.com/ravilushqa/otelgqlgen"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type (
	GQLHandler = handler.Server
)

const (
	MaxQueryComplexity     = 300
	complexityNetworkField = 1
	DefaultTimeout         = time.Second * 5
)

func getHTMLFieldComplexity() int {
	if !featureflags.IsOn(featureflags.FeatureReadability, false) {
		return 0
	}
	return 10
}

// ComplexityMap is a map of field names to their maximum complexity.
// The map is used to calculate the complexity of a query.
//
//nolint:gochecknoglobals // initialize once globally.
var ComplexityMap = complexity.Map{
	"StoryConnection":      complexityNetworkField,
	"Story":                complexityNetworkField,
	"Story.html":           getHTMLFieldComplexity(),
	"CommentConnection":    complexityNetworkField,
	"Comment":              complexityNetworkField,
	"JobConnection":        complexityNetworkField,
	"Job":                  complexityNetworkField,
	"Job.html":             getHTMLFieldComplexity(),
	"PollConnection":       complexityNetworkField,
	"Poll":                 complexityNetworkField,
	"PollOptionConnection": complexityNetworkField,
	"PollOption":           complexityNetworkField,
	"User":                 complexityNetworkField,
	"OpenGraph":            complexityNetworkField * 2, //nolint:gomnd
	"PageInfo":             complexityNetworkField,
}

// NewGQLHandler creates a new graphql handler.
func NewGQLHandler(resolver *Resolver, logger *zap.Logger, tracer trace.TracerProvider) *GQLHandler {
	config := generated.Config{Resolvers: resolver}
	server := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	// setup extensions.
	server.Use(timeout.NewExtension(DefaultTimeout))
	server.Use(complexity.NewExtension(MaxQueryComplexity, ComplexityMap))
	server.Use(otelgqlgen.Middleware(otelgqlgen.WithTracerProvider(tracer)))
	// setup error presenters.
	server.SetErrorPresenter(newErrorPresenterFunc(logger))
	server.SetRecoverFunc(newRecoverFunc(logger))
	return server
}

// newErrorPresenterFunc returns a function that omits all system errors.
// But if the error is a msgerr.Error, it will be presented.
// All other errors will be presented as a generic error message.
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

// newRecoverFunc returns a recover function that logs the panic message.
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
