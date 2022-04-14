package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
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

func NewImageHandler(client *httpclient.CachedClient, logger *zap.Logger) *ImageHandler {
	return images.NewImageResizeHandler(client, logger)
}

func NewGQLHandler(resolver *Resolver, logger *zap.Logger) *GQLHandler {
	server := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: resolver},
		),
	)
	server.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
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
	})
	server.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		log := logger.With(zap.String("component", "graphql_recover_func"))
		log.Error("graphql panic", zap.Any("recovered", err))
		return gqlerror.Errorf("Internal server error!")
	})
	return server
}
