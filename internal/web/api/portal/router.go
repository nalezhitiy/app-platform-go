package portal

import (
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

var basePath = "/api/portal"

type PortalAPI struct {
	web.BaseAPI
}

func NewPortalAPIRouter(router *chi.Mux) PortalAPI {
	cfg := huma.DefaultConfig("Portal API", "1.0.0")

	cfg.Servers = []*huma.Server{
		{
			URL:         basePath,
			Description: "Development server",
		},
	}

	cfg.Info.Description =
		"<h2>Information</h2>" +
			"Provides authentication, account management and business operations for the Platform Portal.<br>" +
			"<b>Authentication</b>: JWT Bearer or Session Cookie.<br>" +
			"Responses use JSON.<br>Errors follow RFC 7807.<br>" +
			"All timestamps are returned in UTC (ISO 8601)."

	cfg.Info.Contact = &huma.Contact{
		Name: "Platform Team",
	}

	cfg.CreateHooks = nil

	cfg.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
		"sessionCookie": {
			Type: "apiKey",
			In:   "cookie",
			Name: "session",
		},
	}

	route := router.Route(basePath, func(r chi.Router) {})
	api := humachi.New(route, cfg)

	return PortalAPI{web.BaseAPI{
		API:      api,
		Name:     "portal-api",
		BasePath: basePath,
	}}
}
