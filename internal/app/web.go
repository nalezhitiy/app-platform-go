package app

import (
	"platform/internal/web/api/admin"

	"go.uber.org/fx"
)

import "platform/internal/infrastructure/transport/http"

func NewWebAppOptions() fx.Option {
	return fx.Options(
		NewBaseOptions(),
		fx.Provide(http.NewBaseRouter),
		fx.Provide(http.NewHTTPServer),
		fx.Invoke(http.ServerRegisterHooks),
	)
}

func NewWebAdminApp() *fx.App {
	AdminModule := fx.Module("admin-api", admin.InjectDI())
	return fx.New(
		NewWebAppOptions(),
		AdminModule,
	)
}
