package admin

import (
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
			func(api AdminAPI, args common.VersionInputArgs) {
				common.RegisterHealthcheck(api)
				common.RegisterVersion(api, args)
				auth.RegisterLogin(api)
			},
			testingCache,
		),
	)
}
