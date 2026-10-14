package configs

import (
	"fmt"
	"time"
)

type CacheConfig struct {
	Name   string          `yaml:"name"`
	Driver string          `yaml:"driver"`
	DSN    string          `yaml:"dsn"`
	Pool   CachePoolConfig `yaml:"pool"`
}

type CachePoolConfig struct {
	Min      int           `yaml:"min"`
	Max      int           `yaml:"max"`
	Lifetime time.Duration `yaml:"lifetime"`
}

type PlatformCacheConfig map[string]CacheConfig

func (c PlatformCacheConfig) Get(name string) (*CacheConfig, error) {
	cfg, ok := c[name]
	if !ok {
		return nil, fmt.Errorf("cache config %q not found", name)
	}

	return &cfg, nil
}
