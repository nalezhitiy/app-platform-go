package app

import (
	"platform/internal/infrastructure/configs"
	"platform/internal/infrastructure/logger"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewBaseOptions() fx.Option {
	return fx.Options(
		fx.Provide(
			configs.NewConfigs,
			logger.NewLogger,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			fxLogger := &fxevent.ZapLogger{
				Logger: log.With(zap.String("kind", "fx")),
			}
			fxLogger.UseLogLevel(zapcore.DebugLevel)
			return fxLogger
		}),

		fx.Invoke(func(log *zap.Logger) {
			log.Info("Application initialized successfully")
		}),
	)
}
