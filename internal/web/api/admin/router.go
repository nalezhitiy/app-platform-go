package admin

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

type AdminAPI huma.API

var basePath = "/api/admin"

func NewAdminAPIRouter(router *chi.Mux) AdminAPI {
	cfg := huma.DefaultConfig("Admin API", "1.0.0")
	// huma.DocsRendererSwaggerUI
	// huma.DocsRendererScalar
	// huma.DocsRendererStoplightElements
	//cfg.DocsRenderer = huma.DocsRendererStoplightElements
	cfg.Servers = []*huma.Server{
		{
			URL:         basePath,
			Description: "V0 Development server",
		},
	}

	cfg.CreateHooks = nil

	cfg.Tags = []*huma.Tag{
		{
			Name:        "Users",
			Description: "User management",
		},
		{
			Name: "Projects",
		},
		{
			Name: "Statistics",
		},
	}

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
	return humachi.New(route, cfg)
}
