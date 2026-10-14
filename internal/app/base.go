package app

import (
	"platform/internal/infrastructure/cache"
	"platform/internal/infrastructure/configs"
	"platform/internal/infrastructure/databases"
	"platform/internal/infrastructure/logger"
	"platform/internal/infrastructure/tools"
	"platform/internal/usecases"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func NewBaseOptions() fx.Option {
	return fx.Options(
		configs.Module,
		logger.Module,
		databases.Module,
		cache.Module,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{
				Logger: log.With(zap.String("kind", "fx")),
			}
		}),

		fx.Invoke(func(log *zap.Logger) {
			log.Info("Application initialized successfully")
		}),

		tools.InjectDI(),
		usecases.InjectDI(),
	)
}
