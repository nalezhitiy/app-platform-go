package app

import "go.uber.org/fx"

func NewCliAppOptions() fx.Option {
	return fx.Options(
		fx.NopLogger, // disable default logger fx
		NewBaseOptions(),
	)
}
