package backend

import (
	"go.uber.org/fx"
)

// Module for fx.
func Module() fx.Option {
	return fx.Options(
		fx.Invoke(RegisterRoutes),
	)
}
