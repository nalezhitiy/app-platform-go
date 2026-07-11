package admin

import (
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

var basePath = "/api/admin"

func NewAdminAPIRouter(router *chi.Mux) web.APIInstance {
	cfg := huma.DefaultConfig("Admin API", "1.0.0")
	// huma.DocsRendererSwaggerUI
	// huma.DocsRendererScalar
	// huma.DocsRendererStoplightElements
	//cfg.DocsRenderer = huma.DocsRendererStoplightElements
	cfg.Servers = []*huma.Server{
		{
			URL:         basePath,
			Description: "Development server",
		},
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
	return web.APIInstance{
		API:      api,
		Name:     "admin-api",
		BasePath: basePath,
	}
}
