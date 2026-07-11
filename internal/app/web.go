package app

import (
	"platform/internal/web/api/admin"
	"platform/internal/web/api/portal"

	"go.uber.org/fx"
)

import "platform/internal/infrastructure/transport/http"

var (
	AdminModule = fx.Module(
		"admin-api",
		admin.InjectDI(),
	)

	PortalModule = fx.Module(
		"portal-api",
		portal.InjectDI(),
	)
)

func NewWebAppOptions() fx.Option {
	return fx.Options(
		NewBaseOptions(),
		fx.Provide(http.NewBaseRouter),
		fx.Provide(http.NewHTTPServer),
		fx.Invoke(http.ServerRegisterHooks),
	)
}

func NewWebAllApp() *fx.App {
	return fx.New(
		NewWebAppOptions(),
		AdminModule,
		PortalModule,
	)
}

func NewWebAdminApp() *fx.App {
	return fx.New(
		NewWebAppOptions(),
		AdminModule,
	)
}
func NewWebPortalApp() *fx.App {
	return fx.New(
		NewWebAppOptions(),
		PortalModule,
	)
}
