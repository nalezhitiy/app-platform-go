package admin

import (
	"platform/internal/infrastructure/configs"
	"platform/internal/web/api/common"
	"platform/internal/web/api/common/auth"

	"go.uber.org/fx"
)

func InjectDI() fx.Option {
	return fx.Options(
		fx.Provide(
			NewAdminAPIRouter,
		),
		fx.Invoke(
			func(api AdminAPI, cfg *configs.Config) {
				common.RegisterHealthcheck(api)
				common.RegisterVersion(api, cfg)
				auth.RegisterLogin(api)
			},
		),
	)
}
