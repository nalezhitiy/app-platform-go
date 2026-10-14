package usecases

import (
	"platform/internal/usecases/system"

	"go.uber.org/fx"
)

func InjectDI() fx.Option {
	return fx.Options(
		fx.Provide(
			system.NewVersionDBUseCase,
			system.NewVersionCacheUseCase,
		),
	)
}
