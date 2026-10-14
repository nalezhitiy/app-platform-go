package common

import (
	"context"
	"net/http"
	"platform/internal/infrastructure/configs"
	"platform/internal/usecases/system"
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

type VersionInputArgs struct {
	fx.In

	Config           *configs.Config
	UseCaseVersionDB *system.VersionDBUseCase
	UseCaseCache     *system.VersionCacheUseCase
}

type version struct {
	Component string                   `json:"component"`
	Version   string                   `json:"version"`
	Commit    string                   `json:"commit"`
	BuildAt   string                   `json:"build_at"`
	Databases []system.DatabaseVersion `json:"databases"`
	Caches    []system.CacheVersion    `json:"caches"`
}

func RegisterVersion(
	API web.APIInterface,
	args VersionInputArgs,
) {
	huma.Register(API, huma.Operation{
		OperationID: "version",
		Summary:     "Application version",
		Method:      http.MethodGet,
		Path:        "/version",
		Tags:        []string{"public"},
	}, func(ctx context.Context, _ *struct{}) (*web.BaseBodyResponse[version], error) {
		cfg := args.Config

		v := version{
			Component: API.GetName(),
			Version:   cfg.App.Version,
			Commit:    cfg.App.Commit,
			BuildAt:   cfg.App.BuildAt,
			Databases: make([]system.DatabaseVersion, 0),
		}

		versionsDB, err := args.UseCaseVersionDB.GetVersion()
		if err != nil {
			return nil, err
		}
		v.Databases = versionsDB

		versionsCache, err := args.UseCaseCache.GetVersion()
		if err != nil {
			return nil, err
		}
		v.Caches = versionsCache

		return web.NewBaseBodyResponse(v), nil
	})
}
