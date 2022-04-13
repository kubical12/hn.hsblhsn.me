package graph

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/grpc/readabilityserver"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/model"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		model.Module(),
		httpclient.Module(),
		hackernews.Module(),
		readabilityserver.Module(),
		fx.Provide(NewResolver),
		fx.Provide(NewImageHandler),
		fx.Provide(NewGraphQLHandler),
	)
}
