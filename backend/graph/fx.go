package graph

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/model"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		model.Module(),
		httpclient.Module(),
		hackernews.Module(),
		fx.Provide(NewResolver),
		fx.Provide(NewServer),
	)
}
