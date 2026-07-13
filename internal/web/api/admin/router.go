package admin

import (
	"platform/internal/web"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

var basePath = "/api/admin"

type AdminAPI struct {
	web.BaseAPI
}

func NewAdminAPIRouter(router *chi.Mux) AdminAPI {
	cfg := huma.DefaultConfig("Admin API", "1.0.0")
	cfg.DocsRenderer = huma.DocsRendererStoplightElements

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
	return AdminAPI{BaseAPI: web.BaseAPI{
		API:            api,
		Name:           "admin-api",
		BasePath:       basePath,
		AuthCookieName: "session-admin",
	}}
}
