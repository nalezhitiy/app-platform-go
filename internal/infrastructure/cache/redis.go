package cache

import (
	"context"
	"errors"
	"fmt"
	"platform/internal/infrastructure/configs"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
	cfg    *configs.CacheConfig
}

func NewRedisCache(cfg *configs.CacheConfig) *RedisCache {
	return &RedisCache{cfg: cfg}
}

func (c *RedisCache) Open(ctx context.Context) error {
	options, err := redis.ParseURL(c.cfg.DSN)
	if err != nil {
		return fmt.Errorf("parse Redis DSN: %w", err)
	}

	client := redis.NewClient(options)

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return fmt.Errorf("connect to Redis: %w", err)
	}

	c.client = client
	return nil
}

func (c *RedisCache) Close() error {
	if c.client == nil {
		return nil
	}

	err := c.client.Close()
	c.client = nil

	return err
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", ErrCacheMiss
	}
	if err != nil {
		return "", fmt.Errorf("redis get %q: %w", key, err)
	}

	return value, nil
}

func (c *RedisCache) Set(
	ctx context.Context,
	key string,
	value string,
	ttl time.Duration,
) error {
	if ttl < 0 {
		return fmt.Errorf("redis set %q: TTL cannot be negative", key)
	}

	if err := c.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return fmt.Errorf("redis set %q: %w", key, err)
	}

	return nil
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	if err := c.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis delete %q: %w", key, err)
	}

	return nil
}

func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	count, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("redis exists %q: %w", key, err)
	}

	return count > 0, nil
}

func (c *RedisCache) Version(ctx context.Context) (string, error) {
	info, err := c.client.Info(ctx, "server").Result()
	if err != nil {
		return "", fmt.Errorf("get redis server info: %w", err)
	}

	for line := range strings.SplitSeq(info, "\r\n") {
		version, ok := strings.CutPrefix(line, "redis_version:")
		if ok {
			return version, nil
		}
	}

	return "", errors.New("redis_version not found in server info")
}

func (c *RedisCache) Name() string {
	return c.cfg.Name
}
