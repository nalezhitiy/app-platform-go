package common

import (
	"context"
	"net/http"
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
)

type plainTextResponse struct {
	Body string `huma:"body,text/plain"`
}

func RegisterHealthcheck(api web.API) {
	desc := huma.Operation{
		OperationID: "healthcheck",
		Summary:     "Plain health check",
		Tags:        []string{"public"},
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Security:    nil,
	}

	huma.Register(api, desc, func(ctx context.Context, _ *struct{}) (*plainTextResponse, error) {
		return &plainTextResponse{Body: "Ok"}, nil
	})
}
