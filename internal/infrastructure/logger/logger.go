package logger

import (
	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger initializes and configures a Uber Zap logger based on application config
func NewLogger(cfg *configs.Config) (*zap.Logger, error) {
	var zapCfg zap.Config

	if cfg.App.IsDevelop() {
		zapCfg = zap.NewDevelopmentConfig()

		zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		zapCfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	} else {
		zapCfg = zap.NewProductionConfig()
		zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	level, err := zapcore.ParseLevel(cfg.App.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}
	zapCfg.Level = zap.NewAtomicLevelAt(level)

	log, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	return log.With(zap.String("app", cfg.App.Name)), nil
}
