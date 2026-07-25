package configs

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConfigs),

	fx.Provide(
		func(cfg *Config) *AppConfig { return &cfg.App },
		func(cfg *Config) *WebServerConfig { return &cfg.WebServer },
		func(cfg *Config) *JWTConfig { return &cfg.JWT },
		func(cfg *Config) *PlatformDBConfig { return &cfg.Databases },
	),
)
