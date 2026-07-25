package portal

import (
	"platform/internal/web/api/common"

	"go.uber.org/fx"
)

func InjectDI() fx.Option {
	return fx.Options(
		fx.Provide(
			NewPortalAPIRouter,
		),
		fx.Invoke(
			func(api PortalAPI, args common.VersionInputArgs) {
				common.RegisterHealthcheck(api)
				common.RegisterVersion(api, args)
			},
		),
	)
}
