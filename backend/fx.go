package backend

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/grpc/readabilityserver"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		images.Module(),
		readabilityserver.Module(),
		fx.Invoke(RegisterRoutes),
	)
}
