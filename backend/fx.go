package backend

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/images"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/caches"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/tracing"
	"go.uber.org/fx"
)

// Module for fx.
func Module() fx.Option {
	return fx.Options(
		tracing.Module(),
		caches.Module(),
		httpclient.Module(),
		images.Module(),
		fx.Invoke(RegisterRoutes),
	)
}
