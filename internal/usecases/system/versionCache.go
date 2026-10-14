package system

import (
	"context"
	"platform/internal/infrastructure/cache"
)

type VersionCacheUseCase struct {
	main     *cache.MainCache
	sessions *cache.SessionCache
}

func NewVersionCacheUseCase(main *cache.MainCache,
	sessions *cache.SessionCache) *VersionCacheUseCase {
	return &VersionCacheUseCase{
		main:     main,
		sessions: sessions,
	}
}

type CacheVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (uc *VersionCacheUseCase) GetVersion() ([]CacheVersion, error) {
	var result []CacheVersion
	ctx := context.Background()

	ver, err := uc.main.Version(ctx)
	if err != nil {
		return nil, err
	}
	result = append(result, CacheVersion{
		Name:    uc.main.Name(),
		Version: ver,
	})

	ver, err = uc.sessions.Version(ctx)
	if err != nil {
		return nil, err
	}
	result = append(result, CacheVersion{
		Name:    uc.main.Name(),
		Version: ver,
	})

	return result, nil
}
