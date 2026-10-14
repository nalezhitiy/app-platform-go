package admin

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"platform/internal/infrastructure/cache"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

type DataOutput struct {
	Value string `json:"value"`
}

type CacheResponse struct {
	Body DataOutput
}

var TAG = "test"

func testingCache(api AdminAPI, mainCache *cache.MainCache) {
	desc := huma.Operation{
		OperationID: "test cache",
		Summary:     "Test cache to platform",
		Tags:        []string{TAG},
		Method:      http.MethodGet,
		Path:        fmt.Sprintf("/%s/cache", TAG),
		Security:    nil,
	}

	handler := func(ctx context.Context, _ *struct{}) (*CacheResponse, error) {
		var key = "test"
		val, err := mainCache.Get(ctx, key)

		if errors.Is(err, cache.ErrCacheMiss) {
			val = time.Now().String()
			mainCache.Set(ctx, key, val, 5*time.Second)
		} else if err != nil {
			return nil, err
		}

		return &CacheResponse{DataOutput{
			Value: val,
		},
		}, nil
	}

	huma.Register(api, desc, handler)
}
