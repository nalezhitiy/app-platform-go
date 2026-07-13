package app

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewCliAppOptions() fx.Option {
	return fx.Options(
		NewBaseOptions(),
		// Виключаємо всі логи для консольних команд
		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),
		fx.Replace(zap.NewNop()),
	)
}
