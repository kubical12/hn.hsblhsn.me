package images

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewSocialPreviewGenerator),
		fx.Provide(NewImageProxyHandler),
		fx.Provide(NewSocialPreviewHandler),
	)
}
