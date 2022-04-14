package relays

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Invoke(registerDependencies)
}
