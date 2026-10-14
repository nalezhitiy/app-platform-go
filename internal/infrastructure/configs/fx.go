package configs

import "go.uber.org/fx"

var Module = fx.Module(
	"configs",

	fx.Provide(
		NewConfigs,
		NewAppConfig,
		NewWebServerConfig,
		NewJWTConfig,
		NewPlatformDBConfig,
		NewCacheConfig,
	),
)

func NewAppConfig(cfg *Config) *AppConfig {
	return &cfg.App
}

func NewWebServerConfig(cfg *Config) *WebServerConfig {
	return &cfg.WebServer
}

func NewJWTConfig(cfg *Config) *JWTConfig {
	return &cfg.JWT
}

func NewPlatformDBConfig(cfg *Config) *PlatformDBConfig {
	return &cfg.Databases
}

func NewCacheConfig(cfg *Config) *PlatformCacheConfig {
	return &cfg.Cache
}
