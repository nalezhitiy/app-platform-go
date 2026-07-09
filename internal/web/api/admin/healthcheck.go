package admin

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type PlainTextResponse struct {
	Body string `huma:"body,text/plain"`
}

func healthcheck(API AdminAPI) {
	desc := huma.Operation{
		OperationID: "healthcheck",
		Summary:     "Plain health check",
		Tags:        []string{"public"},
		Method:      http.MethodGet,
		Path:        "/healthcheck",
		Security:    nil,
	}

	huma.Register(API, desc, func(ctx context.Context, _ *struct{}) (*PlainTextResponse, error) {
		return &PlainTextResponse{Body: "Ok"}, nil
	})
}
