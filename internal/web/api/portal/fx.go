package portal

import (
	"platform/internal/web/api/common"

	"go.uber.org/fx"
)

var routeName = "portal-route"

func InjectDI() fx.Option {

	return fx.Options(
		fx.Provide(
			fx.Annotate(
				NewPortalAPIRouter,
				fx.ResultTags(`name:"`+routeName+`"`),
			),
		),
		common.InjectDI(routeName),
	)
}
