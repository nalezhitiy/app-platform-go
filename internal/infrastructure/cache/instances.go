package cache

import (
	"fmt"

	"platform/internal/infrastructure/configs"

	"go.uber.org/zap"
)

type MainCache struct {
	*RedisCache
}

func NewMainCache(
	log *zap.Logger,
	cfg *configs.PlatformCacheConfig,
) (*MainCache, error) {
	cacheCfg, err := cfg.Get("main-redis")
	if err != nil {
		return nil, err
	}

	if cacheCfg.Driver != "redis" {
		return nil, fmt.Errorf(
			"unsupported cache driver %q for cache %q",
			cacheCfg.Driver,
			cacheCfg.Name,
		)
	}

	log = log.With(zap.String("cache", cacheCfg.Name))

	return &MainCache{
		RedisCache: NewRedisCache(cacheCfg),
	}, nil
}

type SessionCache struct {
	*RedisCache
}

func NewSessionCache(
	log *zap.Logger,
	cfg *configs.PlatformCacheConfig,
) (*SessionCache, error) {
	cacheCfg, err := cfg.Get("sessions-redis")
	if err != nil {
		return nil, err
	}

	if cacheCfg.Driver != "redis" {
		return nil, fmt.Errorf(
			"unsupported cache driver %q for cache %q",
			cacheCfg.Driver,
			cacheCfg.Name,
		)
	}

	log = log.With(zap.String("cache", cacheCfg.Name))

	return &SessionCache{
		RedisCache: NewRedisCache(cacheCfg),
	}, nil
}
