package admin

import (
	"platform/internal/web/api/common"

	"go.uber.org/fx"
)

var routeName = "admin-route"

func InjectDI() fx.Option {

	return fx.Options(
		fx.Provide(
			fx.Annotate(
				NewAdminAPIRouter,
				fx.ResultTags(`name:"`+routeName+`"`),
			),
		),
		common.InjectDI(routeName),
	)
}
