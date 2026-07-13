package common

import (
	"context"
	"net/http"
	"platform/internal/infrastructure/configs"
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
)

type version struct {
	Component string `json:"component"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildAt   string `json:"build_at"`
}

func RegisterVersion(api web.APIInterface, cfg *configs.Config) {
	huma.Register(api, huma.Operation{
		OperationID: "version",
		Summary:     "Application version",
		Method:      http.MethodGet,
		Path:        "/version",
		Tags:        []string{"public"},
	}, func(ctx context.Context, _ *struct{}) (*web.BaseBodyResponse[version], error) {

		return web.NewBaseBodyResponse(version{
			Component: api.GetName(),
			Version:   cfg.App.Version,
			Commit:    cfg.App.Commit,
			BuildAt:   cfg.App.BuildAt,
		}), nil
	})
}
