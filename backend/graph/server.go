package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/msgerr"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func NewServer(resolver *Resolver, logger *zap.Logger) *handler.Server {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)
	server.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log := logger.With(
			zap.String("component", "graphql/error_presenter"),
			zap.Error(err),
		)
		gqlErr := graphql.DefaultErrorPresenter(ctx, err)
		var msgErr *msgerr.Error
		if errors.As(err, &msgErr) {
			gqlErr.Message = msgErr.Msg()
		} else {
			gqlErr.Message = "Internal server error"
		}
		log.Error("graphql error", zap.String("message", gqlErr.Message))
		return gqlErr
	})
	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log := logger.With(zap.String("component", "graphql/recover_func"))
		log.Error("graphql panic", zap.Any("recovered", err))
		return gqlerror.Errorf("Internal server error!")
	})
	return server
}
