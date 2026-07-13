package configs

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConfigs),

	fx.Provide(func(cfg *Config) *AppConfig { return &cfg.App }),
	fx.Provide(func(cfg *Config) *WebServerConfig { return &cfg.WebServer }),
	fx.Provide(func(cfg *Config) *JWTConfig { return &cfg.JWT }),
)
