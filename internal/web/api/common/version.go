package common

import (
	"context"
	"net/http"
	"platform/internal/infrastructure/configs"
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
)

type versionBody struct {
	Component string `json:"component"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildAt   string `json:"build_at"`
}

type versionResponse struct {
	Body versionBody `huma:"body, application/json"`
}

func RegisterVersion(api web.APIInstance, cfg *configs.Config) {
	huma.Register(api, huma.Operation{
		OperationID: "version",
		Summary:     "Application version",
		Method:      http.MethodGet,
		Path:        "/version",
		Tags:        []string{"public"},
	}, func(ctx context.Context, _ *struct{}) (*versionResponse, error) {

		return &versionResponse{
			Body: versionBody{
				Component: api.Name,
				Version:   cfg.App.Version,
				Commit:    cfg.App.Commit,
				BuildAt:   cfg.App.BuildAt,
			},
		}, nil
	})
}
