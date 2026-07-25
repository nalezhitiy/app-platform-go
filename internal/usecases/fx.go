package usecases

import "go.uber.org/fx"

func InjectDI() fx.Option {
	return fx.Options(
		fx.Provide(NewVersionDBUseCase),
	)
}
