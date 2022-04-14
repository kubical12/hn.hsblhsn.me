package backend

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		graphql.Module(),
		fx.Invoke(RegisterRoutes),
	)
}
