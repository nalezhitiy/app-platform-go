package portal

import (
	"platform/internal/infrastructure/configs"
	"platform/internal/web/api/common"

	"go.uber.org/fx"
)

func InjectDI() fx.Option {
	return fx.Options(
		fx.Provide(
			NewPortalAPIRouter,
		),
		fx.Invoke(
			func(api PortalAPI, cfg *configs.Config) {
				common.RegisterHealthcheck(api)
				common.RegisterVersion(api, cfg)
			},
		),
	)
}
